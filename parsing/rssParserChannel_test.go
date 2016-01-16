package parsing

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestItemGetChannelGenerator(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	assert.Equal(t, "http://wordpress.org/?v=3.8.1", feed.Channel.Generator)
}

func TestItemGetChannelImageUrl(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	imageUrl := ""

	for _, image := range feed.Channel.Images {
		if image.Url != "" {
			imageUrl = image.Url
		}
	}

	assert.Equal(t, "http://www.phonelosers.org/images/site_images/podcast_144x144.png", imageUrl)
}

func TestItemGetChannelItunesImageUrl(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	imageUrl := ""

	for _, image := range feed.Channel.Images {
		if image.Href != "" {
			imageUrl = image.Href
		}
	}

	assert.Equal(t, "http://www.phonelosers.org/images/site_images/podcast_300x300.png", imageUrl)
}

func TestItemGetChannelSubtitle(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	assert.Equal(t, "Prank phone calls and other phoneloserish things by the people at Phone Losers of America and Cacti Radio.", feed.Channel.Subtitle)
}

func TestItemGetChannelSummary(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	summary := feed.Channel.Summary

	assert.Contains(t, summary, "The Phone Losers of America and their prank calls")
	assert.Contains(t, summary, "past 10 years of content.")
}

func TestItemGetChannelLastBuildDate(t *testing.T) {

	parser := RssParser{}

	feed, _ := parser.ParseFeedContent(getValidTestContent())

	date, _ := feed.Channel.GetLastBuildDate()

	assert.Equal(t, 2014, date.Year(), "Year does not match")
	assert.Equal(t, 2, date.Month(), "Month does not match")
	assert.Equal(t, 13, date.Day(), "Day does not match")
	assert.Equal(t, 5, date.Hour(), "Hour does not match")
	assert.Equal(t, 4, date.Minute(), "Minute does not match")
	assert.Equal(t, 35, date.Second(), "Second does not match")

	_, zoneOffset := date.Zone()

	assert.Equal(t, 0, zoneOffset, "Offset does not match")
}

func getInvalidTestChannelContent() []byte {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "InvalidContent.xml")

	return content
}

func getValidTestChannelContent() []byte {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "PhoneLosersOfAmerica.xml")

	return content
}
