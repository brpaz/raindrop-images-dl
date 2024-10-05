// package raindrop provides an SDK to interact with the Raindrop API.
// An API key is required to use this SDK. Check the official [Raindrop API documentation](https://developer.raindrop.io/v1/authentication/token) for more information.
// To simplify the usage a "test token" is used instead a full OAuth2 flow.
// Example Usage:
//
//	client, err := raindrop.NewClient(raindrop.WithAPIKey("test-api-key"))
//	if err != nil {
//		log.Fatalf("error creating client: %v", err)
//	}
package raindrop
