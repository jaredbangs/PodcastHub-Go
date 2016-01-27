package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"log"
	"os"
)

type FindDuplicates struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func (show *FindDuplicates) Show(feedUrlToShow string) error {

	log.SetOutput(os.Stdout)

	show.initializeRepo()

	show.repo.ForEach(func(feed parsing.Feed) {
		if feedUrlToShow == "" || feedUrlToShow == feed.FeedUrl {
			show.showFeedContent(&feed)
		}
	})

	return nil
}

func (show *FindDuplicates) initializeRepo() {

	if show.repo == nil {
		show.repo = repositories.NewFeedRepository(show.Config)
	}
}

func (show *FindDuplicates) showFeedContent(feed *parsing.Feed) {

	log.Println(feed.Channel.Title + "\t" + feed.Id + "\t" + feed.FeedUrl + "\t" + feed.LastUpdated.String())
}
