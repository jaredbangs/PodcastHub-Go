package actions

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Update struct {
	Config       config.Configuration
	downloadPath string
	logFile      *os.File
	repo         *repositories.FeedRepository
}

func (update *Update) Update() error {

	update.initializeLog()

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

	update.initializeLog()

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

func (u *Update) initializeLog() {

	if u.downloadPath == "" {
		u.prepareDownloadPath()

		t := time.Now()
		filePath := path.Join(u.downloadPath, t.Format("20060102-150405")+"-update.log")

		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		u.logFile = f

		log.SetOutput(io.MultiWriter(u.logFile, os.Stdout))
		log.Println("Initialized log")
	}
}

func (u *Update) prepareDownloadPath() (err error) {

	if u.downloadPath == "" {
		parentDownloadPath := u.Config.DownloadPath

		t := time.Now()
		todayDirectoryName := t.Format("20060102")
		u.downloadPath = path.Join(parentDownloadPath, todayDirectoryName)

		err = os.MkdirAll(u.downloadPath, 0711)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	}

	return nil
}

func (update *Update) recordFeedInfo(feedUrl string, content []byte) {

	if update.repo == nil {
		update.repo = repositories.NewFeedRepository(update.Config)
	}

	currentFeed, err := parsing.TryParse(content)

	if err == nil {

		feedRecord, err := update.repo.ReadByUrl(feedUrl)

		if err == nil {

			update.updateChannelInfo(&feedRecord, &currentFeed.Channel)

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

			feedRecord.Channel.ReprocessExistingInfo(false)
			feedRecord.MakeSureAllItemsHaveIds()

			update.repo.Save(&feedRecord)
		} else {
			log.Println(err)
		}
	}
}

func (u *Update) updateChannelInfo(existingFeed *parsing.Feed, currentChannel *parsing.Channel) {

	if existingFeed.Channel.Description != currentChannel.Description && currentChannel.Description != "" {
		existingFeed.Channel.Description = currentChannel.Description
	}

	if existingFeed.Channel.Generator != currentChannel.Generator && currentChannel.Generator != "" {
		existingFeed.Channel.Generator = currentChannel.Generator
	}

	// Images        []Image

	if existingFeed.Channel.LastBuildDate != currentChannel.LastBuildDate && currentChannel.LastBuildDate != "" {
		existingFeed.Channel.LastBuildDate = currentChannel.LastBuildDate
	}

	if existingFeed.Channel.Link != currentChannel.Link && currentChannel.Link != "" {
		existingFeed.Channel.Link = currentChannel.Link
	}

	if existingFeed.Channel.Subtitle != currentChannel.Subtitle && currentChannel.Subtitle != "" {
		existingFeed.Channel.Subtitle = currentChannel.Subtitle
	}

	if existingFeed.Channel.Summary != currentChannel.Summary && currentChannel.Summary != "" {
		existingFeed.Channel.Summary = currentChannel.Summary
	}

	if existingFeed.Channel.Title != currentChannel.Title && currentChannel.Title != "" {
		existingFeed.Channel.Title = currentChannel.Title
	}

}
