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

func (c *Channel) ParseTimesIfNecessary() (err error) {

	if c.LastBuildTime.IsZero() && c.LastBuildDate != "" {
		parsed, err := time.Parse(time.RFC1123Z, c.LastBuildDate)

		if err != nil {
			return err
		}

		c.LastBuildTime = parsed
	}

	for i, item := range c.ItemList {
		item.ParseTimeIfNecessary()
		c.ItemList[i] = item
	}

	return nil
}
