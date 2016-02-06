package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"github.com/jaredbangs/go-download/download"
	"io"
	"log"
	"net/http"
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

	if d.logFile != nil {
		d.logFile.Close()
	}
}

func (d *downloadFiles) DownloadAllNewFiles() error {

	for _, feedUrl := range d.repo.GetAllUrls() {
		d.DownloadNewFilesInFeed(feedUrl)
	}

	return nil
}

func (d *downloadFiles) DownloadFileInFeed(feedUrl string, enclosureUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.ReadByUrl(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, false, enclosureUrl)
			} else {
				log.Println(err)
			}
		}
	}

	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)

	p := &ProcessDownloadedItems{}
	p.ApplyAllFilters(d.Config)
}

func (d *downloadFiles) DownloadNewFilesInFeed(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.ReadByUrl(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, false, "")
			} else {
				log.Println(err)
			}
		}
	}

	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
	p := &ProcessDownloadedItems{}
	p.ApplyAllFilters(d.Config)
}

func (d *downloadFiles) MarkAllNewFilesDownloaded() error {

	for _, feedUrl := range d.repo.GetAllUrls() {
		d.MarkAllNewFilesDownloadedInFeed(feedUrl)
	}

	return nil
}

func (d *downloadFiles) MarkAllNewFilesDownloadedInFeed(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.ReadByUrl(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, true, "")
			} else {
				log.Println(err)
			}
		}
	}

	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
	p := &ProcessDownloadedItems{}
	p.ApplyAllFilters(d.Config)
}

func (d *downloadFiles) ShouldDownloadResponseFilter(resp *http.Response) bool {

	contentType := resp.Header.Get("Content-Type")

	if strings.Contains(contentType, "text") {
		return false
	} else if strings.Contains(contentType, "javascript") {
		return false
	} else if strings.Contains(contentType, "image/png") {
		return false
	} else if strings.Contains(contentType, "image/gif") {
		return false
	} else if strings.Contains(contentType, "image/jpeg") {
		return false
	} else if strings.Contains(contentType, "image/x-icon") {
		return false
	} else if strings.Contains(contentType, "application/rss+xml") {
		return false
	} else if strings.Contains(contentType, "application/xml") {
		return false
	} else if strings.Contains(contentType, "atom+xml") {
		return false
	}

	return true
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

						savedToPath, err := d.downloader.DownloadFile(enclosure.Url, d.downloadPath)

						if err == nil {
							enclosure.DownloadedFilePath = savedToPath
						}
					}

					enclosure.Downloaded = true
					enclosure.DownloadedAt = time.Now()

					feed.UpdateEnclosure(enclosure)

					d.repo.Save(feed)
				}
			}
		}
	}
}

func (d *downloadFiles) initializeDownloader() {

	if d.downloader == nil {

		d.prepareDownloadPath()

		d.downloader = &download.Download{
			EnableLogging:                true,
			LogAllHeaders:                true,
			ShouldDownloadResponseFilter: d.ShouldDownloadResponseFilter,
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
