package workflow

import (
	"testing"
)

func Test_NewScenarioRunner_GetRaindrops(t *testing.T) {
	reset()
	accessToken = "access token"
	actual := NewScenarioRunner(GetRaindrops)
	expect := getRaindropScenarioRunner{
		accessToken: "access token",
	}

	if actual != expect {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_NewScenarioRunner_GetRaindropsByCollectionID(t *testing.T) {
	reset()
	accessToken = "access token"
	collectionID = "123"
	actual := NewScenarioRunner(GetRaindropsByCollectionID)
	expect := getRaindropsByCollectionIDScenarioRunner{
		accessToken:  "access token",
		collectionID: "123",
	}

	if actual != expect {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_NewScenarioRunner_GetRaindropsByTag(t *testing.T) {
	reset()
	accessToken = "access token"
	tag = "t a g"
	actual := NewScenarioRunner(GetRaindropsByTag)
	expect := getRaindropsByTagScenarioRunner{
		accessToken: "access token",
		tag:         "t a g",
	}

	if actual != expect {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_NewScenarioRunner_GetCollections(t *testing.T) {
	reset()
	accessToken = "access token"
	actual := NewScenarioRunner(GetCollections)
	expect := getCollectionsScenarioRunner{
		accessToken: "access token",
	}

	if actual != expect {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_NewScenarioRunner_GetTags(t *testing.T) {
	reset()
	accessToken = "access token"
	actual := NewScenarioRunner(GetTags)
	expect := getTagsScenarioRunner{
		accessToken: "access token",
	}

	if actual != expect {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_NewScenarioRunner_other(t *testing.T) {
	reset()
	actual := NewScenarioRunner("other")
	expect := noopRunner{}

	if actual != expect {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_validateAndDetermineScenario_accessToken(t *testing.T) {
	_, err := validateAndDetermineScenario()
	if err.Error() != "accessToken is required" {
		t.Errorf("Unexpected error. error:%s", err.Error())
	}
}

func Test_validateAndDetermineScenario_raindrops_error(t *testing.T) {
	reset()
	accessToken = "access-token"
	raindrops = true
	collectionID = "id"
	tag = "tag"

	_, err := validateAndDetermineScenario()
	if err.Error() != "can't specify '-collectionId' and '-tag' at the same time" {
		t.Errorf("Unexpected error. error:%s", err.Error())
	}
}

func Test_validateAndDetermineScenario_raindrops_normal_collection(t *testing.T) {
	reset()
	accessToken = "access-token"
	raindrops = true
	collectionID = "id"

	actual, _ := validateAndDetermineScenario()
	if actual != GetRaindropsByCollectionID {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_validateAndDetermineScenario_raindrops_normal_tag(t *testing.T) {
	reset()
	accessToken = "access-token"
	raindrops = true
	tag = "tag"

	actual, _ := validateAndDetermineScenario()
	if actual != GetRaindropsByTag {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_validateAndDetermineScenario_raindrops_normal(t *testing.T) {
	reset()
	accessToken = "access-token"
	raindrops = true

	actual, _ := validateAndDetermineScenario()
	if actual != GetRaindrops {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_validateAndDetermineScenario_tags(t *testing.T) {
	reset()
	accessToken = "access-token"
	tags = true

	actual, _ := validateAndDetermineScenario()
	if actual != GetTags {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_validateAndDetermineScenario_collections(t *testing.T) {
	reset()
	accessToken = "access-token"
	collections = true

	actual, _ := validateAndDetermineScenario()
	if actual != GetCollections {
		t.Errorf("Unexpected value. actual:%s", actual)
	}
}

func Test_validateAndDetermineScenario_error(t *testing.T) {
	reset()
	accessToken = "access-token"
	raindrops = true
	collections = true

	_, err := validateAndDetermineScenario()
	if err.Error() != "one of '-raindrops' or '-tags' or '-collections' should be specified" {
		t.Errorf("Unexpected error. error:%s", err.Error())
	}
}

func reset() {
	accessToken = ""
	raindrops = false
	collections = false
	tags = false
	collectionID = ""
	tag = ""
}
