package raindrop_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

var mockDrop = raindrop.Drop{
	ID:         123,
	Title:      "Test Title",
	Note:       "Test Note",
	Link:       "https://test.com",
	Cover:      "https://test.com/cover.jpg",
	Tags:       []string{"tag1", "tag2"},
	LastUpdate: "2021-01-01T00:00:00Z",
	Collection: raindrop.CollectionRef{
		ID:  456,
		Ref: "Test Collection",
	},
}

func TestDrop_GetDescription(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Test Note", mockDrop.GetDescription())
}

func TestDrop_GetName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Test_Title", mockDrop.GetName())
}

func TestDrop_GetFileLink(t *testing.T) {
	t.Parallel()

	assert.Equal(t, mockDrop.Cover, mockDrop.GetFileLink())
}
