package alfred

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/ytakahashi/raindrop-alfred-workflow/pkg/raindrop"
)

// Mod controls how the modifier keys react
type Mod struct {
	Valid    bool   `json:"valid"`
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
}

// Mods is modifier commands
type Mods struct {
	Cmd Mod `json:"cmd"`
}

// Item is an alfred item (https://www.alfredapp.com/help/workflows/inputs/script-filter/json/)
type Item struct {
	UID          string `json:"uid"`
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Match        string `json:"match"`
	QuickLookURL string `json:"quicklookurl"`
	Mods         Mods   `json:"mods"`
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

// ConvertToAlfredJSONFromTags creates json string from Tags
func ConvertToAlfredJSONFromTags(tags raindrop.Tags) string {
	items := []Item{}
	for _, t := range tags.Items {
		item := newItemFromTag(t)
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
		UID:      fmt.Sprint(collection.ID),
		Title:    collection.Title,
		SubTitle: fmt.Sprintf("%d items", collection.Count),
		Arg:      fmt.Sprint(collection.ID),
		Match:    collection.Title,
	}
}

func newItemFromRaindrop(raindrop raindrop.Raindrop) Item {

	u, err := url.Parse(raindrop.Link)
	if err != nil {
		return Item{
			UID:      raindrop.Title,
			Title:    raindrop.Title,
			SubTitle: raindrop.Excerpt,
			Arg:      raindrop.Link,
			Match:    raindrop.Title,
		}
	}

	host := strings.Replace(u.Hostname(), ".", " ", -1)
	path := strings.Replace(u.Path, "/", " ", -1)
	return Item{
		UID:          raindrop.Title,
		Title:        raindrop.Title,
		SubTitle:     raindrop.Excerpt,
		Arg:          raindrop.Link,
		Match:        fmt.Sprintf("%s %s %s", raindrop.Title, host, path),
		QuickLookURL: raindrop.Link,
		Mods: Mods{
			Cmd: Mod{
				Valid:    true,
				Arg:      raindrop.Link,
				Subtitle: fmt.Sprintf("Cmd + c to copy %s, Cmd + y to Quick Look", raindrop.Link),
			},
		},
	}
}

func newItemFromTag(tag raindrop.Tag) Item {
	return Item{
		UID:      tag.ID,
		Title:    tag.ID,
		SubTitle: fmt.Sprintf("%d items", tag.Count),
		Arg:      tag.ID,
		Match:    tag.ID,
	}
}
