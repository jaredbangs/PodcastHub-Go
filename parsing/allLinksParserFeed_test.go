package parsing

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestAllLinksFeedContentIsNotParseable(t *testing.T) {

	parser := AllLinksParser{}

	feed, err := parser.ParseFeedContent(getInvalidAllLinksTestContent())

	assert.NotNil(t, err, "error should not be nil")
	assert.NotNil(t, feed, "Rss should not be nil")
}

func TestAllLinksFeedContentIsParseable(t *testing.T) {

	parser := AllLinksParser{}

	feed, err := parser.ParseFeedContent(getValidAllLinksTestContent())

	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, feed, "Rss should not be nil")
	assert.NotNil(t, feed.Channel, "Channel should not be nil")
}

func TestAllLinksFeedContentHasItems(t *testing.T) {

	parser := AllLinksParser{}

	feed, _ := parser.ParseFeedContent(getValidAllLinksTestContent())

	assert.NotNil(t, feed.Channel.ItemList, "ItemList should not be nil")
	assert.Equal(t, 276, len(feed.Channel.ItemList), "ItemList should have 824 items")
}

func TestAllLinksFeedContentHasTitle(t *testing.T) {

	parser := AllLinksParser{}

	feed, _ := parser.ParseFeedContent(getValidAllLinksTestContent())

	assert.NotNil(t, feed.Channel.Title, "Title should not be nil")
}

func TestAllLinksFeedShouldHave276Enclosures(t *testing.T) {

	parser := AllLinksParser{}

	feed, _ := parser.ParseFeedContent(getValidAllLinksTestContent())

	enclosureCount := 0

	for _, item := range feed.Channel.ItemList {
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {
				enclosureCount = enclosureCount + 1
			}
		}
	}

	assert.Equal(t, 276, enclosureCount, "276 items should have enclosures")

}

func TestAllLinksItemsAlwaysHaveEnclosures(t *testing.T) {

	parser := AllLinksParser{}

	feed, _ := parser.ParseFeedContent(getValidAllLinksTestContent())

	emptyEnclosureItemCount := 0

	for _, item := range feed.Channel.ItemList {
		if len(item.Enclosures) == 0 {
			emptyEnclosureItemCount = emptyEnclosureItemCount + 1
		}
	}

	assert.Equal(t, 0, emptyEnclosureItemCount, "No items should not have enclosures")

}

func getInvalidAllLinksTestContent() []byte {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "InvalidContent.xml")

	return content
}

func getValidAllLinksTestContent() []byte {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "c697f48392907a0Err.xml")

	return content
}
