package archive

import (
	"errors"
	"github.com/jaredbangs/PodcastHub/parsing"
	"reflect"
)

type moveFileToFeedArchivePath struct {
	feedArchivePath string
}

func init() {

	thisType := &moveFileToFeedArchivePath{}

	AvailableArchiveStrategies[thisType.GetName()] = reflect.TypeOf(thisType)
}

func NewMoveFileToFeedArchivePath(feed parsing.Feed) (ArchiveStrategy, error) {

	if feed.ArchiveStrategy == "moveFileToFeedArchivePath" {

		if feed.ArchivePath != "" {

			a := &moveFileToFeedArchivePath{}
			a.feedArchivePath = feed.ArchivePath
			return a, nil
		} else {

			a := &NullArchiveStrategy{}
			err := errors.New("Feed does not have archive path")
			return a, err
		}
	} else {

		a := &NullArchiveStrategy{}
		err := errors.New("Feed does not have the moveFileToFeedArchivePath archive strategy defined")
		return a, err
	}

}

func (a *moveFileToFeedArchivePath) ArchiveFeedItems(feed parsing.Feed) (err error) {
	return nil
}

func (a *moveFileToFeedArchivePath) ArchiveItem(item parsing.Item) (err error) {
	return nil
}

func (a *moveFileToFeedArchivePath) GetName() (name string) {
	return reflect.TypeOf(a).Elem().Name()
}
