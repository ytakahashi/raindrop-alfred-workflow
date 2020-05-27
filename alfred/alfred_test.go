package alfred

import (
	"testing"

	"github.com/ytakahashi/raindrop-alfred-workflow/raindrop"
)

func Test_ConvertToAlfredJSON(t *testing.T) {
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

	actual := ConvertToAlfredJSON(raindrops)
	expected := `{"items":[{"uid":"Test 1","title":"Test 1","subtitle":"excerpt 1","arg":"https://example.com/1"},{"uid":"Test 2","title":"Test 2","subtitle":"excerpt 2","arg":"https://example.com/2"}]}`
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}
