package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	xml, _ := ioutil.ReadFile("PhoneLosersOfAmerica.xml")

	processFeedContent(xml)

	path := "subscriptions"

	file, err := os.Open(path)
	if err != nil {
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//	processFeedUrl(scanner.Text())
	}

}

func processFeedUrl(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			fmt.Println("Getting " + feedUrl)
			response, _ := http.Get(feedUrl)
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			processFeedContent(body)
		}
	}
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	ItemList []Item `xml:"item"`
	Title    string `xml:"title"`
}

type Item struct {
	// Have to specify where to find the series title since
	// the field of this struct doesn't match the xml tag
	Title     string    `xml:"title"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Url string `xml:"url,attr"`
}

func processFeedContent(content []byte) {

	var channel Rss

	err := xml.Unmarshal(content, &channel)
	if err != nil {
		fmt.Printf("%s\n", content)
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(channel)
		fmt.Println(channel.Channel.Title)
		for _, item := range channel.Channel.ItemList {
			fmt.Println(item.Title)
			fmt.Println(item.Enclosure.Url)
		}
	}

}
