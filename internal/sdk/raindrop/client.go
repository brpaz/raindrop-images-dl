package raindrop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	defaultBaseURL = "https://api.raindrop.io/rest/v1"
	itemsPerPage   = 50 // Number of items to retrieve per page. Max is 50 according to the Raindrop API docs
)

var (
	ErrMissingAPIKey  = errors.New("API key is required")
	ErrInvalidBaseURL = errors.New("Invalid base URL")
)

// Client is a client for the Raindrop API
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// Option defines a functional option type for configuring the Client
type Option func(*Client)

// NewClient creates a new Raindrop client with optional configurations
func NewClient(opts ...Option) (*Client, error) {
	// Set default values
	client := &Client{
		baseURL:    defaultBaseURL,
		httpClient: http.DefaultClient,
	}

	// Apply any options passed in
	for _, opt := range opts {
		opt(client)
	}

	if client.apiKey == "" {
		return nil, ErrMissingAPIKey
	}

	if client.baseURL == "" {
		return nil, ErrInvalidBaseURL
	}

	return client, nil
}

// WithAPIKey sets the API key for the client
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithHTTPClient allows providing a custom http.Client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// setAuthHeader sets the Authorization header with the API key
func (c *Client) setAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
}

// GetImagesDropsFromCollection retrieves all image drops from a collection
func (c *Client) GetImagesDropsFromCollection(ctx context.Context, collectionID int, page int) (*ImageDrops, error) {
	// Construct the URL for the Raindrop API
	url := fmt.Sprintf("%s/raindrops/%d", c.baseURL, collectionID)

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the API key for authorization
	c.setAuthHeader(req)

	// Add query parameters
	q := req.URL.Query()
	q.Add("perpage", fmt.Sprintf("%d", itemsPerPage))
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add(("search"), "type:image") // Filter for items of type "image"
	req.URL.RawQuery = q.Encode()

	// Send the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Check for a non-200 status code
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, respBody)
	}

	// Parse the response body
	var drops GetRaindropsResponse
	if err := json.NewDecoder(resp.Body).Decode(&drops); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// base on the page, count and itemsPerPage, determine if there are more items
	totalItems := drops.Count
	hasMore := true
	if totalItems == 0 || totalItems <= page*itemsPerPage {
		hasMore = false
	}

	return &ImageDrops{
		Items:   drops.Items,
		HasMore: hasMore,
	}, nil
}

func (c *Client) GetCollectionByID(ctx context.Context, collectionID int) (*CollectionItem, error) {
	url := fmt.Sprintf("%s/collection/%d", c.baseURL, collectionID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setAuthHeader(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, respBody)
	}

	var collection GetCollectionResponse
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &collection.Item, nil
}
