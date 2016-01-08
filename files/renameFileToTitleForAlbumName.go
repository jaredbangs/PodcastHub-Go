package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RenameFileToTitleForAlbumName struct {
	AlbumName string
}

func (p RenameFileToTitleForAlbumName) ProcessFile(path string, info os.FileInfo) error {

	reader := TagReader{}
	err := reader.Read(path)

	if err == nil {
		dir := filepath.Dir(path)
		base := filepath.Base(path)
		newname := filepath.Join(dir, p.AlbumName+" - "+reader.Name+" - "+base)
		os.Rename(path, newname)

		fmt.Println(newname)
	}

	return nil
}

func (p RenameFileToTitleForAlbumName) ShouldProcessFile(path string, info os.FileInfo) bool {

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
