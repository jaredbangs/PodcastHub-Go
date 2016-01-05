package actions

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Download struct {
	Config       config.Configuration
	Repo         *boltrepository.Repository
	downloadPath string
}

func (d *Download) DownloadAllNewFiles() error {

	d.Repo = boltrepository.NewRepository(d.Config.RepositoryFile)

	file, err := os.Open(d.Config.SubscriptionFile)
	defer file.Close()

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			d.DownloadNewFilesInFeed(scanner.Text())
		}
	}
	return err
}

func (d *Download) DownloadNewFilesInFeed(feedUrl string) {

	if d.Repo == nil {
		d.Repo = boltrepository.NewRepository(d.Config.RepositoryFile)
	}

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed := parsing.Feed{}

			err := d.Repo.ReadInto("Feeds", feedUrl, &feed)

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

	d.Repo = boltrepository.NewRepository(d.Config.RepositoryFile)

	file, err := os.Open(d.Config.SubscriptionFile)
	defer file.Close()

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			d.MarkAllNewFilesDownloadedInFeed(scanner.Text())
		}
	}
	return err
}

func (d *Download) MarkAllNewFilesDownloadedInFeed(feedUrl string) {

	if d.Repo == nil {
		d.Repo = boltrepository.NewRepository(d.Config.RepositoryFile)
	}

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {

			feed := parsing.Feed{}

			err := d.Repo.ReadInto("Feeds", feedUrl, &feed)

			if err == nil {
				d.downloadNewFilesInFeed(&feed, feedUrl, true)
			}
		}
	}
}

func (d *Download) downloadFile(url string) (err error) {

	resp, err := http.Get(url)
	if err != nil {
		d.logError(err)
		return err
	}
	defer resp.Body.Close()

	filePath, err := d.getTargetFilePath(url, resp.Header.Get("Content-Disposition"))
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

	for _, item := range feed.Channel.ItemList {
		for _, enclosure := range item.Enclosures {
			if len(enclosure.Url) != 0 && !enclosure.Downloaded {

				d.log("Downloading new file: " + enclosure.Url)

				if !markOnly {
					d.downloadFile(enclosure.Url)
				}

				enclosure.Downloaded = true

				feed.UpdateEnclosure(enclosure)

				d.Repo.Save("Feeds", feedUrl, feed)
			}
		}
	}
}

func (d *Download) getTargetFilePath(url string, contentDisposition string) (filePath string, err error) {

	err = d.prepareDownloadPath()

	fileNamePart := path.Base(url)

	if err == nil {

		if contentDisposition != "" {

			_, params, err := mime.ParseMediaType(contentDisposition)

			if err == nil && params["filename"] != "" {
				filePath = path.Join(d.downloadPath, params["filename"])
			} else {
				filePath = path.Join(d.downloadPath, fileNamePart)
			}
			d.log(contentDisposition)
		} else {
			filePath = path.Join(d.downloadPath, fileNamePart)
		}
	}

	return filePath, err
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
