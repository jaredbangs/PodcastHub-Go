package files

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
)

type ItemFileProcessor struct {
	Config  config.Configuration
	Filters []ItemFileProcessingFilter
	repo    *repositories.FeedRepository
}

func (f *ItemFileProcessor) ProcessFiles() (err error) {

	f.initializeRepo()

	allFeedUrls := f.repo.GetAllKeys()

	for _, feedUrl := range allFeedUrls {
		err = f.processFeed(feedUrl)
	}

	return err
}

func (f *ItemFileProcessor) initializeRepo() {

	if f.repo == nil {
		f.repo = repositories.NewFeedRepository(f.Config)
	}
}

func (f *ItemFileProcessor) processFeed(feedUrl string) (err error) {

	feed, err := f.repo.ReadByUrl(feedUrl)

	if err == nil {
		for _, item := range feed.Channel.ItemList {
			err = f.processItem(feedUrl, &feed, &item)
		}
	}

	return err
}

func (f *ItemFileProcessor) processItem(feedUrl string, feed *parsing.Feed, item *parsing.Item) (err error) {

	for _, enclosure := range item.Enclosures {
		if enclosure.Downloaded && enclosure.DownloadedFilePath != "" {
			for _, filter := range f.Filters {
				if filter.ShouldProcessItem(feed, item, &enclosure) {
					err = filter.ProcessItem(feed, item, &enclosure)

					if err == nil {

						feed.UpdateEnclosure(enclosure)

						f.repo.Save(feedUrl, feed)
					}
				}
			}
		}
	}

	return err
}
