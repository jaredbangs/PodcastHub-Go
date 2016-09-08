package archive

import (
	"errors"
	"github.com/jaredbangs/PodcastHub/parsing"
	"log"
	"reflect"
)

type deleteFile struct {
}

func init() {

	thisType := &deleteFile{}

	AvailableArchiveStrategies[thisType.GetName()] = reflect.TypeOf(thisType)
}

func NewDeleteFile(feed *parsing.Feed) (ArchiveStrategy, error) {

	if feed.ArchiveStrategy == "deleteFile" {

		a := &deleteFile{}
		return a, nil
	} else {

		a := &NullArchiveStrategy{}
		err := errors.New("Feed does not have the deleteFile archive strategy defined")
		return a, err
	}
}

func (a *deleteFile) ArchiveFeedItems(feed *parsing.Feed) (itemsUpdated bool, err error) {

	itemsUpdated = false

	log.Println("Archiving items in feed " + feed.Channel.Title)
	for _, item := range feed.Channel.ItemList {
		if !item.IsArchived && item.IsToBeArchived {
			log.Println("Deleting item " + item.Title + "in feed " + feed.Channel.Title)
			itemUpdated, _ := a.ArchiveItem(feed, &item)

			itemsUpdated = itemsUpdated || itemUpdated
		}
	}

	return itemsUpdated, err
}

func (a *deleteFile) ArchiveItem(feed *parsing.Feed, item *parsing.Item) (enclosuresUpdated bool, err error) {

	enclosuresUpdated = false

	for _, enclosure := range item.Enclosures {

		if len(enclosure.DownloadedFilePath) != 0 {

			mover := &FileMover{}

			if mover.FileExists(enclosure.DownloadedFilePath) {

				mover.DeleteFile(enclosure.DownloadedFilePath)

				log.Println("Deleted " + enclosure.DownloadedFilePath)

				enclosure.DownloadedFilePath = ""

			} else {
				log.Println(enclosure.DownloadedFilePath + " file does not exist")
			}
			enclosuresUpdated = true
		} else if len(enclosure.Url) != 0 {
			enclosuresUpdated = true
		}
	}

	if enclosuresUpdated {
		item.IsArchived = true
		feed.UpdateItem(*item)
	}

	return enclosuresUpdated, nil
}

func (a *deleteFile) GetName() (name string) {
	return reflect.TypeOf(a).Elem().Name()
}
