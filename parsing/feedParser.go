package parsing

import ()

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
	Title     string    `xml:"title"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Url string `xml:"url,attr"`
}
