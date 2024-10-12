package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd creates a new instance of the root command.
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "raindrop-images-dl",
		Short: "A CLI tool to download images from Raindrop.io collections",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
}
