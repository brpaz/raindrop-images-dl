<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# raindrop

```go
import "github.com/brpaz/raindrop-images-dl/internal/sdk/raindrop"
```

package raindrop provides an SDK to interact with the Raindrop API. An API key is required to use this SDK. Check the official \[Raindrop API documentation\]\(https://developer.raindrop.io/v1/authentication/token\) for more information. To simplify the usage a "test token" is used instead a full OAuth2 flow. Example Usage:

```
client, err := raindrop.NewClient(raindrop.WithAPIKey("test-api-key"))
if err != nil {
	log.Fatalf("error creating client: %v", err)
}
```

## Index

- [Variables](<#variables>)
- [type Access](<#Access>)
- [type Client](<#Client>)
  - [func NewClient\(opts ...Option\) \(\*Client, error\)](<#NewClient>)
  - [func \(c \*Client\) GetCollectionByID\(ctx context.Context, collectionID int\) \(\*CollectionItem, error\)](<#Client.GetCollectionByID>)
  - [func \(c \*Client\) GetImagesDropsFromCollection\(ctx context.Context, collectionID int, page int\) \(\*ImageDrops, error\)](<#Client.GetImagesDropsFromCollection>)
- [type CollectionItem](<#CollectionItem>)
- [type CollectionRef](<#CollectionRef>)
- [type Creator](<#Creator>)
- [type Drop](<#Drop>)
  - [func \(d Drop\) GetDescription\(\) string](<#Drop.GetDescription>)
  - [func \(d Drop\) GetFileLink\(\) string](<#Drop.GetFileLink>)
  - [func \(d Drop\) GetName\(\) string](<#Drop.GetName>)
- [type GetCollectionResponse](<#GetCollectionResponse>)
- [type GetRaindropsResponse](<#GetRaindropsResponse>)
- [type ImageDrops](<#ImageDrops>)
- [type Option](<#Option>)
  - [func WithAPIKey\(apiKey string\) Option](<#WithAPIKey>)
  - [func WithBaseURL\(baseURL string\) Option](<#WithBaseURL>)
  - [func WithHTTPClient\(httpClient \*http.Client\) Option](<#WithHTTPClient>)
- [type Ref](<#Ref>)


## Variables

<a name="ErrMissingAPIKey"></a>

```go
var (
    ErrMissingAPIKey  = errors.New("API key is required")
    ErrInvalidBaseURL = errors.New("Invalid base URL")
)
```

<a name="Access"></a>
## type Access



```go
type Access struct {
    For       int  `json:"for"`
    Level     int  `json:"level"`
    Root      bool `json:"root"`
    Draggable bool `json:"draggable"`
}
```

<a name="Client"></a>
## type Client

Client is a client for the Raindrop API

```go
type Client struct {
    // contains filtered or unexported fields
}
```

<a name="NewClient"></a>
### func NewClient

```go
func NewClient(opts ...Option) (*Client, error)
```

NewClient creates a new Raindrop client with optional configurations

<a name="Client.GetCollectionByID"></a>
### func \(\*Client\) GetCollectionByID

```go
func (c *Client) GetCollectionByID(ctx context.Context, collectionID int) (*CollectionItem, error)
```



<a name="Client.GetImagesDropsFromCollection"></a>
### func \(\*Client\) GetImagesDropsFromCollection

```go
func (c *Client) GetImagesDropsFromCollection(ctx context.Context, collectionID int, page int) (*ImageDrops, error)
```

GetImagesDropsFromCollection retrieves all image drops from a collection

<a name="CollectionItem"></a>
## type CollectionItem



```go
type CollectionItem struct {
    ID          int64    `json:"_id"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    User        Ref      `json:"user"`
    Parent      Ref      `json:"parent"`
    Public      bool     `json:"public"`
    View        string   `json:"view"`
    Count       int      `json:"count"`
    Cover       []string `json:"cover"`
    Sort        int      `json:"sort"`
    CreatorRef  Creator  `json:"creatorRef"`
    LastAction  string   `json:"lastAction"`
    Created     string   `json:"created"`
    LastUpdate  string   `json:"lastUpdate"`
    Slug        string   `json:"slug"`
    Color       string   `json:"color"`
    Access      Access   `json:"access"`
    Author      bool     `json:"author"`
}
```

<a name="CollectionRef"></a>
## type CollectionRef



```go
type CollectionRef struct {
    Ref string `json:"$ref"`
    ID  int64  `json:"$id"`
    Oid int64  `json:"oid"`
}
```

<a name="Creator"></a>
## type Creator



```go
type Creator struct {
    ID    int    `json:"_id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

<a name="Drop"></a>
## type Drop



```go
type Drop struct {
    ID           int64         `json:"_id"`
    Link         string        `json:"link"`
    Title        string        `json:"title"`
    Excerpt      string        `json:"excerpt"`
    Note         string        `json:"note"`
    Type         string        `json:"type"`
    Cover        string        `json:"cover"`
    Media        []string      `json:"media"`
    Tags         []string      `json:"tags"`
    Created      time.Time     `json:"created"`
    Collection   CollectionRef `json:"collection"`
    LastUpdate   string        `json:"lastUpdate"`
    CollectionID int64         `json:"collectionId"`
}
```

<a name="Drop.GetDescription"></a>
### func \(Drop\) GetDescription

```go
func (d Drop) GetDescription() string
```



<a name="Drop.GetFileLink"></a>
### func \(Drop\) GetFileLink

```go
func (d Drop) GetFileLink() string
```



<a name="Drop.GetName"></a>
### func \(Drop\) GetName

```go
func (d Drop) GetName() string
```



<a name="GetCollectionResponse"></a>
## type GetCollectionResponse



```go
type GetCollectionResponse struct {
    Result bool           `json:"result"`
    Item   CollectionItem `json:"item"`
}
```

<a name="GetRaindropsResponse"></a>
## type GetRaindropsResponse



```go
type GetRaindropsResponse struct {
    Result bool   `json:"result"`
    Items  []Drop `json:"items"`
    Count  int    `json:"count"`
}
```

<a name="ImageDrops"></a>
## type ImageDrops



```go
type ImageDrops struct {
    Items   []Drop
    HasMore bool
}
```

<a name="Option"></a>
## type Option

Option defines a functional option type for configuring the Client

```go
type Option func(*Client)
```

<a name="WithAPIKey"></a>
### func WithAPIKey

```go
func WithAPIKey(apiKey string) Option
```

WithAPIKey sets the API key for the client

<a name="WithBaseURL"></a>
### func WithBaseURL

```go
func WithBaseURL(baseURL string) Option
```



<a name="WithHTTPClient"></a>
### func WithHTTPClient

```go
func WithHTTPClient(httpClient *http.Client) Option
```

WithHTTPClient allows providing a custom http.Client

<a name="Ref"></a>
## type Ref



```go
type Ref struct {
    Ref string `json:"$ref"`
    ID  int    `json:"$id"`
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
