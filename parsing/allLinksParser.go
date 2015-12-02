package parsing

import (
	"errors"
	"github.com/mvdan/xurls"
	"io/ioutil"
	"net/http"
)

type AllLinksParser struct {
}

func (parser *AllLinksParser) ParseFeedContent(content []byte) (feed Feed, err error) {

	feed = Feed{}
	feed.Channel = Channel{}
	feed.Channel.Title = ""

	urls := xurls.Strict.FindAllString(string(content), -1)

	for _, url := range urls {
		enclosure := Enclosure{Url: url}
		item := Item{}
		item.Title = ""
		item.Enclosures = append(item.Enclosures, enclosure)

		feed.Channel.ItemList = append(feed.Channel.ItemList, item)
	}

	if len(feed.Channel.ItemList) == 0 {
		err = errors.New("No Items")
	}

	return feed, err
}

func (parser *AllLinksParser) ParseFeedUrl(feedUrl string) (Feed, error) {

	response, _ := http.Get(feedUrl)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return parser.ParseFeedContent(body)
}
