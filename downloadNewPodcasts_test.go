package main

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestProcessFeedContentContainsTitle(t *testing.T) {

	var content []byte

	content, _ = ioutil.ReadFile("TestFiles" + string(os.PathSeparator) + "PhoneLosersOfAmerica.xml")

	var channel Rss

	err := xml.Unmarshal(content, &channel)

	assert.Nil(t, err, "error should be nil")

	if err != nil {
		fmt.Printf("%s\n", content)
		fmt.Printf("error: %v", err)
	} else {
		assert.NotNil(t, channel.Channel.Title, "Title should not be nil")
	}

}
