package actions

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Download struct {
	Config       config.Configuration
	downloadPath string
	repo         *repositories.FeedRepository
}

func (d *Download) DownloadAllNewFiles() error {

	d.initializeRepo()

	for _, feedUrl := range d.repo.GetAllKeys() {
		d.DownloadNewFilesInFeed(feedUrl)
	}

	return nil
}

func (d *Download) DownloadFileInFeed(feedUrl string, enclosureUrl string) {

	d.initializeRepo()

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.Read(feedUrl)

			if err == nil {
				for _, item := range feed.Channel.ItemList {
					for _, enclosure := range item.Enclosures {
						if len(enclosure.Url) != 0 && enclosure.Url == enclosureUrl {

							d.log("Downloading new file: " + enclosure.Url)

							d.downloadFile(enclosure.Url)

							enclosure.Downloaded = true

							feed.UpdateEnclosure(enclosure)

							d.repo.Save(feedUrl, &feed)
						}
					}
				}
			}
		}
	}

	d.prepareDownloadPath()
	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
}

func (d *Download) DownloadNewFilesInFeed(feedUrl string) {

	d.initializeRepo()

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.Read(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, false)
			}
		}
	}

	d.prepareDownloadPath()
	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
}

func (d *Download) MarkAllNewFilesDownloaded() error {

	d.initializeRepo()

	for _, feedUrl := range d.repo.GetAllKeys() {
		d.MarkAllNewFilesDownloadedInFeed(feedUrl)
	}

	return nil
}

func (d *Download) MarkAllNewFilesDownloadedInFeed(feedUrl string) {

	d.initializeRepo()

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed, err := d.repo.Read(feedUrl)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, true)
			}
		}
	}

	d.prepareDownloadPath()
	process := &ProcessDownloadedFiles{}
	process.ApplyAllFilters(d.downloadPath)
}

func (d *Download) TrimExtraPartsFromFileName(fileName string) (adjustedFileName string) {

	if strings.Contains(fileName, "?") {
		adjustedFileName = strings.Split(fileName, "?")[0]
	} else {
		adjustedFileName = fileName
	}

	return adjustedFileName
}

func (d *Download) downloadFile(url string) (err error) {

	//transport := &RedirectHandlingTransport{}
	//client := &http.Client{Transport: transport}

	resp, err := http.Get(url)
	if err != nil {
		d.logError(err)
		return err
	}
	defer resp.Body.Close()

	filePath, err := d.getTargetFilePath(url, resp)
	if err != nil {
		d.logError(err)
		return err
	}

	out, err := os.Create(filePath)
	if err != nil {
		d.logError(err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		d.logError(err)
		return err
	}

	d.log("Downloaded " + filePath)

	return nil
}

func (d *Download) downloadNewFilesInFeed(feed *parsing.Feed, feedUrl string, markOnly bool) {

	if feed.Channel.Title != "" {
		d.log("Feed: " + feed.Channel.Title)
	} else {
		d.log("Feed: " + feedUrl)
	}

	d.initializeRepo()

	for _, item := range feed.Channel.ItemList {
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 && !enclosure.Downloaded {

				d.log("Downloading new file: " + enclosure.Url)

				if !markOnly {
					d.downloadFile(enclosure.Url)
				}

				enclosure.Downloaded = true

				feed.UpdateEnclosure(enclosure)

				d.repo.Save(feedUrl, feed)
			}
		}
	}
}

func (d *Download) getTargetFilePath(url string, resp *http.Response) (filePath string, err error) {

	err = d.prepareDownloadPath()

	if url != resp.Request.URL.String() {
		url = resp.Request.URL.String()
	}

	fileNamePart := path.Base(url)

	if err == nil {

		if d.Config.LogAllHeaders {

			for k, v := range resp.Header {
				log.Println("Header:", k, "value:", v)
			}
		}

		contentDisposition := resp.Header.Get("Content-Disposition")

		if contentDisposition != "" {

			_, params, err := mime.ParseMediaType(contentDisposition)

			if err == nil && params["filename"] != "" {
				fileNamePart = params["filename"]
			}
			d.log(contentDisposition)
		}

		fileNamePart = d.TrimExtraPartsFromFileName(fileNamePart)

		filePath = path.Join(d.downloadPath, fileNamePart)
	}

	return filePath, err
}

func (d *Download) initializeRepo() {
	if d.repo == nil {
		d.repo = repositories.NewFeedRepository(d.Config)
	}
}

func (d *Download) logError(err error) {
	d.log(err.Error())
}

func (d *Download) log(text string) {

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

func (d *Download) prepareDownloadPath() (err error) {

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
