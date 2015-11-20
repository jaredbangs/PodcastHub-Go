package main

import (
	"github.com/jaredbangs/PodcastHub/files"
	"os"
)

func main() {

	p := files.FileProcessor{}

	//p.Filters = append(p.Filters, files.Id3InfoPrinter{})

	p.Filters = append(p.Filters, files.AddPrefixToFileNameForAlbumName{AlbumName: "Turing-Incomplete"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForAlbumName{AlbumName: "Ventura Vineyard"})

	p.Filters = append(p.Filters, files.AddPrefixToFileNameForArtistName{ArtistName: "This American Life"})
	p.Filters = append(p.Filters, files.AddPrefixToFileNameForArtistName{ArtistName: "MadhouseLive.com"})

	p.WalkPath(os.Args[1])
}
