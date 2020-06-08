package workflow

import (
	"errors"
	"flag"
	"fmt"
	"os"
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
	s, e := validateAndDetermineScenario()
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	return s
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
	flag.BoolVar(&raindrops, "raindrops", false, "boolean flag whether to call Get raindrops api or not. defaults to false")
	flag.BoolVar(&collections, "collections", false, "boolean flag whether to call Get collections api or not. defaults to false")
	flag.BoolVar(&tags, "tags", false, "boolean flag whether to call Get tags api or not. defaults to false")
	flag.StringVar(&collectionID, "collectionId", "", "collection id used to call raindrop api")
	flag.StringVar(&tag, "tag", "", "tag value used to call raindrop api")
	flag.Parse()
}

func validateAndDetermineScenario() (Scenario, error) {
	if accessToken == "" {
		return "", errors.New("accessToken is required")
	}

	if raindrops && !tags && !collections {
		if collectionID != "" && tag != "" {
			return "", errors.New("can't specify '-collectionId' and '-tag' at the same time")
		} else if collectionID != "" {
			return GetRaindropsByCollectionID, nil
		} else if tag != "" {
			return GetRaindropsByTag, nil
		} else {
			return GetRaindrops, nil
		}
	}

	if !raindrops && tags && !collections {
		return GetTags, nil
	}

	if !raindrops && !tags && collections {
		return GetCollections, nil
	}

	return "", errors.New("one of '-raindrops' or '-tags' or '-collections' should be specified")
}
