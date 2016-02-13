package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/repositories"
)

type parseTimesIfNecessary struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func ParseTimesIfNecessary(config config.Configuration) *parseTimesIfNecessary {

	d := &parseTimesIfNecessary{Config: config}

	d.initializeRepo()

	return d
}

func (d *parseTimesIfNecessary) Run() error {

	for _, feedId := range d.repo.GetAllIds() {

		feed, err := d.repo.ReadById(feedId)
		if err != nil {
			return err
		}

		feed.Channel.ParseTimesIfNecessary()

		d.repo.Save(&feed)
	}

	return nil
}

func (d *parseTimesIfNecessary) initializeRepo() {

	if d.repo == nil {
		d.repo = repositories.NewFeedRepository(d.Config)
	}
}
