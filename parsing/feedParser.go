package parsing

type FeedParser interface {
	parseFeedContent(content []byte) (Feed, error)
	parseFeedUrl(feedUrl string) (Feed, error)
}

func TryParse(content []byte, forceAllLinksParser bool) (feed Feed, parserUsed string, err error) {

	parser := RssParser{}
	parserUsed = "RssParser"

	feed, err = parser.ParseFeedContent(content)

	if !forceAllLinksParser && err == nil {
		return feed, parserUsed, err
	} else {

		parser := AllLinksParser{}
		parserUsed = "AllLinksParser"

		feed, err = parser.ParseFeedContent(content)

		return feed, parserUsed, err
	}
}
