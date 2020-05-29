package alfred

import (
	"encoding/json"
	"fmt"

	"github.com/ytakahashi/raindrop-alfred-workflow/raindrop"
)

// Item is an alfred item (https://www.alfredapp.com/help/workflows/inputs/script-filter/json/)
type Item struct {
	UID      string `json:"uid"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Items ia an array of items
type Items struct {
	Items []Item `json:"items"`
}

var emptyItems = "{\"items\":[]}"

// ConvertToAlfredJSONFromCollections creates json string from Collections
func ConvertToAlfredJSONFromCollections(collections raindrop.Collections) string {
	items := []Item{}
	for _, r := range collections.Items {
		item := newItemFromCollection(r)
		items = append(items, item)
	}

	json, err := json.Marshal(Items{Items: items})
	if err != nil {
		return emptyItems
	}

	return string(json)
}

// ConvertToAlfredJSONFromRaindrops creates json string from Raindrops
func ConvertToAlfredJSONFromRaindrops(raindrops raindrop.Raindrops) string {
	items := []Item{}
	for _, r := range raindrops.Items {
		item := newItemFromRaindrop(r)
		items = append(items, item)
	}

	json, err := json.Marshal(Items{Items: items})
	if err != nil {
		return emptyItems
	}

	return string(json)
}

func newItemFromCollection(collection raindrop.Collection) Item {
	return Item{
		UID:   fmt.Sprint(collection.ID),
		Title: collection.Title,
		Arg:   fmt.Sprint(collection.ID),
	}
}

func newItemFromRaindrop(raindrop raindrop.Raindrop) Item {
	return Item{
		UID:      raindrop.Title,
		Title:    raindrop.Title,
		SubTitle: raindrop.Excerpt,
		Arg:      raindrop.Link,
	}
}
