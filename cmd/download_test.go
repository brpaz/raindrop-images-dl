package cmd_test

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/brpaz/raindrop-images-dl/cmd"
)

func setupTestDownloadCmd() *cobra.Command {
	return cmd.NewDownloadCmd()
}

func resetEnv(t *testing.T) {
	t.Helper()
	t.Setenv("RAINDROP_COLLECTION", "")
	t.Setenv("OUTPUT_DIR", "")
	t.Setenv("RAINDROP_API_KEY", "")
}

func TestNewDownloadCmd(t *testing.T) {
	t.Parallel()

	downloadCmd := setupTestDownloadCmd()

	assert.IsType(t, &cobra.Command{}, downloadCmd)
	assert.Equal(t, "download", downloadCmd.Use)
}

func TestDownloadPreFn(t *testing.T) {
	t.Run("sets flags from environment variables", func(t *testing.T) {
		t.Setenv("RAINDROP_COLLECTION", "123")
		t.Setenv("OUTPUT_DIR", "/some/path")
		t.Setenv("RAINDROP_API_KEY", "test-key")

		downloadCmd := setupTestDownloadCmd()

		err := downloadCmd.PreRunE(downloadCmd, []string{})
		require.NoError(t, err)

		collection, err := downloadCmd.Flags().GetInt("collection")
		require.NoError(t, err)
		assert.Equal(t, 123, collection)

		output, err := downloadCmd.Flags().GetString("output")
		require.NoError(t, err)
		assert.Equal(t, "/some/path", output)

		apiKey, err := downloadCmd.Flags().GetString("api-key")
		require.NoError(t, err)
		assert.Equal(t, "test-key", apiKey)
	})
}

func TestDownloadExecute(t *testing.T) {
	t.Run("returns errors when required flags are not set", func(t *testing.T) {
		resetEnv(t)
		downloadCmd := setupTestDownloadCmd()
		err := downloadCmd.ExecuteContext(context.Background())
		require.Error(t, err)

		assert.Contains(t, err.Error(), "required flag(s) \"api-key\", \"collection\", \"output\" not set")
	})
}
