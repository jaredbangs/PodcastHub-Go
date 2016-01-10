package actions

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"github.com/jaredbangs/go-download/download"
	"os"
	"path"
	"strings"
	"time"
)

type downloadFiles struct {
	Config       config.Configuration
	downloader   *download.Download
	downloadPath string
	repo         *repositories.FeedRepository
}

func NewDownload(config config.Configuration) *downloadFiles {

	d := &downloadFiles{Config: config}

	d.initializeRepo()
	d.initializeDownloader()

	return d
}

func (d *downloadFiles) DownloadAllNewFiles() error {

	for _, feedUrl := range d.repo.GetAllKeys() {
		d.DownloadNewFilesInFeed(feedUrl)
	}

	return nil
}

func (d *downloadFiles) DownloadFileInFeed(feedUrl string, enclosureUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.Read(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, false, enclosureUrl)
			}
		}
	}

	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
}

func (d *downloadFiles) DownloadNewFilesInFeed(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.Read(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, false, "")
			}
		}
	}

	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
}

func (d *downloadFiles) MarkAllNewFilesDownloaded() error {

	for _, feedUrl := range d.repo.GetAllKeys() {
		d.MarkAllNewFilesDownloadedInFeed(feedUrl)
	}

	return nil
}

func (d *downloadFiles) MarkAllNewFilesDownloadedInFeed(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.Read(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, true, "")
			}
		}
	}

	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
}

func (d *downloadFiles) downloadNewFilesInFeed(feed *parsing.Feed, feedUrl string, markOnly bool, specificEnclosureUrl string) {

	d.logFeedName(feed, feedUrl)

	for _, item := range feed.Channel.ItemList {
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 {

				shouldDownload := !enclosure.Downloaded

				if specificEnclosureUrl != "" {
					shouldDownload = specificEnclosureUrl == enclosure.Url
				}

				if shouldDownload {
					d.log("Downloading new file: " + enclosure.Url)

					if !markOnly {
						d.downloader.DownloadFile(enclosure.Url, d.downloadPath)
					}

					enclosure.Downloaded = true

					feed.UpdateEnclosure(enclosure)

					d.repo.Save(feedUrl, feed)
				}
			}
		}
	}
}

func (d *downloadFiles) initializeDownloader() {

	if d.downloader == nil {

		d.prepareDownloadPath()

		d.downloader = &download.Download{
			EnableLogging: true,
			LogAllHeaders: true,
			LogToFilePath: d.downloadPath,
		}

	}
}

func (d *downloadFiles) initializeRepo() {

	if d.repo == nil {
		d.repo = repositories.NewFeedRepository(d.Config)
	}
}

func (d *downloadFiles) logError(err error) {
	d.log(err.Error())
}

func (d *downloadFiles) logFeedName(feed *parsing.Feed, feedUrl string) {

	if feed.Channel.Title != "" {
		d.log("Feed: " + feed.Channel.Title)
	} else {
		d.log("Feed: " + feedUrl)
	}
}

func (d *downloadFiles) log(text string) {

	err := d.prepareDownloadPath()
	if err != nil {
		panic(err)
	}

	filePath := path.Join(d.downloadPath, "download.log")

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(text + "\n"); err != nil {
		panic(err)
	}

	fmt.Println(text)
}

func (d *downloadFiles) prepareDownloadPath() (err error) {

	if d.downloadPath == "" {
		parentDownloadPath := d.Config.DownloadPath

		t := time.Now()
		todayDirectoryName := t.Format("20060102")
		d.downloadPath = path.Join(parentDownloadPath, todayDirectoryName)

		err = os.MkdirAll(d.downloadPath, 0711)
		if err != nil {
			d.logError(err)
			return err
		}

		filePath := path.Join(d.downloadPath, "download.log")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			out, err := os.Create(filePath)
			if err != nil {
				d.logError(err)
				return err
			}
			defer out.Close()
		}
	}

	return nil
}
