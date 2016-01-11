package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"github.com/jaredbangs/go-download/download"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type downloadFiles struct {
	Config       config.Configuration
	downloader   *download.Download
	downloadPath string
	logFile      *os.File
	repo         *repositories.FeedRepository
}

func NewDownload(config config.Configuration) *downloadFiles {

	d := &downloadFiles{Config: config}

	d.initializeRepo()
	d.initializeDownloader()
	d.initializeLog()

	return d
}

func (d *downloadFiles) Close() {
	//TODO: close log file here, always use Defer to call this from any endpoints
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
					log.Println("Downloading new file: " + enclosure.Url)

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
		}

	}
}

func (d *downloadFiles) initializeLog() {

	t := time.Now()
	filePath := path.Join(d.downloadPath, t.Format("20060102-150405")+"-download.log")

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	d.logFile = f

	log.SetOutput(io.MultiWriter(d.logFile, os.Stdout))
	log.Println("Initialized log")
}

func (d *downloadFiles) initializeRepo() {

	if d.repo == nil {
		d.repo = repositories.NewFeedRepository(d.Config)
	}
}

func (d *downloadFiles) logFeedName(feed *parsing.Feed, feedUrl string) {

	if feed.Channel.Title != "" {
		log.Println("Feed: " + feed.Channel.Title)
	} else {
		log.Println("Feed: " + feedUrl)
	}
}

func (d *downloadFiles) prepareDownloadPath() (err error) {

	if d.downloadPath == "" {
		parentDownloadPath := d.Config.DownloadPath

		t := time.Now()
		todayDirectoryName := t.Format("20060102")
		d.downloadPath = path.Join(parentDownloadPath, todayDirectoryName)

		err = os.MkdirAll(d.downloadPath, 0711)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	}

	return nil
}
