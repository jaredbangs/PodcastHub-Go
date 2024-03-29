package parsing

import (
	"log"
	"path/filepath"
	"time"
)

type Enclosure struct {
	Downloaded          bool
	DownloadedAt        time.Time
	DownloadedDirectory string
	DownloadedFilePath  string
	Length              int    `xml:"length,attr"`
	Type                string `xml:"type,attr"`
	Url                 string `xml:"url,attr"`
}

func (e *Enclosure) UpdateCalculatedInfo(logInfo bool) {

	if e.DownloadedFilePath != "" && e.DownloadedDirectory == "" {

		if logInfo {
			log.Println("Needs DownloadedDirectory")
		}

		downloadedDirectory := filepath.Dir(e.DownloadedFilePath)

		if downloadedDirectory != e.DownloadedDirectory {
			log.Println("Updating dl directory to " + downloadedDirectory + " for " + e.DownloadedFilePath)
			e.DownloadedDirectory = downloadedDirectory
		}
	} else {
		if logInfo {
			log.Println("Already has DownloadedDirectory")
		}
	}
}
