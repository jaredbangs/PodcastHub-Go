package actions

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Update struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
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

	if update.repo == nil {
		update.repo = repositories.NewFeedRepository(update.Config)
	}

	currentFeed, err := parsing.TryParse(content)

	if err == nil {

		feedRecord, _ := update.repo.ReadByUrl(feedUrl)

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

		update.repo.Save(feedUrl, &feedRecord)
	}
}
