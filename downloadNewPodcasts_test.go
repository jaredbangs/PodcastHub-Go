package main

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFeedContentIsParseable(t *testing.T) {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "PhoneLosersOfAmerica.xml")

	var rss Rss

	err := xml.Unmarshal(content, &rss)

	assert.Nil(t, err, "error should be nil")
	assert.True(t, len(content) > 0, "content array should be full")
	assert.NotNil(t, rss, "Rss should not be nil")
	assert.NotNil(t, rss.Channel, "Channel should not be nil")
}

func TestFeedContentHasTitle(t *testing.T) {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "PhoneLosersOfAmerica.xml")

	var rss Rss

	xml.Unmarshal(content, &rss)

	assert.NotNil(t, rss.Channel.Title, "Title should not be nil")
}
