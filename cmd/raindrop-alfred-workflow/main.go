package main

import (
	"flag"
	"fmt"

	"github.com/ytakahashi/raindrop-alfred-workflow/alfred"
	"github.com/ytakahashi/raindrop-alfred-workflow/raindrop"
)

func main() {
	flag.Parse()
	accessToken := flag.Arg(0)
	api := flag.Arg(1)

	if api == "raindrops" {
		id := flag.Arg(2)
		getRaindrops(accessToken, id)
	} else if api == "collections" {
		getCollections(accessToken)
	}
}

func getRaindrops(accessToken, id string) {
	client, err := raindrop.NewClient(accessToken)

	res, err := client.GetRaindrops(id)
	if err != nil {
		fmt.Print(err)
	} else {
		jsonString := alfred.ConvertToAlfredJSONFromRaindrops(*res)
		fmt.Println(jsonString)
	}
}

func getCollections(accessToken string) {
	client, err := raindrop.NewClient(accessToken)

	res, err := client.GetCollections()
	if err != nil {
		fmt.Print(err)
	} else {
		jsonString := alfred.ConvertToAlfredJSONFromCollections(*res)
		fmt.Println(jsonString)
	}
}
