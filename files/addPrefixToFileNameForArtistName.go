package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type AddPrefixToFileNameForArtistName struct {
	ArtistName string
}

func (p AddPrefixToFileNameForArtistName) ProcessFile(path string, info os.FileInfo) error {

	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newname := filepath.Join(dir, p.ArtistName+" - "+base)
	os.Rename(path, newname)

	fmt.Println(newname)

	return nil
}

func (p AddPrefixToFileNameForArtistName) ShouldProcessFile(path string, info os.FileInfo) bool {

	var shouldProcess = false

	base := filepath.Base(path)

	if strings.HasSuffix(path, ".mp3") && !strings.HasPrefix(base, p.ArtistName) {

		reader := TagReader{}

		err := reader.Read(path)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read %s: %s\n", path, err)
		}

		shouldProcess = strings.TrimSpace(reader.Artist) == p.ArtistName

		if shouldProcess {
			reader.PrintFileInfo()
		}
	}

	return shouldProcess
}
