package downloader_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/brpaz/raindrop-images-dl/internal/downloader"
	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

type MockRaindropClient struct {
	mock.Mock
}

func (m *MockRaindropClient) GetCollectionByID(ctx context.Context, collectionID int) (*raindrop.CollectionItem, error) {
	args := m.Called(ctx, collectionID)
	return args.Get(0).(*raindrop.CollectionItem), args.Error(1)
}

func (m *MockRaindropClient) GetImagesDropsFromCollection(ctx context.Context, collectionID int, page int) (*raindrop.ImageDrops, error) {
	args := m.Called(ctx, collectionID, page)
	return args.Get(0).(*raindrop.ImageDrops), args.Error(1)
}

func setupTestDownloader(t *testing.T) (*downloader.Downloader, *MockRaindropClient) {
	t.Helper()

	rdClient := &MockRaindropClient{}

	dl, err := downloader.NewDownloader(
		downloader.WithRaindropClient(rdClient),
	)
	require.NoError(t, err)

	return dl, rdClient
}

func generateTmpDir(t *testing.T) string {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "raindrop-dl-images-test")
	require.NoError(t, err)

	return tmpDir
}

func TestNewDownloader(t *testing.T) {
	t.Parallel()

	client := &MockRaindropClient{}

	t.Run("WithMissingRaindropClient_ReturnsError", func(t *testing.T) {
		t.Parallel()
		dl, err := downloader.NewDownloader()

		assert.ErrorIs(t, err, downloader.ErrRaindropClientNotSet)
		assert.Nil(t, dl)
	})

	t.Run("WithRaindropClient", func(t *testing.T) {
		t.Parallel()
		dl, err := downloader.NewDownloader(downloader.WithRaindropClient(client))

		assert.NoError(t, err)
		assert.NotNil(t, dl)
	})
}

func TestDownloader_DownloadCollection(t *testing.T) {
	t.Parallel()

	t.Run("WithInvalidCollectionID_ReturnsError", func(t *testing.T) {
		t.Parallel()

		dl, _ := setupTestDownloader(t)

		err := dl.DownloadCollection(context.Background(), 0, "output", false)
		assert.ErrorIs(t, err, downloader.ErrCollectionIDNotSet)
	})

	t.Run("WithEmptyOutputDir_ReturnsError", func(t *testing.T) {
		t.Parallel()

		dl, _ := setupTestDownloader(t)

		err := dl.DownloadCollection(context.Background(), 123, "", false)
		assert.ErrorIs(t, err, downloader.ErrOutputDirNotSet)
	})

	t.Run("WithNonExistentOutputDir_ReturnsError", func(t *testing.T) {
		t.Parallel()

		dl, _ := setupTestDownloader(t)
		err := dl.DownloadCollection(context.Background(), 123, "non-existent-dir", false)
		assert.ErrorIs(t, err, downloader.ErrOutputDirNotExists)
	})

	t.Run("Success With Info File", func(t *testing.T) {
		t.Parallel()

		dl, rdClient := setupTestDownloader(t)
		// create tmp dir
		outputDir := generateTmpDir(t)
		defer os.RemoveAll(outputDir)

		collectionID := 123

		rdClient.On("GetCollectionByID", mock.Anything, collectionID).Return(&raindrop.CollectionItem{
			ID: int64(collectionID),
		}, nil)

		mockDrop := raindrop.Drop{
			ID:      1,
			Title:   "Image 1",
			Note:    "Description for image 1",
			Tags:    []string{"tag1", "tag2"},
			Created: time.Now(),
			Link:    "https://example.com/image1.jpg",
			Cover:   "https://placehold.co/600x400/000000/FFFFFF/png",
		}

		mockDrops := &raindrop.ImageDrops{
			Items: []raindrop.Drop{
				mockDrop,
			},
		}
		rdClient.On("GetImagesDropsFromCollection", mock.Anything, collectionID, 0).Return(mockDrops, nil)

		err := dl.DownloadCollection(context.Background(), collectionID, outputDir, true)
		require.NoError(t, err)

		rdClient.AssertExpectations(t)

		// Check if the file was downloaded
		downloadedFilePath := filepath.Join(outputDir, mockDrop.GetName()+".png")
		downloadedFileInfoPath := filepath.Join(outputDir, mockDrop.GetName()+".info.json")

		_, err = os.Stat(downloadedFilePath)
		assert.NoError(t, err)

		_, err = os.Stat(downloadedFileInfoPath)
		assert.NoError(t, err)

		// Check contents of the info file
		infoFile, err := os.ReadFile(downloadedFileInfoPath)
		require.NoError(t, err)

		var infoFileContent downloader.InfoFile
		err = json.NewDecoder(bytes.NewReader(infoFile)).Decode(&infoFileContent)

		require.NoError(t, err)

		assert.Equal(t, mockDrop.Title, infoFileContent.Title)
		assert.Equal(t, mockDrop.GetDescription(), infoFileContent.Description)
		assert.Equal(t, mockDrop.Tags, infoFileContent.Tags)
		assert.NotEmpty(t, infoFileContent.CreatedAt)
		assert.Equal(t, mockDrop.Link, infoFileContent.OriginalURL)
	})

	t.Run("Success Without Info file", func(t *testing.T) {
		t.Parallel()

		dl, rdClient := setupTestDownloader(t)
		// create tmp dir
		outputDir := generateTmpDir(t)
		defer os.RemoveAll(outputDir)

		collectionID := 123

		rdClient.On("GetCollectionByID", mock.Anything, collectionID).Return(&raindrop.CollectionItem{
			ID: int64(collectionID),
		}, nil)

		mockDrop := raindrop.Drop{
			ID:      1,
			Title:   "Image 1",
			Note:    "Description for image 1",
			Tags:    []string{"tag1", "tag2"},
			Created: time.Now(),
			Link:    "https://example.com/image1.jpg",
			Cover:   "https://placehold.co/600x400/000000/FFFFFF/png",
		}

		mockDrops := &raindrop.ImageDrops{
			Items: []raindrop.Drop{
				mockDrop,
			},
		}
		rdClient.On("GetImagesDropsFromCollection", mock.Anything, collectionID, 0).Return(mockDrops, nil)

		err := dl.DownloadCollection(context.Background(), collectionID, outputDir, false)
		require.NoError(t, err)

		rdClient.AssertExpectations(t)

		// Check if the file was downloaded
		downloadedFilePath := filepath.Join(outputDir, mockDrop.GetName()+".png")
		downloadedFileInfoPath := filepath.Join(outputDir, mockDrop.GetName()+".info.json")

		_, err = os.Stat(downloadedFilePath)
		assert.NoError(t, err)

		_, err = os.Stat(downloadedFileInfoPath)
		assert.True(t, os.IsNotExist(err))
	})
}
