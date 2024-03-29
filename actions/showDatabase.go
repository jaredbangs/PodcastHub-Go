package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"log"
	"os"
)

type ShowDatabase struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func (show *ShowDatabase) Show(feedUrlToShow string) error {

	log.SetOutput(os.Stdout)

	show.initializeRepo()

	show.repo.ForEach(func(feed parsing.Feed) {
		if feedUrlToShow == "" || feedUrlToShow == feed.FeedUrl {
			show.showFeedContent(&feed)
		}
	})

	return nil
}

func (show *ShowDatabase) initializeRepo() {

	if show.repo == nil {
		show.repo = repositories.NewFeedRepository(show.Config)
	}
}

func (show *ShowDatabase) showFeedContent(feed *parsing.Feed) {

	log.Println(feed.Channel.Title + "\t" + feed.FeedUrl)
	log.Println(feed.Id + "\t" + feed.LastUpdated.String())
	log.Println("ArchivePath\t" + feed.ArchivePath)
	log.Println("ArchiveStrategy\t" + feed.ArchiveStrategy)
	if feed.ForceAllLinksParser {
		log.Println("ForceAllLinksParser\ttrue")
	}
	log.Println("Subtitle:\t" + feed.Channel.Subtitle)
	log.Println("LastBuildDate:\t" + feed.Channel.LastBuildTime.String())
	log.Println("Summary:\t" + feed.Channel.Summary)
	log.Println("LastFileDownloaded:\t" + feed.LastFileDownloadedTime.Format("2006-01-02"))
	for _, item := range feed.Channel.ItemList {
		log.Println("\t" + item.Title + "\t- Published: " + item.PubTime.String())
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {
				log.Printf("\t\tDownloaded: %t\t"+enclosure.Url+"\n", enclosure.Downloaded)
				log.Printf("\t\tDownloaded to: \t" + enclosure.DownloadedFilePath + "\n")
			}
		}
	}
	log.Println(feed.Channel.Title + "\t" + feed.Id + "\t" + feed.FeedUrl + "\t" + feed.LastUpdated.String())
	log.Println("Subtitle:\t" + feed.Channel.Subtitle)
	log.Println("LastBuildDate:\t" + feed.Channel.LastBuildTime.String())
	log.Println("Summary:\t" + feed.Channel.Summary)
	log.Println("LastFileDownloaded:\t" + feed.LastFileDownloadedTime.Format("2006-01-02"))
}
