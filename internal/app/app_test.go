package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brpaz/raindrop-images-dl/internal/app"
)

func TestNewApp(t *testing.T) {
	t.Parallel()

	cmd := app.New()

	assert.IsType(t, &app.App{}, cmd)
}
