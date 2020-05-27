package alfred

import (
	"encoding/json"

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

// ConvertToAlfredJSON converts a TimeStruct to JSON string
func ConvertToAlfredJSON(raindrops raindrop.Raindrops) string {
	items := []Item{}
	for _, r := range raindrops.Items {
		item := newItemFrom(r)
		items = append(items, item)
	}

	json, err := json.Marshal(Items{Items: items})
	if err != nil {
		return emptyItems
	}

	return string(json)
}

func newItemFrom(raindrop raindrop.Raindrop) Item {
	return Item{
		UID:      raindrop.Title,
		Title:    raindrop.Title,
		SubTitle: raindrop.Excerpt,
		Arg:      raindrop.Link,
	}
}
