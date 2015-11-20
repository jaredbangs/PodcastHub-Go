package files

import (
	"fmt"
	"github.com/ascherkus/go-id3/src/id3"
	"os"
	"strings"
)

type TagReader struct {
	Name   string
	Artist string
	Album  string
	Year   string
	Track  string
	Disc   string
	Genre  string
	Length string
}

func (r *TagReader) PrintFileInfo() {

	r.printIfNotEmpty("Name", r.Name)
	r.printIfNotEmpty("Artisit", r.Artist)
	r.printIfNotEmpty("Album", r.Album)
	r.printIfNotEmpty("Year", r.Year)
	r.printIfNotEmpty("Track", r.Track)
	r.printIfNotEmpty("Disc", r.Disc)
	r.printIfNotEmpty("Genre", r.Genre)
	r.printIfNotEmpty("Length", r.Length)

	fmt.Println()
}

func (r *TagReader) Read(path string) error {

	var fd, err = os.Open(path)

	defer fd.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s: %s\n", path, err)
		return err
	}

	defer func() {
		if e := recover(); e != nil {
		}
	}()

	file := id3.Read(fd)

	if file == nil {
		fmt.Fprintf(os.Stderr, "Could not read ID3 information from %s\n", path)
	} else {
		r.copyInfoFromId3File(file)
	}

	return nil
}

func (r *TagReader) ReadFileAndPrintFileInfo(path string) {

	err := r.Read(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read %s: %s\n", path, err)
	}

	r.PrintFileInfo()
}

func (r *TagReader) copyInfoFromId3File(file *id3.File) {

	r.Name = file.Name
	r.Artist = file.Artist
	r.Album = file.Album
	r.Year = file.Year
	r.Track = file.Track
	r.Disc = file.Disc
	r.Genre = file.Genre
	r.Length = file.Length
}

func (r *TagReader) printIfNotEmpty(label string, value string) {

	if strings.TrimSpace(value) != "" {

		fmt.Printf("%s\t%s\n", label, value)
	}
}
