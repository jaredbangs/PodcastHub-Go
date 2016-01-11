package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"log"
	"strings"
)

type ShowDatabase struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func (show *ShowDatabase) Show(feedUrlToShow string) error {

	show.initializeRepo()

	show.repo.ForEach(func(feedUrl string, feed parsing.Feed) {
		if feedUrlToShow == "" || feedUrlToShow == feedUrl {
			show.showFeedUrl(feedUrl)
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

	log.Println(feed.Channel.Title)
	for _, item := range feed.Channel.ItemList {
		log.Println("\t" + item.Title)
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {
				log.Printf("\t\tDownloaded: %t\t"+enclosure.Url+"\n", enclosure.Downloaded)
			}
		}
	}
}

func (show *ShowDatabase) showFeedUrl(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			log.Println("Showing " + feedUrl)

			feed, err := show.repo.Read(feedUrl)

			if err == nil {
				show.showFeedContent(&feed)
			}
		}
	}
}
