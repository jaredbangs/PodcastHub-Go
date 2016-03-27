package parsing

import (
	"github.com/satori/go.uuid"
	"log"
	"time"
)

type Channel struct {
	Description   string  `xml:"description"`
	Generator     string  `xml:"generator"`
	Images        []Image `xml:"image"`
	ItemList      []Item  `xml:"item"`
	LastBuildDate string  `xml:"lastBuildDate"`
	LastBuildTime time.Time
	Link          string `xml:"link"`
	Subtitle      string `xml:"subtitle"`
	Summary       string `xml:"summary"`
	Title         string `xml:"title"`
}

type Image struct {
	Href string `xml:"href,attr"`
	Url  string `xml:"url"`
}

func (c *Channel) GetLastBuildDate() (time.Time, error) {

	// input format: Thu, 06 Feb 2014 08:00:10 +0000

	return time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", c.LastBuildDate)
}

func (c *Channel) MakeSureAllItemsHaveIds(feedId string) (err error) {

	for i, item := range c.ItemList {

		if item.FeedId == "" || item.ItemId == "" {

			newId := uuid.NewV4()
			item.ItemId = newId.String()
			item.FeedId = feedId

			c.ItemList[i] = item
		}
	}

	return nil
}

func (c *Channel) ReprocessExistingInfo(logInfo bool) (err error) {

	if logInfo {
		log.Println("Reprocessing " + c.Title)
	}

	if c.LastBuildTime.IsZero() && c.LastBuildDate != "" {
		parsed, err := time.Parse(time.RFC1123Z, c.LastBuildDate)

		if err != nil {
			if logInfo {
				log.Println("Error parsing time " + c.LastBuildDate)
			}
		}

		c.LastBuildTime = parsed
	}

	for i, item := range c.ItemList {
		item.ParseTimeIfNecessary()
		item.ReprocessExistingInfo(logInfo)
		c.ItemList[i] = item
	}

	return nil
}
