package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/brpaz/raindrop-images-dl/internal/cmd"
)

func TestRootCmd(t *testing.T) {
	t.Parallel()

	rootCmd := cmd.NewRootCmd()

	assert.IsType(t, &cobra.Command{}, rootCmd)
	assert.Equal(t, "raindrop-images-dl", rootCmd.Use)
}
