package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/brpaz/raindrop-images-dl/cmd"
)

func setupTestDownloadCmd() *cobra.Command {
	return cmd.NewDownloadCmd()
}

func TestNewDownloadCmd(t *testing.T) {
	t.Parallel()

	downloadCmd := setupTestDownloadCmd()

	assert.IsType(t, &cobra.Command{}, downloadCmd)
	assert.Equal(t, "download", downloadCmd.Use)
}

func TestDownloadCmdFlags(t *testing.T) {
}
