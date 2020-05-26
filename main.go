package main

import (
	"flag"
	"fmt"

	"github.com/ytakahashi/raindrop-alfred-workflow/raindrop"
)

func main() {
	flag.Parse()
	accessToken := flag.Arg(0)

	client, err := raindrop.NewClient(accessToken)

	res, err := client.GetRaindrops("0")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v", *res)
	}
}
