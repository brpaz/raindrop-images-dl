package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brpaz/raindrop-images-dl/cmd"
)

func TestNewApp(t *testing.T) {
	t.Parallel()

	app := cmd.NewApp()

	assert.IsType(t, &cmd.App{}, app)
}
