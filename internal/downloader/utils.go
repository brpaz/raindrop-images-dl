package downloader

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

// MIME type to file extension map
var mimeMap = map[string]string{
	"image/jpeg":               ".jpg",
	"image/png":                ".png",
	"image/gif":                ".gif",
	"image/bmp":                ".bmp",
	"image/webp":               ".webp",
	"image/svg+xml":            ".svg",
	"application/pdf":          ".pdf",
	"text/plain":               ".txt",
	"application/octet-stream": ".bin",
}

// DownloadFile downloads a file from a URL and saves it to the destination path.
func downloadFile(url, dest string) error {
	resp, err := http.Get(url) // #nosec
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	extension, supported := mimeMap[contentType]
	if !supported {
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	dest += extension

	if fileExists(dest) {
		slog.Info("File already exists, skipping", "path", dest)
		return nil
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// dirExists checks if a directory exists.
func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// fileExists checks if a file already exists at the specified path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ensureDir ensures that a directory exists, creating it if necessary.
func ensureDir(dir string) error {
	if !dirExists(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}
