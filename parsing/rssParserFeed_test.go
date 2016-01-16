package parsing

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFeedContentDescription(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	assert.Equal(t, "The happiest place in Roy, New Mexico!", feed.Channel.Description)
}

func TestFeedContentIsNotParseable(t *testing.T) {

	parser := RssParser{}

	feed, err := parser.ParseFeedContent(getInvalidTestContent())

	assert.NotNil(t, err, "error should not be nil")
	assert.NotNil(t, feed, "Rss should not be nil")
}

func TestFeedContentIsParseable(t *testing.T) {

	parser := RssParser{}

	feed, err := parser.ParseFeedContent(getValidTestContent())

	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "Rss should not be nil")
	assert.NotNil(t, feed.Channel, "Channel should not be nil")
}

func TestFeedContentHasItems(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	assert.NotNil(t, feed.Channel.ItemList, "ItemList should not be nil")
	assert.Equal(t, 50, len(feed.Channel.ItemList), "ItemList should have 50 items")
}

func TestFeedContentLink(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	assert.Equal(t, "http://www.phonelosers.org", feed.Channel.Link)
}

func TestFeedContentTitle(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	assert.Equal(t, "Phone Losers of America", feed.Channel.Title)
}

func TestFeedShouldHave47Enclosures(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	enclosureCount := 0

	for _, item := range feed.Channel.ItemList {
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {
				enclosureCount = enclosureCount + 1
			}
		}
	}

	assert.Equal(t, 47, enclosureCount, "47 items should have enclosures")

}

func TestItemsUsuallyHaveEnclosures(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	emptyEnclosureItemCount := 0

	for _, item := range feed.Channel.ItemList {
		if len(item.Enclosures) == 0 {
			emptyEnclosureItemCount = emptyEnclosureItemCount + 1
		}
	}

	assert.Equal(t, 5, emptyEnclosureItemCount, "Only five items should not have enclosures")

}

func TestItemsAlwaysHaveTitles(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	for _, item := range feed.Channel.ItemList {
		assert.NotEmpty(t, item.Title)
	}
}

func getInvalidTestContent() []byte {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "InvalidContent.xml")

	return content
}

func getValidTestContent() []byte {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "PhoneLosersOfAmerica.xml")

	return content
}
