package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/brpaz/raindrop-images-dl/internal/downloader"
	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

const (
	FlagDownloadCollection = "collection"
	FlagDownloadOutput     = "output"
	FlagDownloadGenInfo    = "gen-info-json"
	FlagDownloadApiKey     = "api-key"
)

func downloadPreFn(cmd *cobra.Command, args []string) error {
	// Set the flags from environment variables if not provided
	collection, _ := cmd.Flags().GetInt(FlagDownloadCollection)
	if collection == 0 {
		envCollection := os.Getenv("RAINDROP_COLLECTION")
		if envCollection != "" {
			_ = cmd.Flags().Set(FlagDownloadCollection, envCollection)
		}
	}

	output, _ := cmd.Flags().GetString(FlagDownloadOutput)
	if output == "" {
		envOutput := os.Getenv("OUTPUT_DIR")
		if envOutput != "" {
			_ = cmd.Flags().Set(FlagDownloadOutput, envOutput)
		}
	}

	apiKey, _ := cmd.Flags().GetString(FlagDownloadApiKey)
	if apiKey == "" {
		envApiKey := os.Getenv("RAINDROP_API_KEY")

		if envApiKey != "" {
			_ = cmd.Flags().Set(FlagDownloadApiKey, envApiKey)
		}
	}

	isGenEnvInfoJsonSwet := cmd.Flags().Changed(FlagDownloadGenInfo)

	if !isGenEnvInfoJsonSwet {
		envGenInfoJson := os.Getenv("GEN_INFO_JSON")

		if envGenInfoJson != "" {
			_ = cmd.Flags().Set(FlagDownloadGenInfo, envGenInfoJson)
		}
	}

	return nil
}

func downloadRunFn(cmd *cobra.Command, args []string) error {
	collection, _ := cmd.Flags().GetInt(FlagDownloadCollection)
	output, _ := cmd.Flags().GetString(FlagDownloadOutput)
	apiKey, _ := cmd.Flags().GetString(FlagDownloadApiKey)
	infoJson, _ := cmd.Flags().GetBool(FlagDownloadGenInfo)

	raindropClient, err := raindrop.NewClient(raindrop.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("failed to initialize Raindrop.io client: %w", err)
	}

	dl, err := downloader.NewDownloader(downloader.WithRaindropClient(raindropClient))
	if err != nil {
		return fmt.Errorf("failed to initialize downloader: %w", err)
	}

	err = dl.DownloadCollection(cmd.Context(), collection, output, infoJson)
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

	downloadCmd.Flags().IntP(FlagDownloadCollection, "c", 0, "The collection ID to download images from")
	downloadCmd.Flags().StringP(FlagDownloadOutput, "o", "", "The output directory to save the images")
	downloadCmd.Flags().StringP(FlagDownloadApiKey, "k", "", "The Raindrop.io API key")
	downloadCmd.Flags().BoolP(FlagDownloadGenInfo, "i", true, "Generate a JSON file with the image metadata")

	_ = downloadCmd.MarkFlagRequired(FlagDownloadCollection)
	_ = downloadCmd.MarkFlagRequired(FlagDownloadApiKey)
	_ = downloadCmd.MarkFlagRequired(FlagDownloadOutput)

	return downloadCmd
}
