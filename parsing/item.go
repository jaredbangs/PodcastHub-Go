package parsing

import (
	"time"
)

type Item struct {
	// Have to specify where to find the series title since
	// the field of this struct doesn't match the xml tag
	Enclosures []Enclosure `xml:"enclosure"`
	PubDate    string      `xml:"pubDate"`
	Title      string      `xml:"title"`
}

func (i Item) GetPublicationDate() (time.Time, error) {

	// input format: Thu, 06 Feb 2014 08:00:10 +0000

	return time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", i.PubDate)
}
