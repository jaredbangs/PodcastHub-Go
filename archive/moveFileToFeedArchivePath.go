package archive

import (
	"errors"
	"github.com/jaredbangs/PodcastHub/parsing"
	"log"
	"reflect"
)

type moveFileToFeedArchivePath struct {
	feedArchivePath string
}

func init() {

	thisType := &moveFileToFeedArchivePath{}

	AvailableArchiveStrategies[thisType.GetName()] = reflect.TypeOf(thisType)
}

func NewMoveFileToFeedArchivePath(feed *parsing.Feed) (ArchiveStrategy, error) {

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

func (a *moveFileToFeedArchivePath) ArchiveFeedItems(feed *parsing.Feed) (itemsUpdated bool, err error) {

	itemsUpdated = false

	log.Println("Archiving items in feed " + feed.Channel.Title)
	for _, item := range feed.Channel.ItemList {
		if !item.IsArchived && item.IsToBeArchived {
			log.Println("Moving item " + item.Title + "in feed " + feed.Channel.Title)
			itemUpdated, _ := a.ArchiveItem(feed, &item)

			itemsUpdated = itemsUpdated || itemUpdated
		}
	}

	return itemsUpdated, err
}

func (a *moveFileToFeedArchivePath) ArchiveItem(feed *parsing.Feed, item *parsing.Item) (enclosuresUpdated bool, err error) {

	enclosuresUpdated = false

	for _, enclosure := range item.Enclosures {

		if len(enclosure.DownloadedFilePath) != 0 {

			mover := &FileMover{}

			if mover.FileExists(enclosure.DownloadedFilePath) {

				archiveFilePath := mover.MoveFileIntoDirectory(enclosure.DownloadedFilePath, feed.ArchivePath)

				log.Println("Moved " + enclosure.DownloadedFilePath + " to " + archiveFilePath)

				enclosure.DownloadedFilePath = archiveFilePath

				enclosuresUpdated = true
				//feed.UpdateEnclosure(enclosure)
			} else {
				log.Println(enclosure.DownloadedFilePath + " file does not exist")
			}
		}
	}

	if enclosuresUpdated {
		item.IsArchived = true
		feed.UpdateItem(*item)
	}

	return enclosuresUpdated, nil
}

func (a *moveFileToFeedArchivePath) GetName() (name string) {
	return reflect.TypeOf(a).Elem().Name()
}
