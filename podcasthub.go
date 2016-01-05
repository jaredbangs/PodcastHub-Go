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
		case "download":
			d := &actions.Download{Config: config}

			if len(argsWithoutProg) > 1 {
				d.DownloadNewFilesInFeed(argsWithoutProg[1])
			} else {
				d.DownloadAllNewFiles()
			}
		case "markdownloaded":
			d := &actions.Download{Config: config}

			if len(argsWithoutProg) > 1 {
				d.MarkAllNewFilesDownloadedInFeed(argsWithoutProg[1])
			} else {
				d.MarkAllNewFilesDownloaded()
			}
		case "show":
			show := &actions.ShowDatabase{Config: config}

			if len(argsWithoutProg) > 1 {
				show.Show(argsWithoutProg[1])
			} else {
				show.Show("")
			}
		case "update":
			u := &actions.Update{Config: config}

			if len(argsWithoutProg) > 1 {
				u.UpdateFeed(argsWithoutProg[1])
			} else {
				u.Update()
			}
		}
	} else {
		fmt.Println("Usage: podcasthub action [options]")
	}
}
