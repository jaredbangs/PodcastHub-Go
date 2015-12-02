package parsing

import (
	"time"
)

type FeedParser interface {
	parseFeedContent(content []byte) (Feed, error)
	parseFeedUrl(feedUrl string) (Feed, error)
}

type Feed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	ItemList []Item `xml:"item"`
	Title    string `xml:"title"`
}

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

type Enclosure struct {
	Url string `xml:"url,attr"`
}
