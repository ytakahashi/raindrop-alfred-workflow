package raindrop

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func Test_NewClient(t *testing.T) {
	actual, err := NewClient("access-token")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	actualURL := actual.baseURL.String()
	expectedURL := "https://api.raindrop.io"
	if actualURL != expectedURL {
		t.Errorf("assert failed. expect:%s actual:%s", expectedURL, actualURL)
	}

	actualAccessToken := actual.accessToken
	expectedAccessToken := "access-token"
	if actualURL != expectedURL {
		t.Errorf("assert failed. expect:%s actual:%s", expectedAccessToken, actualAccessToken)
	}
}

func Test_GetRaindrops(t *testing.T) {
	// Given
	raindrop1 := Raindrop{
		Tags:  []string{"foo", "bar"},
		Title: "Test 1",
		Link:  "https://example.com/1",
	}
	raindrop2 := Raindrop{
		Tags:  []string{"baz"},
		Title: "Test 2",
		Link:  "https://example.com/2",
	}
	expected := Raindrops{
		Result: true,
		Items:  []Raindrop{raindrop1, raindrop2},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(expected)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}))

	defer ts.Close()

	// When
	sut := createTestClient(ts, t)

	// Then
	actual, err := sut.GetRaindrops("1")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if actual.Result != true {
		t.Error("actual.Result")
	}
	if len(actual.Items) != 2 {
		t.Errorf("Unexpected length: %d", len(actual.Items))
	}
	if !reflect.DeepEqual(actual.Items[0], raindrop1) {
		t.Errorf("Unexpected: %v, %v", actual.Items[0], raindrop1)
	}
	if !reflect.DeepEqual(actual.Items[1], raindrop2) {
		t.Errorf("Unexpected: %v, %v", actual.Items[1], raindrop2)
	}
}

func createTestClient(ts *httptest.Server, t *testing.T) Client {
	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	return Client{
		baseURL:     u,
		httpClient:  &http.Client{},
		accessToken: "",
	}
}
