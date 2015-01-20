package main

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/parsing"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	xml, _ := ioutil.ReadFile("parsing/TestFiles/PhoneLosersOfAmerica.xml")

	processFeedContent(&xml)

	path := "parsing/TestFiles/subscriptions"

	file, err := os.Open(path)
	if err != nil {
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		processFeedUrl(scanner.Text())
	}

}

func processFeedUrl(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			fmt.Println("Getting " + feedUrl)
			response, _ := http.Get(feedUrl)
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			processFeedContent(&body)
		}
	}
}

func processFeedContent(content *[]byte) parsing.Feed {

	parser := parsing.RssParser{}

	feed, err := parser.ParseFeedContent(content)

	if err != nil {
		fmt.Printf("%s\n", content)
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(feed)
		fmt.Println(feed.Channel.Title)
		for _, item := range feed.Channel.ItemList {
			fmt.Println(item.Title)
			fmt.Println(item.Enclosure.Url)
		}
	}

	return feed
}
