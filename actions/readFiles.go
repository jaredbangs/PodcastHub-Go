package actions

import (
	"github.com/jaredbangs/PodcastHub/files"
)

type ReadFiles struct {
}

func (process *ReadFiles) ApplyAllFilters(downloadPath string) {

	p := files.FileProcessor{}

	p.Filters = append(p.Filters, files.Id3InfoPrinter{})

	p.WalkPath(downloadPath)
}
