package main

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/parsing"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	//xml, _ := ioutil.ReadFile("parsing/TestFiles/PhoneLosersOfAmerica.xml")

	//processFeedContent(&xml)

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

			response, httpErr := http.Get(feedUrl)

			if httpErr != nil {

				fmt.Println("ERR\t" + httpErr.Error())

			} else {

				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)
				//processFeedContent(&body)
				fmt.Println(checkRssParser(body))
			}
		}
	}
}

func checkRssParser(content []byte) string {

	parser := parsing.RssParser{}

	feed, err := parser.ParseFeedContent(content)

	if err != nil {
		errFileName := strconv.FormatInt(int64(rand.Int()), 16) + "Err.xml"

		ioutil.WriteFile("parsing"+string(os.PathSeparator)+"TestFiles"+string(os.PathSeparator)+errFileName, content, 0644)
		return "ERR\t" + err.Error()
	} else {
		ioutil.WriteFile("parsing"+string(os.PathSeparator)+"TestFiles"+string(os.PathSeparator)+feed.Channel.Title+".xml", content, 0644)
		return "OK\t" + feed.Channel.Title
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
			fmt.Println(item.Enclosure.Url)
		}
	}

	return feed
}
