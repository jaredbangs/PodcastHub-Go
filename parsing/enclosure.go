package parsing

type Enclosure struct {
	Downloaded bool
	Url        string `xml:"url,attr"`
}
