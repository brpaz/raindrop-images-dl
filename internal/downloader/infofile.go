package downloader

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

type InfoFile struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	OriginalURL string    `json:"original_url"`
}

// createInfoFile generates a metadata file for a given Raindrop bookmark.
func createInfoFile(baseFilePath string, bookmark raindrop.Drop) error {
	infoFilePath := fmt.Sprintf("%s.info.json", baseFilePath)

	if fileExists(infoFilePath) {
		slog.Info("Info file already exists, skipping", "path", infoFilePath)
		return nil
	}

	infoFile, err := os.Create(infoFilePath)
	if err != nil {
		return err
	}
	defer infoFile.Close()

	info := InfoFile{
		Title:       bookmark.Title,
		Description: bookmark.GetDescription(),
		Tags:        bookmark.Tags,
		CreatedAt:   bookmark.Created,
		OriginalURL: bookmark.Link,
	}

	return json.NewEncoder(infoFile).Encode(info)
}
