package main

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/actions"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/repositories"
	"github.com/jaredbangs/PodcastHub/web"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	config := config.GetConfig()

	if len(argsWithoutProg) > 0 {
		switch argsWithoutProg[0] {
		case "archive":
			a := actions.NewArchive(config)
			a.ArchivePendingItems()
		case "clone":
			d := repositories.NewFeedRepository(config)
			d.Clone()
		case "download":
			d := actions.NewDownload(config)
			defer d.Close()

			if len(argsWithoutProg) > 2 {
				d.DownloadFileInFeed(argsWithoutProg[1], argsWithoutProg[2])
			} else if len(argsWithoutProg) > 1 {
				d.DownloadNewFilesInFeed(argsWithoutProg[1])
			} else {
				d.DownloadAllNewFiles()
			}
		case "dups":
			show := &actions.FindDuplicates{Config: config}
			if len(argsWithoutProg) > 1 {
				show.Show(argsWithoutProg[1])
			} else {
				show.Show("")
			}
		case "markdownloaded":
			d := actions.NewDownload(config)
			defer d.Close()

			if len(argsWithoutProg) > 1 {
				d.MarkAllNewFilesDownloadedInFeed(argsWithoutProg[1])
			} else {
				d.MarkAllNewFilesDownloaded()
			}
		case "move":
			p := &actions.Reorg{Config: config}
			p.CopyFromOldToNew()
		case "processitems":
			p := &actions.ProcessDownloadedItems{}
			p.ApplyAllFilters(config)
		case "processfiles":
			p := &actions.ProcessDownloadedFiles{}

			if len(argsWithoutProg) > 1 {
				p.ApplyAllFilters(argsWithoutProg[1])
			} else {
				fmt.Println("Usage: podcasthub processfiles path")
			}
		case "readfiles":
			p := &actions.ReadFiles{}

			if len(argsWithoutProg) > 1 {
				p.ApplyAllFilters(argsWithoutProg[1])
			} else {
				fmt.Println("Usage: podcasthub readfiles path")
			}
		case "reprocess":
			d := actions.ReprocessExistingInfo(config)
			d.Run()
		case "show":
			show := &actions.ShowDatabase{Config: config}

			if len(argsWithoutProg) > 1 {
				show.Show(argsWithoutProg[1])
			} else {
				show.Show("")
			}
		case "subscribe":
			s := &actions.Subscribe{Config: config}

			s.Subscribe(argsWithoutProg[1])
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
