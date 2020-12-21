package workflow

import (
	"fmt"
	"os"

	"github.com/ytakahashi/raindrop-alfred-workflow/pkg/alfred"
	"github.com/ytakahashi/raindrop-alfred-workflow/pkg/raindrop"
)

// ScenarioRunner is a scenario runner
type ScenarioRunner interface {
	Run()
}

type getRaindropScenarioRunner struct {
	accessToken string
}

type getRaindropsByCollectionIDScenarioRunner struct {
	accessToken  string
	collectionID string
}

type getRaindropsByTagScenarioRunner struct {
	accessToken string
	tag         string
}

type getCollectionsScenarioRunner struct {
	accessToken string
}

type getTagsScenarioRunner struct {
	accessToken string
}

type noopRunner struct{}

// Run executes GetRaindrop Scenario
func (r getRaindropScenarioRunner) Run() {
	client, err := raindrop.NewClient(r.accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetRaindrops("0", 50)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonString := alfred.ConvertToAlfredJSONFromRaindrops(*res)
	fmt.Println(jsonString)
}

// Run executes getRaindropsByCollectionID Scenario
func (r getRaindropsByCollectionIDScenarioRunner) Run() {
	client, err := raindrop.NewClient(r.accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetRaindrops(r.collectionID, 50)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonString := alfred.ConvertToAlfredJSONFromRaindrops(*res)
	fmt.Println(jsonString)
}

// Run executes getCollections Scenario
func (r getCollectionsScenarioRunner) Run() {
	client, err := raindrop.NewClient(r.accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetCollections()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonString := alfred.ConvertToAlfredJSONFromCollections(*res)
	fmt.Println(jsonString)
}

// Run executes getTags Scenario
func (r getTagsScenarioRunner) Run() {
	client, err := raindrop.NewClient(r.accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetTags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonString := alfred.ConvertToAlfredJSONFromTags(*res)
	fmt.Println(jsonString)
}

// Run executes getRaindropsByTag Scenario
func (r getRaindropsByTagScenarioRunner) Run() {
	client, err := raindrop.NewClient(r.accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetTaggedRaindrops(r.tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonString := alfred.ConvertToAlfredJSONFromRaindrops(*res)
	fmt.Println(jsonString)
}

// Run does nothing
func (r noopRunner) Run() {
	// do nothing
}
