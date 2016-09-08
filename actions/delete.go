package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/repositories"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type Delete struct {
	Config         config.Configuration
	downloadPath   string
	logFile        *os.File
	repo           *repositories.FeedRepository
	verboseLogging bool
}

func (d *Delete) DeleteFeed(feedUrl string, verbose bool) {

	d.verboseLogging = verbose

	d.initializeLog()

	if d.repo == nil {
		d.repo = repositories.NewFeedRepository(d.Config)
	}

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			log.Println("Getting " + feedUrl)

			err := d.repo.DeleteByUrl(feedUrl)

			if err != nil {
				log.Println("ERR\t" + err.Error())
			}
		}
	}
}

func (d *Delete) initializeLog() {

	if d.downloadPath == "" {
		d.prepareDownloadPath()

		t := time.Now()
		filePath := path.Join(d.downloadPath, t.Format("20060102-150405")+"-update.log")

		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		d.logFile = f

		log.SetOutput(io.MultiWriter(d.logFile, os.Stdout))
		log.Println("Initialized log")
	}
}

func (d *Delete) prepareDownloadPath() (err error) {

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
