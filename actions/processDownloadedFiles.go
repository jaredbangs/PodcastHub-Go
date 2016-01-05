package actions

import (
	"github.com/jaredbangs/PodcastHub/files"
)

type ProcessDownloadedFiles struct {
}

func (process *ProcessDownloadedFiles) ApplyAllFilters(downloadPath string) {

	p := files.FileProcessor{}

	//p.Filters = append(p.Filters, files.Id3InfoPrinter{})

	p.Filters = append(p.Filters, files.AddPrefixToFileNameForAlbumName{AlbumName: "Turing-Incomplete"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForAlbumName{AlbumName: "Ventura Vineyard"})

	p.Filters = append(p.Filters, files.AddPrefixToFileNameForArtistName{ArtistName: "This American Life"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForArtistName{ArtistName: "MadhouseLive.com"})

	p.WalkPath(downloadPath)
}
