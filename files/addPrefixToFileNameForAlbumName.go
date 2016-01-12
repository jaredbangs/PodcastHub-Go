package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type AddPrefixToFileNameForAlbumName struct {
	AlbumName string
}

func (p AddPrefixToFileNameForAlbumName) ProcessFile(path string, info os.FileInfo) (bool, error) {

	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newname := filepath.Join(dir, p.AlbumName+" - "+base)
	os.Rename(path, newname)

	fmt.Println(newname)

	return true, nil
}

func (p AddPrefixToFileNameForAlbumName) ShouldProcessFile(path string, info os.FileInfo) bool {

	var shouldProcess = false

	base := filepath.Base(path)

	if strings.HasSuffix(path, ".mp3") && !strings.HasPrefix(base, p.AlbumName) {

		reader := TagReader{}

		err := reader.Read(path)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read %s: %s\n", path, err)
		}

		shouldProcess = strings.TrimSpace(reader.Album) == p.AlbumName

		if shouldProcess {
			reader.PrintFileInfo()
		}
	}

	return shouldProcess
}
