package archive

import (
	"github.com/jaredbangs/PodcastHub/parsing"
	"reflect"
)

type ArchiveStrategy interface {
	ArchiveFeedItems(feed *parsing.Feed) (itemsUpdated bool, err error)
	ArchiveItem(feed *parsing.Feed, item *parsing.Item) (enclosuresUpdated bool, err error)
	GetName() string
}

var AvailableArchiveStrategies = make(map[string]reflect.Type)

func init() {

	null := &NullArchiveStrategy{}

	AvailableArchiveStrategies[null.GetName()] = reflect.TypeOf(null)
}

func GetArchiveStrategyByName(name string) (strategy ArchiveStrategy, err error) {

	if name == "" {
		null := &NullArchiveStrategy{}
		return null, nil
	} else {

		strategyType := AvailableArchiveStrategies[name]

		strategy = reflect.New(strategyType).Elem().Interface().(ArchiveStrategy)

		return strategy, nil
	}
}

type NullArchiveStrategy struct {
}

func (a *NullArchiveStrategy) ArchiveFeedItems(feed *parsing.Feed) (itemsUpdated bool, err error) {
	return false, nil
}

func (a *NullArchiveStrategy) ArchiveItem(feed *parsing.Feed, item *parsing.Item) (enclosuresUpdated bool, err error) {
	return false, nil
}

func (a *NullArchiveStrategy) GetName() (name string) {
	return reflect.TypeOf(a).Elem().Name()
}
