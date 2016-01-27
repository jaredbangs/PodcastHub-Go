package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"log"
	"os"
	"strings"
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
			show.showFeedUrl(feed.FeedUrl)
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

func (show *FindDuplicates) showFeedUrl(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := show.repo.ReadByUrl(feedUrl)

			if err == nil {
				show.showFeedContent(&feed)
			}
		}
	}
}
