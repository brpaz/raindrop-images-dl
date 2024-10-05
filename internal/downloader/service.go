package downloader

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

var (
	ErrRaindropClientNotSet = errors.New("raindrop client not set")
	ErrCollectionIDNotSet   = errors.New("collection ID not set")
	ErrOutputDirNotSet      = errors.New("output directory not set")
	ErrOutputDirNotExists   = errors.New("output directory does not exist")
)

type RaindropClient interface {
	GetCollectionByID(ctx context.Context, collectionID int) (*raindrop.CollectionItem, error)
	GetImagesDropsFromCollection(ctx context.Context, collectionID int, page int) (*raindrop.ImageDrops, error)
}

// Downloader is a client for the Raindrop API
type Downloader struct {
	rdClient RaindropClient
}

// Validate validates the Downloader configuration
func (d *Downloader) Validate() error {
	if d.rdClient == nil {
		return ErrRaindropClientNotSet
	}
	return nil
}

// Option defines a functional option type for configuring the Downloader
type Option func(*Downloader)

// WithRaindropClient is a functional option to set the Raindrop client
func WithRaindropClient(client RaindropClient) Option {
	return func(d *Downloader) {
		d.rdClient = client
	}
}

// NewDownloader creates a new Downloader instance, applying any provided options
func NewDownloader(opts ...Option) (*Downloader, error) {
	dl := &Downloader{}

	for _, opt := range opts {
		opt(dl)
	}

	if err := dl.Validate(); err != nil {
		return nil, err
	}

	return dl, nil
}

// DownloadCollection downloads all images from a Raindrop collection
func (d *Downloader) DownloadCollection(ctx context.Context, collectionID int, outputDir string) error {
	if collectionID == 0 {
		return ErrCollectionIDNotSet
	}

	if outputDir == "" {
		return ErrOutputDirNotSet
	}

	// Ensure the output directory exists
	if !dirExists(outputDir) {
		return ErrOutputDirNotExists
	}

	collection, err := d.rdClient.GetCollectionByID(ctx, collectionID)
	if err != nil {
		return fmt.Errorf("failed to get collection with id %d: %w", collectionID, err)
	}

	slog.Info("Downloading collection", "name", collection.Title)

	page := 0
	for {
		slog.Info("Processing page", "page", page)

		items, err := d.rdClient.GetImagesDropsFromCollection(ctx, collectionID, page)
		if err != nil {
			slog.Error("Failed to get images from collection", "collection", collection.Title, "page", page, "error", err)
			break
		}

		// Download each item
		for _, item := range items.Items {
			slog.Info("Downloading item", "title", item.Title)

			if err := d.downloadItem(ctx, item, collection.Title, outputDir); err != nil {
				slog.Error("Failed to download item", "title", item.Title, "error", err)
			}
		}

		// Exit if no more items to process
		if !items.HasMore {
			break
		}
		page++
	}

	return nil
}

// downloadItem handles downloading an individual item
func (d *Downloader) downloadItem(ctx context.Context, item raindrop.Drop, collectionName, outputDir string) error {
	// Ensure collection-specific directory exists
	itemOutputDir := filepath.Join(outputDir, collectionName)
	if err := ensureDir(itemOutputDir); err != nil {
		return fmt.Errorf("failed to create directory for collection: %w", err)
	}

	imageURL := item.GetFileLink()
	if imageURL == "" {
		slog.Warn("Bookmark has no URL field", "title", item.Title)
		return nil
	}

	// Download image
	baseFilePath := filepath.Join(itemOutputDir, item.GetName())
	if err := downloadFile(imageURL, baseFilePath); err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}

	// Create info.json file
	if err := createInfoFile(baseFilePath, item); err != nil {
		return fmt.Errorf("failed to create info file: %w", err)
	}

	return nil
}
