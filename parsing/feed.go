package parsing

type Feed struct {
	Channel      Channel `xml:"channel"`
	existingUrls map[string]int
}

func (feed *Feed) AddItem(item Item) {
	feed.Channel.ItemList = append(feed.Channel.ItemList, item)
	feed.existingUrls = nil
}

func (feed *Feed) ContainsEnclosureUrl(url string) bool {

	feed.populateExistingUrls()

	_, containsKey := feed.existingUrls[url]

	return containsKey
}

func (feed *Feed) UpdateEnclosure(enclosure Enclosure) {

	for _, item := range feed.Channel.ItemList {
		if item.Enclosures != nil {
			for i, existingEnclosure := range item.Enclosures {
				if enclosure.Url == existingEnclosure.Url {
					item.Enclosures = append(item.Enclosures[:i], item.Enclosures[i+1:]...)
					item.Enclosures = append(item.Enclosures, enclosure)
					break
				}
			}
		}
	}
}

func (feed *Feed) populateExistingUrls() {

	if feed.existingUrls == nil {

		feed.existingUrls = make(map[string]int)

		for _, item := range feed.Channel.ItemList {
			if item.Enclosures != nil {
				for _, enclosure := range item.Enclosures {
					feed.existingUrls[enclosure.Url] = 1
				}
			}
		}
	}
}