package parsing

type Channel struct {
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	ItemList []Item `xml:"item"`
	Title    string `xml:"title"`
}
