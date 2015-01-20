package parsing

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type RssParser struct {
}

func (parser *RssParser) parseFeedContent(content []byte) (Feed, error) {

	var feed Feed

	err := xml.Unmarshal(content, &feed)

	return feed, err
}

func (parser *RssParser) parseFeedUrl(feedUrl string) (Feed, error) {

	response, _ := http.Get(feedUrl)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return parser.parseFeedContent(body)
}
