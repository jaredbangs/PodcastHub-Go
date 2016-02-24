package archive

import (
	"github.com/jaredbangs/PodcastHub/parsing"
	"reflect"
)

type ArchiveStrategy interface {
	ArchiveFeedItems(feed parsing.Feed) (err error)
	ArchiveItem(item parsing.Item) (err error)
	GetName() string
}

var AvailableArchiveStrategies = make(map[string]reflect.Type)

func init() {

	null := &NullArchiveStrategy{}

	AvailableArchiveStrategies[null.GetName()] = reflect.TypeOf(null)
}

type NullArchiveStrategy struct {
}

func (a *NullArchiveStrategy) ArchiveFeedItems(feed parsing.Feed) (err error) {
	return nil
}

func (a *NullArchiveStrategy) ArchiveItem(item parsing.Item) (err error) {
	return nil
}

func (a *NullArchiveStrategy) GetName() (name string) {
	return reflect.TypeOf(a).Elem().Name()
}
