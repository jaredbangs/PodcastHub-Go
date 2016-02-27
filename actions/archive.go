package actions

import (
	"github.com/jaredbangs/PodcastHub/archive"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/repositories"
)

type archiveItems struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func NewArchive(config config.Configuration) *archiveItems {

	a := &archiveItems{Config: config}
	a.initializeRepo()
	return a
}

func (a *archiveItems) ArchivePendingItems() {

	for _, feedId := range a.repo.GetAllIds() {

		feed, _ := a.repo.ReadById(feedId)

		archiver, _ := archive.GetArchiveStrategyByName(feed.ArchiveStrategy)

		itemsUpdated, err := archiver.ArchiveFeedItems(&feed)

		if itemsUpdated && (err == nil) {
			a.repo.Save(&feed)
		}
	}
}

func (a *archiveItems) initializeRepo() {

	if a.repo == nil {
		a.repo = repositories.NewFeedRepository(a.Config)
	}
}
