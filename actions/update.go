package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type update struct {
	Config         config.Configuration
	downloadPath   string
	logFile        *os.File
	repo           *repositories.FeedRepository
	verboseLogging bool
}

func NewUpdate(config config.Configuration) *update {

	u := &update{Config: config}

	u.initializeLog()
	u.repo = repositories.NewFeedRepository(u.Config)

	return u
}

func (update *update) Update() error {

	for _, feedUrl := range update.repo.GetAllUrls() {
		update.UpdateFeed(feedUrl, false)
	}

	return nil
}

func (update *update) UpdateFeed(feedUrl string, verbose bool) {

	update.verboseLogging = verbose

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			log.Println("Getting " + feedUrl)

			response, httpErr := http.Get(feedUrl)

			if httpErr != nil {

				log.Println("ERR\t" + httpErr.Error())
			} else {

				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)
				update.recordFeedInfo(feedUrl, body)
			}
		}
	}
}

func (u *update) initializeLog() {

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

func (u *update) prepareDownloadPath() (err error) {

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

func (update *update) recordFeedInfo(feedUrl string, content []byte) {

	if update.repo == nil {
		update.repo = repositories.NewFeedRepository(update.Config)
	}

	if update.verboseLogging {
		//		log.Println(string(content))
	}

	forceAllLinksParser := false
	tempFeedRecord, err := update.repo.ReadByUrl(feedUrl)
	if err == nil && tempFeedRecord.ForceAllLinksParser {
		forceAllLinksParser = true
	}

	currentFeed, parserUsed, err := parsing.TryParse(content, forceAllLinksParser)
	log.Println("Parser used: " + parserUsed)

	if err == nil {

		feedRecord, err := update.repo.ReadByUrl(feedUrl)

		if err == nil {

			update.updateChannelInfo(&feedRecord, &currentFeed.Channel)

			for _, item := range currentFeed.Channel.ItemList {

				hasUrl := true

				for _, enclosure := range item.Enclosures {
					if hasUrl {
						hasUrl = feedRecord.ContainsEnclosureUrl(enclosure.Url)
						if hasUrl && update.verboseLogging {
							log.Println("Already has item: " + item.Title + " : " + enclosure.Url)
							log.Println(enclosure.Url)
							log.Println("IsArchived: " + strconv.FormatBool(item.IsArchived) + " IsToBeArchived: " + strconv.FormatBool(item.IsToBeArchived))
							log.Println("DownloadedAt: " + enclosure.DownloadedAt.Format("2006-01-02T15:04:05.999999-07:00") + " DownloadedFilePath: " + enclosure.DownloadedFilePath)
						}
					}
				}

				if !hasUrl {
					feedRecord.AddItem(item)
					for _, enclosure := range item.Enclosures {
						log.Println("Added item: " + item.Title + " : " + enclosure.Url)
					}
				}
			}

			feedRecord.Channel.ReprocessExistingInfo(false)
			feedRecord.MakeSureAllItemsHaveIds()

			update.repo.Save(&feedRecord)
		} else {
			log.Println("Error reading existing feed record")
			log.Println(err)
		}
	} else {
		log.Println("Error parsing feed")
		log.Println(err)
	}
}

func (u *update) updateChannelInfo(existingFeed *parsing.Feed, currentChannel *parsing.Channel) {

	if u.verboseLogging {
		log.Println("Updating channel Info")
	}

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
