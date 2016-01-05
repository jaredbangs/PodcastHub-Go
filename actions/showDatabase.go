package actions

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
	"os"
	"strings"
)

type ShowDatabase struct {
	Config config.Configuration
	Repo   *boltrepository.Repository
}

func (show *ShowDatabase) Show(feedUrlToShow string) error {

	show.Repo = boltrepository.NewRepository(show.Config.RepositoryFile)

	file, err := os.Open(show.Config.SubscriptionFile)
	defer file.Close()

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			feedUrl := scanner.Text()
			if feedUrlToShow == "" || feedUrlToShow == feedUrl {
				show.showFeedUrl(feedUrl)
			}
		}
	}
	return err
}

func (show *ShowDatabase) showFeedContent(feed *parsing.Feed) {

	fmt.Println(feed.Channel.Title)
	for _, item := range feed.Channel.ItemList {
		fmt.Println("\t" + item.Title)
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {
				fmt.Printf("\t\tDownloaded: %t\t"+enclosure.Url+"\n", enclosure.Downloaded)
			}
		}
	}
}

func (show *ShowDatabase) showFeedUrl(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			fmt.Println("Showing " + feedUrl)

			feed := parsing.Feed{}

			err := show.Repo.ReadInto("Feeds", feedUrl, &feed)

			if err == nil {
				show.showFeedContent(&feed)
			}
		}
	}
}
