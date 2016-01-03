package main

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/actions"
	"github.com/jaredbangs/PodcastHub/config"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	config := config.GetConfig()

	if len(argsWithoutProg) > 0 {
		switch argsWithoutProg[0] {
		case "update":
			update := &actions.Update{Config: config}
			update.Update()
		case "show":
			show := &actions.ShowDatabase{Config: config}

			if len(argsWithoutProg) > 1 {
				show.Show(argsWithoutProg[1])

			} else {
				show.Show("")
			}
		}
	} else {
		fmt.Println("Usage: podcasthub action [options]")
	}
}
