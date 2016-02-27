package archive

import (
	"errors"
	"github.com/jaredbangs/PodcastHub/parsing"
	"reflect"
)

type deleteFile struct {
}

func init() {

	thisType := &deleteFile{}

	AvailableArchiveStrategies[thisType.GetName()] = reflect.TypeOf(thisType)
}

func NewDeleteFile(feed parsing.Feed) (ArchiveStrategy, error) {

	if feed.ArchiveStrategy == "deleteFile" {

		a := &deleteFile{}
		return a, nil
	} else {

		a := &NullArchiveStrategy{}
		err := errors.New("Feed does not have the deleteFile archive strategy defined")
		return a, err
	}
}

func (a *deleteFile) ArchiveFeedItems(feed parsing.Feed) (err error) {
	return nil
}

func (a *deleteFile) ArchiveItem(item parsing.Item) (err error) {
	return nil
}

func (a *deleteFile) GetName() (name string) {
	return reflect.TypeOf(a).Elem().Name()
}