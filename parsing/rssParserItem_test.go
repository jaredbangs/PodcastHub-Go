package parsing

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestItemEnclosures(t *testing.T) {

	item := getValidTestItem(0)

	assert.Equal(t, "http://www.phonelosers.org/podpress_trac/feed/3455/0/snow_plow_show_2014-02-05.mp3", item.Enclosures[0].Url, "Enclosure does not match")
	assert.Equal(t, "http://friday.erinthepir8.com/misc/carlito%27s%20roommate.mp3", item.Enclosures[1].Url, "Enclosure does not match")
}

func TestItemGetPublicationDate(t *testing.T) {

	item := getValidTestItem(0)

	pub, _ := item.GetPublicationDate()

	assert.Equal(t, 2014, pub.Year(), "Year does not match")
	assert.Equal(t, 2, pub.Month(), "Month does not match")
	assert.Equal(t, 6, pub.Day(), "Day does not match")
	assert.Equal(t, 8, pub.Hour(), "Hour does not match")
	assert.Equal(t, 0, pub.Minute(), "Minute does not match")
	assert.Equal(t, 10, pub.Second(), "Second does not match")

	_, zoneOffset := pub.Zone()

	assert.Equal(t, 0, zoneOffset, "Offset does not match")
}

func TestItemPubDateString(t *testing.T) {

	item := getValidTestItem(0)

	assert.Equal(t, "Thu, 06 Feb 2014 08:00:10 +0000", item.PubDate, "PubDate does not match")
}

func TestItemTitle(t *testing.T) {

	item := getValidTestItem(0)

	assert.Equal(t, "Snow Plow Show – February 5th, 2014 – Coke Terrorism", item.Title, "Title does not match")
}

func getValidTestItem(index int) Item {

	content, _ := ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "PhoneLosersOfAmerica.xml")

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(content)

	return feed.Channel.ItemList[index]
}
