package parsing

import (
	"time"
)

type Feed struct {
	ArchivePath     string
	ArchiveStrategy string
	Channel         Channel `xml:"channel"`
	FeedUrl         string
	Id              string
	LastUpdated     time.Time
	existingUrls    map[string]int
}

func (feed *Feed) AddItem(item Item) {
	feed.Channel.ItemList = append(feed.Channel.ItemList, item)
	feed.existingUrls = nil
}

func (feed *Feed) ContainsEnclosureUrl(url string) bool {

	feed.populateExistingUrls()

	_, containsKey := feed.existingUrls[url]

	return containsKey
}

func (f *Feed) MakeSureAllItemsHaveIds() (err error) {

	f.Channel.MakeSureAllItemsHaveIds(f.Id)

	return nil
}

func (feed *Feed) UpdateEnclosure(enclosure Enclosure) {

	for _, item := range feed.Channel.ItemList {
		if item.Enclosures != nil {
			for i, existingEnclosure := range item.Enclosures {
				if enclosure.Url == existingEnclosure.Url {
					item.Enclosures = append(item.Enclosures[:i], item.Enclosures[i+1:]...)
					item.Enclosures = append(item.Enclosures, enclosure)
					break
				}
			}
		}
	}
}

func (feed *Feed) UpdateItem(item Item) {

	for i, existingItem := range feed.Channel.ItemList {
		if existingItem.ItemId == item.ItemId {
			feed.Channel.ItemList[i] = item
		}
	}
}

func (feed *Feed) populateExistingUrls() {

	if feed.existingUrls == nil {

		feed.existingUrls = make(map[string]int)

		for _, item := range feed.Channel.ItemList {
			if item.Enclosures != nil {
				for _, enclosure := range item.Enclosures {
					feed.existingUrls[enclosure.Url] = 1
				}
			}
		}
	}
}
