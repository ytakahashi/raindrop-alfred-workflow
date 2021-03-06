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
	actual, err := sut.GetRaindrops("1", 50)
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

func Test_GetCollections(t *testing.T) {
	// Given
	collection1 := Collection{
		ID:    1,
		Title: "Test 1",
	}
	collection2 := Collection{
		ID:    2,
		Title: "Test 2",
	}
	expected := Collections{
		Result: true,
		Items:  []Collection{collection1, collection2},
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
	actual, err := sut.GetCollections()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if actual.Result != true {
		t.Error("actual.Result")
	}
	if len(actual.Items) != 2 {
		t.Errorf("Unexpected length: %d", len(actual.Items))
	}
	if !reflect.DeepEqual(actual.Items[0], collection1) {
		t.Errorf("Unexpected: %v, %v", actual.Items[0], collection1)
	}
	if !reflect.DeepEqual(actual.Items[1], collection2) {
		t.Errorf("Unexpected: %v, %v", actual.Items[1], collection2)
	}
}

func Test_GetTags(t *testing.T) {
	// Given
	tag1 := Tag{
		ID:    "tag 1",
		Count: 10,
	}
	tag2 := Tag{
		ID:    "tag 2",
		Count: 100,
	}
	expected := Tags{
		Result: true,
		Items:  []Tag{tag1, tag2},
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
	actual, err := sut.GetTags()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if actual.Result != true {
		t.Error("actual.Result")
	}
	if len(actual.Items) != 2 {
		t.Errorf("Unexpected length: %d", len(actual.Items))
	}
	if !reflect.DeepEqual(actual.Items[0], tag1) {
		t.Errorf("Unexpected: %v, %v", actual.Items[0], tag1)
	}
	if !reflect.DeepEqual(actual.Items[1], tag2) {
		t.Errorf("Unexpected: %v, %v", actual.Items[1], tag2)
	}
}

func Test_GetTaggedRaindrops(t *testing.T) {
	// Given
	raindrop1 := Raindrop{
		Tags:  []string{"foo", "bar"},
		Title: "Test 1",
		Link:  "https://example.com/1",
	}
	expected := Raindrops{
		Result: true,
		Items:  []Raindrop{raindrop1},
	}
	tag := "tag 1"
	expectedQuery := `[{"key":"tag","val":"tag 1"}]`
	var actualQuery string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(expected)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		actualQuery = r.URL.Query().Get("search")

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}))

	defer ts.Close()

	// When
	sut := createTestClient(ts, t)

	// Then
	actual, err := sut.GetTaggedRaindrops(tag)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if actualQuery != expectedQuery {
		t.Errorf("Unexpected: %v, %v", actualQuery, expectedQuery)
	}
	if actual.Result != true {
		t.Error("actual.Result")
	}
	if len(actual.Items) != 1 {
		t.Errorf("Unexpected length: %d", len(actual.Items))
	}
	if !reflect.DeepEqual(actual.Items[0], raindrop1) {
		t.Errorf("Unexpected: %v, %v", actual.Items[0], raindrop1)
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
