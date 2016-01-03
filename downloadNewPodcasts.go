package main

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	subscriptionFilePath := "parsing/TestFiles/subscriptions"

	file, err := os.Open(subscriptionFilePath)
	if err != nil {
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	repository := boltrepository.NewRepository("PodcastHub.bolt")

	for scanner.Scan() {
		processFeedUrl(scanner.Text(), repository)
	}
}

func processFeedUrl(feedUrl string, repository *boltrepository.Repository) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			fmt.Println("Getting " + feedUrl)

			response, httpErr := http.Get(feedUrl)

			if httpErr != nil {

				fmt.Println("ERR\t" + httpErr.Error())
			} else {

				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)
				//processFeedContent(&body)
				//fmt.Println(checkRssParser(body))
				recordFeedInfo(feedUrl, body, repository)
			}
		}
	}
}

func recordFeedInfo(feedUrl string, content []byte, repository *boltrepository.Repository) {

	feed, err := tryParse(content)

	if err == nil {
		ioutil.WriteFile("parsing"+string(os.PathSeparator)+"TestFiles"+string(os.PathSeparator)+feed.Channel.Title+".xml", content, 0644)
		fmt.Println("OK\t" + feed.Channel.Title + " " + strconv.FormatInt(int64(len(feed.Channel.ItemList)), 10) + " items")

		repository.Save("Feeds", feedUrl, feed)

	} else {
		errFileName := strconv.FormatInt(int64(rand.Int()), 10) + "Err.xml"

		ioutil.WriteFile("parsing"+string(os.PathSeparator)+"TestFiles"+string(os.PathSeparator)+errFileName, content, 0644)
		fmt.Println("ERR\t" + err.Error())
	}
}

func tryParse(content []byte) (feed parsing.Feed, err error) {

	parser := parsing.RssParser{}

	feed, err = parser.ParseFeedContent(content)

	if err == nil {
		return feed, err
	} else {

		parser := parsing.AllLinksParser{}

		feed, err = parser.ParseFeedContent(content)

		return feed, err
	}
}

func checkRssParser(content []byte) string {

	feed, err := tryParse(content)

	if err != nil {
		errFileName := strconv.FormatInt(int64(rand.Int()), 10) + "Err.xml"

		ioutil.WriteFile("parsing"+string(os.PathSeparator)+"TestFiles"+string(os.PathSeparator)+errFileName, content, 0644)
		return "ERR\t" + err.Error()
	} else {
		ioutil.WriteFile("parsing"+string(os.PathSeparator)+"TestFiles"+string(os.PathSeparator)+feed.Channel.Title+".xml", content, 0644)
		return "OK\t" + feed.Channel.Title + " " + strconv.FormatInt(int64(len(feed.Channel.ItemList)), 10) + " items"
	}
}

func processFeedContent(content []byte) parsing.Feed {

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
			fmt.Println(item.Enclosures[0].Url)
		}
	}

	return feed
}
