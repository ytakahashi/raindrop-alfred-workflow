package main

import (
	"github.com/ytakahashi/raindrop-alfred-workflow/internal/workflow"
)

/*
Usage:
	(cmd) -accessToken {accessToken} -raindrops
	(cmd) -accessToken {accessToken} -raindrops -collectionId {collection id}
	(cmd) -accessToken {accessToken} -raindrops -tag {tag name}
	(cmd) -accessToken {accessToken} -collections
	(cmd) -accessToken {accessToken} -tags
*/
func main() {
	scenario := workflow.NewScenario()
	runner := workflow.NewScenarioRunner(scenario)
	runner.Run()
}
