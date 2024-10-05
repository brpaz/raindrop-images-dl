package raindrop

import (
	"strings"
	"time"
)

type ImageDrops struct {
	Items   []Drop
	HasMore bool
}

type GetRaindropsResponse struct {
	Result bool   `json:"result"`
	Items  []Drop `json:"items"`
	Count  int    `json:"count"`
}

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

func (d Drop) GetFileLink() string {
	return d.Cover
}

func (d Drop) GetName() string {
	return strings.ReplaceAll(d.Title, " ", "_")
}

func (d Drop) GetDescription() string {
	return d.Note
}

type CollectionRef struct {
	Ref string `json:"$ref"`
	ID  int64  `json:"$id"`
	Oid int64  `json:"oid"`
}

type GetCollectionResponse struct {
	Result bool           `json:"result"`
	Item   CollectionItem `json:"item"`
}

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

type Ref struct {
	Ref string `json:"$ref"`
	ID  int    `json:"$id"`
}

type Creator struct {
	ID    int    `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Access struct {
	For       int  `json:"for"`
	Level     int  `json:"level"`
	Root      bool `json:"root"`
	Draggable bool `json:"draggable"`
}
