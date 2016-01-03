package parsing

type FeedParser interface {
	parseFeedContent(content []byte) (Feed, error)
	parseFeedUrl(feedUrl string) (Feed, error)
}

func TryParse(content []byte) (feed Feed, err error) {

	parser := RssParser{}

	feed, err = parser.ParseFeedContent(content)

	if err == nil {
		return feed, err
	} else {

		parser := AllLinksParser{}

		feed, err = parser.ParseFeedContent(content)

		return feed, err
	}
}
