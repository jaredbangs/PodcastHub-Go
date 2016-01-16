package parsing

import (
	"time"
)

type Channel struct {
	Description   string  `xml:"description"`
	Generator     string  `xml:"generator"`
	Images        []Image `xml:"image"`
	ItemList      []Item  `xml:"item"`
	LastBuildDate string  `xml:"lastBuildDate"`
	Link          string  `xml:"link"`
	Subtitle      string  `xml:"subtitle"`
	Summary       string  `xml:"summary"`
	Title         string  `xml:"title"`
}

type Image struct {
	Href string `xml:"href,attr"`
	Url  string `xml:"url"`
}

func (c Channel) GetLastBuildDate() (time.Time, error) {

	// input format: Thu, 06 Feb 2014 08:00:10 +0000

	return time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", c.LastBuildDate)
}
