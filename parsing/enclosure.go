package parsing

type Enclosure struct {
	Downloaded         bool
	DownloadedFilePath string
	Url                string `xml:"url,attr"`
}
