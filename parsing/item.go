package parsing

import (
	"time"
)

type Item struct {
	Description string      `xml:"description"`
	Duration    string      `xml:"duration"`
	Enclosures  []Enclosure `xml:"enclosure"`
	Encoded     string      `xml:"encoded"`
	PubDate     string      `xml:"pubDate"`
	Subtitle    string      `xml:"subtitle"`
	Summary     string      `xml:"summary"`
	Title       string      `xml:"title"`
}

func (i Item) GetPublicationDate() (time.Time, error) {

	// input format: Thu, 06 Feb 2014 08:00:10 +0000

	return time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", i.PubDate)
}
