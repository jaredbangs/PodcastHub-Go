package files

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/parsing"
	"os"
	"path/filepath"
	"strings"
)

type AddPrefixToFileNameForFeedName struct {
	FeedName string
}

func (p AddPrefixToFileNameForFeedName) ProcessItem(feed *parsing.Feed, item *parsing.Item, enclosure *parsing.Enclosure) (err error) {

	path := enclosure.DownloadedFilePath
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newname := filepath.Join(dir, p.FeedName+" - "+base)

	enclosure.DownloadedFilePath = newname

	os.Rename(path, newname)

	fmt.Println(newname)

	return nil
}

func (p AddPrefixToFileNameForFeedName) ShouldProcessItem(feed *parsing.Feed, item *parsing.Item, enclosure *parsing.Enclosure) bool {

	var shouldProcess = false

	if enclosure.Downloaded && enclosure.DownloadedFilePath != "" {

		path := enclosure.DownloadedFilePath
		base := filepath.Base(path)

		if strings.HasSuffix(path, ".mp3") && !strings.HasPrefix(base, p.FeedName) {

			shouldProcess = strings.TrimSpace(feed.Channel.Title) == p.FeedName
		}
	}

	return shouldProcess
}
