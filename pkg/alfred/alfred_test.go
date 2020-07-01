package alfred

import (
	"testing"

	"github.com/ytakahashi/raindrop-alfred-workflow/pkg/raindrop"
)

func Test_ConvertToAlfredJSONFromRaindrop(t *testing.T) {
	raindrop1 := raindrop.Raindrop{
		Tags:    []string{"foo", "bar"},
		Title:   "Test 1",
		Excerpt: "excerpt 1",
		Link:    "https://example.com/1",
	}
	raindrop2 := raindrop.Raindrop{
		Tags:    []string{"baz"},
		Title:   "Test 2",
		Excerpt: "excerpt 2",
		Link:    "https://example.com/2",
	}
	raindrops := raindrop.Raindrops{
		Result: true,
		Items:  []raindrop.Raindrop{raindrop1, raindrop2},
	}

	// When
	actual := ConvertToAlfredJSONFromRaindrops(raindrops)

	// Then
	expected := `{"items":[{"uid":"Test 1","title":"Test 1","subtitle":"excerpt 1","arg":"https://example.com/1"},{"uid":"Test 2","title":"Test 2","subtitle":"excerpt 2","arg":"https://example.com/2"}]}`
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}

func Test_ConvertToAlfredJSONFromCollections(t *testing.T) {
	// Given
	collection1 := raindrop.Collection{
		ID:    1,
		Title: "Test 1",
	}
	collection2 := raindrop.Collection{
		ID:    2,
		Title: "Test 2",
	}
	collections := raindrop.Collections{
		Result: true,
		Items:  []raindrop.Collection{collection1, collection2},
	}

	// When
	actual := ConvertToAlfredJSONFromCollections(collections)

	// Then
	expected := `{"items":[{"uid":"1","title":"Test 1","subtitle":"","arg":"1"},{"uid":"2","title":"Test 2","subtitle":"","arg":"2"}]}`
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}

func Test_ConvertToAlfredJSONFromTags(t *testing.T) {
	// When
	tag1 := raindrop.Tag{
		ID:    "tag 1",
		Count: 123,
	}
	tag2 := raindrop.Tag{
		ID:    "tag 2",
		Count: 456,
	}
	tags := raindrop.Tags{
		Result: true,
		Items:  []raindrop.Tag{tag1, tag2},
	}

	// When
	actual := ConvertToAlfredJSONFromTags(tags)

	// Then
	expected := `{"items":[{"uid":"tag 1","title":"tag 1","subtitle":"","arg":"tag 1"},{"uid":"tag 2","title":"tag 2","subtitle":"","arg":"tag 2"}]}`
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}
