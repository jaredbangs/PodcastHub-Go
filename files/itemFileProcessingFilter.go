package files

import (
	"github.com/jaredbangs/PodcastHub/parsing"
)

type ItemFileProcessingFilter interface {
	ProcessItem(feed *parsing.Feed, item *parsing.Item, enclosure *parsing.Enclosure) (err error)
	ShouldProcessItem(feed *parsing.Feed, item *parsing.Item, enclosure *parsing.Enclosure) bool
}
