package workflow

import (
	"flag"
	"fmt"
)

// Scenario represents a scenario
type Scenario string

const (
	// GetRaindrops scenario calls Get Raindrops API
	GetRaindrops = Scenario("GetRaindrops")

	// GetRaindropsByCollectionID scenario calls Get Raindrops API using given collection id
	GetRaindropsByCollectionID = Scenario("GetRaindropsByCollectionID")

	// GetRaindropsByTag scenario calls Get Raindrops API using given tag
	GetRaindropsByTag = Scenario("GetTaggedRaindrops")

	// GetCollections scenario calls Get Collections API
	GetCollections = Scenario("GetCollections")

	// GetTags scenario calls Get Tags API
	GetTags = Scenario("GetTags")
)

// NewScenario returns scenario
func NewScenario() Scenario {
	parseFlag()
	return validateAndDetermineScenario()
}

// NewScenarioRunner returns scenario runner
func NewScenarioRunner(target Scenario) ScenarioRunner {
	switch target {
	case GetRaindrops:
		return getRaindropScenarioRunner{
			accessToken: accessToken,
		}

	case GetRaindropsByCollectionID:
		return getRaindropsByCollectionIDScenarioRunner{
			accessToken:  accessToken,
			collectionID: collectionID,
		}

	case GetRaindropsByTag:
		return getRaindropsByTagScenarioRunner{
			accessToken: accessToken,
			tag:         tag,
		}

	case GetCollections:
		return getCollectionsScenarioRunner{
			accessToken: accessToken,
		}

	case GetTags:
		return getTagsScenarioRunner{
			accessToken: accessToken,
		}

	default:
		fmt.Printf("Undefined Scenario (%s)", target)
		return noopRunner{}
	}
}

var (
	accessToken  string
	raindrops    bool
	collections  bool
	tags         bool
	collectionID string
	tag          string
)

func parseFlag() {
	flag.StringVar(&accessToken, "accessToken", "", "access token to call raindrop api")
	flag.BoolVar(&raindrops, "raindrops", false, "calls Get raindrops api or not")
	flag.BoolVar(&collections, "collections", false, "calls Get collections api or not")
	flag.BoolVar(&tags, "tags", false, "calls Get tags api or not")
	flag.StringVar(&collectionID, "collectionId", "", "collection id used to call raindrop api")
	flag.StringVar(&tag, "tag", "", "tag value used to call raindrop api")
	flag.Parse()
}

func validateAndDetermineScenario() Scenario {
	return GetRaindropsByTag
}
