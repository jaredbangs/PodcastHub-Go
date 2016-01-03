package main

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
	"os"
	"strings"
)

func main() {

	argsWithoutProg := os.Args[1:]

	subscriptionFilePath := "parsing/TestFiles/subscriptions"

	file, err := os.Open(subscriptionFilePath)
	if err != nil {
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	repository := boltrepository.NewRepository("PodcastHub.bolt")

	for scanner.Scan() {
		feedUrl := scanner.Text()

		if len(argsWithoutProg) == 0 || argsWithoutProg[0] == feedUrl {

			showFeedUrl(feedUrl, repository)
		}
	}
}

func showFeedUrl(feedUrl string, repository *boltrepository.Repository) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			fmt.Println("Showing " + feedUrl)

			feed := parsing.Feed{}

			err := repository.ReadInto("Feeds", feedUrl, &feed)

			if err == nil {
				showFeedContent(&feed)
			}
		}
	}
}

func showFeedContent(feed *parsing.Feed) {

	fmt.Println(feed.Channel.Title)
	for _, item := range feed.Channel.ItemList {
		fmt.Println("\t" + item.Title)
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {
				fmt.Printf("\t\t%t"+"  "+enclosure.Url+"\n", enclosure.Downloaded)
			}
		}
	}
}
