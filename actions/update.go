package actions

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Update struct {
	Config config.Configuration
	Repo   *boltrepository.Repository
}

func (update *Update) Update() error {

	file, err := os.Open(update.Config.SubscriptionFile)
	defer file.Close()

	if err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			update.UpdateFeed(scanner.Text())
		}
	}

	return err
}

func (update *Update) UpdateFeed(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			fmt.Println("Getting " + feedUrl)

			response, httpErr := http.Get(feedUrl)

			if httpErr != nil {

				fmt.Println("ERR\t" + httpErr.Error())
			} else {

				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)
				update.recordFeedInfo(feedUrl, body)
			}
		}
	}
}

func (update *Update) recordFeedInfo(feedUrl string, content []byte) {

	if update.Repo == nil {
		update.Repo = boltrepository.NewRepository(update.Config.RepositoryFile)
	}

	currentFeed, err := parsing.TryParse(content)

	if err == nil {

		feedRecord := parsing.Feed{}
		update.Repo.ReadInto("Feeds", feedUrl, &feedRecord)

		for _, item := range currentFeed.Channel.ItemList {

			hasUrl := true

			for _, enclosure := range item.Enclosures {
				if hasUrl {
					hasUrl = feedRecord.ContainsEnclosureUrl(enclosure.Url)
				}
			}

			if !hasUrl {
				feedRecord.AddItem(item)
				for _, enclosure := range item.Enclosures {
					fmt.Println("Added item: " + enclosure.Url)
				}
			}
		}

		update.Repo.Save("Feeds", feedUrl, feedRecord)
	}
}
