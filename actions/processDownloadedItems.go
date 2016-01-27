package actions

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/files"
)

type ProcessDownloadedItems struct {
}

func (process *ProcessDownloadedItems) ApplyAllFilters(config config.Configuration) {

	p := files.ItemFileProcessor{Config: config}
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "Fighting In The War Room"})
	//	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "NPR Politics Podcast"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "The Bike Shed"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "The Ventura Vineyard Podcast"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "This Agile Life"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "This American Life"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "Travel with Rick Steves"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "Turing-Incomplete"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForFeedName{FeedName: "Wait Wait... Don't Tell Me!"})

	p.ProcessFiles()
}
