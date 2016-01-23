package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/files"
)

type ProcessDownloadedItems struct {
}

func (process *ProcessDownloadedItems) ApplyAllFilters(config config.Configuration) {

	p := files.ItemFileProcessor{Config: config}
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "This Agile Life"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "This American Life"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "Fighting In The War Room"})

	p.ProcessFiles()
}
