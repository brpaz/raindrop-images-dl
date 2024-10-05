package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/brpaz/raindrop-images-dl/internal/downloader"
	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

func downloadPreFn(cmd *cobra.Command, args []string) error {
	// Set the flags from environment variables if not provided
	collection, _ := cmd.Flags().GetInt("collection")
	if collection == 0 {
		envCollection := os.Getenv("RAINDROP_COLLECTION")
		if envCollection != "" {
			_ = cmd.Flags().Set("collection", envCollection)
		}
	}

	output, _ := cmd.Flags().GetString("output")
	if output == "" {
		envOutput := os.Getenv("OUTPUT_DIR")
		if envOutput != "" {
			_ = cmd.Flags().Set("output", envOutput)
		}
	}

	apiKey, _ := cmd.Flags().GetString("api-key")
	if apiKey == "" {
		envApiKey := os.Getenv("RAINDROP_API_KEY")

		if envApiKey != "" {
			_ = cmd.Flags().Set("api-key", envApiKey)
		}
	}

	return nil
}

func downloadRunFn(cmd *cobra.Command, args []string) error {
	collection, _ := cmd.Flags().GetInt("collection")
	output, _ := cmd.Flags().GetString("output")

	apiKey, _ := cmd.Flags().GetString("api-key")

	raindropClient, err := raindrop.NewClient(raindrop.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("failed to initialize Raindrop.io client: %w", err)
	}

	dl, err := downloader.NewDownloader(downloader.WithRaindropClient(raindropClient))
	if err != nil {
		return fmt.Errorf("failed to initialize downloader: %w", err)
	}

	err = dl.DownloadCollection(cmd.Context(), collection, output)
	if err != nil {
		return fmt.Errorf("failed to download collection: %w", err)
	}

	return nil
}

func NewDownloadCmd() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:     "download",
		Short:   "Download images from Raindrop.io collections",
		PreRunE: downloadPreFn,
		RunE:    downloadRunFn,
	}

	downloadCmd.Flags().IntP("collection", "c", 0, "The collection ID to download images from")
	downloadCmd.Flags().StringP("output", "o", "", "The output directory to save the images")
	downloadCmd.Flags().StringP("api-key", "k", "", "The Raindrop.io API key")

	_ = downloadCmd.MarkFlagRequired("api-key")
	_ = downloadCmd.MarkFlagRequired("collection")
	_ = downloadCmd.MarkFlagRequired("output")

	return downloadCmd
}
