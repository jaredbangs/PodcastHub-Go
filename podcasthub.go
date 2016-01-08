package main

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/actions"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/web"
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
		case "processfiles":
			p := &actions.ProcessDownloadedFiles{}

			if len(argsWithoutProg) > 1 {
				p.ApplyAllFilters(argsWithoutProg[1])
			} else {
				fmt.Println("Usage: podcasthub processfiles path")
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
		case "web":
			w := &web.Web{Config: config}
			w.Start()
		}
	} else {
		fmt.Println("Usage: podcasthub action [options]")
	}
}
