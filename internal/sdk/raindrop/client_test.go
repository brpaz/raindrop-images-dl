package raindrop_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
)

const (
	testAPIKey = "test-api-key"
)

// Helper function to load testdata
func loadTestData(t *testing.T, filePath string) []byte {
	t.Helper()
	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	return data
}

func setupTestClient(t *testing.T, server *httptest.Server) *raindrop.Client {
	t.Helper()
	client, err := raindrop.NewClient(
		raindrop.WithAPIKey(testAPIKey),
		raindrop.WithHTTPClient(server.Client()),
		raindrop.WithBaseURL(server.URL),
	)

	require.NoError(t, err, "error creating raindrop client")

	return client
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("WithoutAPIKey_ReturnsError", func(t *testing.T) {
		t.Parallel()

		client, err := raindrop.NewClient()

		assert.Nil(t, client)
		assert.ErrorIs(t, err, raindrop.ErrMissingAPIKey)
	})

	t.Run("WithAPIKey_Empty_ReturnsError", func(t *testing.T) {
		t.Parallel()

		client, err := raindrop.NewClient(raindrop.WithAPIKey(""))

		assert.Nil(t, client)
		assert.ErrorIs(t, err, raindrop.ErrMissingAPIKey)
	})

	t.Run("WithAPIKey", func(t *testing.T) {
		t.Parallel()

		client, err := raindrop.NewClient(raindrop.WithAPIKey(testAPIKey))

		require.NoError(t, err)

		assert.IsType(t, &raindrop.Client{}, client)
	})

	t.Run("WithHTTPClient", func(t *testing.T) {
		t.Parallel()

		client, err := raindrop.NewClient(
			raindrop.WithAPIKey(testAPIKey),
			raindrop.WithHTTPClient(&http.Client{}),
		)

		require.NoError(t, err)

		assert.IsType(t, &raindrop.Client{}, client)
	})

	t.Run("WithBaseURL", func(t *testing.T) {
		t.Parallel()

		client, err := raindrop.NewClient(
			raindrop.WithAPIKey(testAPIKey),
			raindrop.WithBaseURL("http://localhost"),
		)

		require.NoError(t, err)

		assert.IsType(t, &raindrop.Client{}, client)
	})

	t.Run("WithBaseURL_Empty_ReturnsError", func(t *testing.T) {
		t.Parallel()

		client, err := raindrop.NewClient(
			raindrop.WithAPIKey(testAPIKey),
			raindrop.WithBaseURL(""),
		)

		assert.Nil(t, client)
		assert.ErrorIs(t, err, raindrop.ErrInvalidBaseURL)
	})
}

func TestGetImagesDropsFromCollection(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		mockData := loadTestData(t, "testdata/get_raindrops_response_success.json")

		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/raindrops/123", r.URL.Path) // Check collection ID in URL
			assert.Equal(t, fmt.Sprintf("Bearer %s", testAPIKey), r.Header.Get("Authorization"))
			assert.Equal(t, "type:image", r.URL.Query().Get("search"))
			assert.Equal(t, "1", r.URL.Query().Get("page"))
			assert.Equal(t, "50", r.URL.Query().Get("perpage"))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData) // Serve the mock data from file
		}))
		defer server.Close()

		// Create a client with the test server URL
		client := setupTestClient(t, server)

		// Call the method
		imageDrops, err := client.GetImagesDropsFromCollection(context.Background(), 123, 1)
		require.NoError(t, err)

		// Check the returned result
		assert.Len(t, imageDrops.Items, 2)
		assert.Equal(t, "Image 1", imageDrops.Items[0].Title)
		assert.Equal(t, "Image 2", imageDrops.Items[1].Title)
		assert.False(t, imageDrops.HasMore) // Since the total count matches the returned items
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		t.Parallel()
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		defer server.Close()

		// Create a client with the test server URL
		client := setupTestClient(t, server)

		// Call the method expecting an error due to non-200 status code
		_, err := client.GetImagesDropsFromCollection(context.Background(), 123, 1)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected status code: 500")
	})

	t.Run("Invalid JSON Response", func(t *testing.T) {
		t.Parallel()

		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{invalid-json}`)) // Invalid JSON
		}))
		defer server.Close()

		// Create a client with the test server URL
		client := setupTestClient(t, server)

		// Call the method expecting a JSON decoding error
		_, err := client.GetImagesDropsFromCollection(context.Background(), 123, 1)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}

func TestGetCollectionByID(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		mockData := loadTestData(t, "testdata/get_collection_response_success.json")

		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/collection/47919682", r.URL.Path) // Check collection ID in URL
			assert.Equal(t, fmt.Sprintf("Bearer %s", testAPIKey), r.Header.Get("Authorization"))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockData) // Serve the mock data from file
		}))
		defer server.Close()

		// Create a client with the test server URL
		client := setupTestClient(t, server)

		// Call the method
		collection, err := client.GetCollectionByID(context.Background(), 47919682)
		require.NoError(t, err)

		// Check the returned result
		assert.Equal(t, "Images", collection.Title)
		assert.Equal(t, int64(47919682), collection.ID)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		t.Parallel()
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		defer server.Close()

		// Create a client with the test server URL
		client := setupTestClient(t, server)

		// Call the method expecting an error due to non-200 status code
		_, err := client.GetCollectionByID(context.Background(), 123)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected status code: 500")
	})

	t.Run("Invalid JSON Response", func(t *testing.T) {
		t.Parallel()

		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{invalid-json}`)) // Invalid JSON
		}))
		defer server.Close()

		// Create a client with the test server URL
		client := setupTestClient(t, server)

		// Call the method expecting a JSON decoding error
		_, err := client.GetCollectionByID(context.Background(), 123)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode response")
	})
}
