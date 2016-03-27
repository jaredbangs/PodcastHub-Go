package parsing

import (
	"log"
	"time"
)

type Item struct {
	Description    string      `xml:"description"`
	Duration       string      `xml:"duration"`
	Enclosures     []Enclosure `xml:"enclosure"`
	Encoded        string      `xml:"encoded"`
	FeedId         string
	IsArchived     bool
	IsToBeArchived bool
	ItemId         string
	PubDate        string `xml:"pubDate"`
	PubTime        time.Time
	Subtitle       string `xml:"subtitle"`
	Summary        string `xml:"summary"`
	Title          string `xml:"title"`
}

func (i *Item) GetPublicationDate() (time.Time, error) {

	// input format: Thu, 06 Feb 2014 08:00:10 +0000

	return time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", i.PubDate)
}

func (i *Item) ParseTimeIfNecessary() (err error) {

	if i.PubTime.IsZero() && i.PubDate != "" {
		parsed, err := time.Parse(time.RFC1123Z, i.PubDate)

		if err != nil {
			return err
		}

		i.PubTime = parsed
	}

	return nil
}

func (i *Item) ReprocessExistingInfo(logInfo bool) (err error) {

	if logInfo {
		log.Println("Processing item " + i.Title)
	}

	for index, enclosure := range i.Enclosures {
		enclosure.UpdateCalculatedInfo(logInfo)
		i.Enclosures[index] = enclosure
	}

	return nil
}
