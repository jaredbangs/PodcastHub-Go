package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/repositories"
)

type reprocessExistingInfo struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func ReprocessExistingInfo(config config.Configuration) *reprocessExistingInfo {

	d := &reprocessExistingInfo{Config: config}

	d.initializeRepo()

	return d
}

func (d *reprocessExistingInfo) ReprocessFeed(feedUrl string) error {

	feed, err := d.repo.ReadByUrl(feedUrl)
	if err != nil {
		return err
	}

	feed.ReprocessExistingInfo(true)

	d.repo.Save(&feed)

	return nil
}

func (d *reprocessExistingInfo) Run() error {

	for _, feedId := range d.repo.GetAllIds() {

		feed, err := d.repo.ReadById(feedId)
		if err != nil {
			return err
		}

		feed.ReprocessExistingInfo(true)

		d.repo.Save(&feed)
	}

	return nil
}

func (d *reprocessExistingInfo) initializeRepo() {

	if d.repo == nil {
		d.repo = repositories.NewFeedRepository(d.Config)
	}
}
