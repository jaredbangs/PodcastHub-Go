package parsing

import (
	"time"
)

type Enclosure struct {
	Downloaded         bool
	DownloadedAt       time.Time
	DownloadedFilePath string
	Length             int    `xml:"length,attr"`
	Type               string `xml:"type,attr"`
	Url                string `xml:"url,attr"`
}
