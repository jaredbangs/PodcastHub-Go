package archive

import (
	"io"
	"os"
	"path/filepath"
)

type FileMover struct {
}

func (m *FileMover) CopyFile(sourceFilePath, destinationFilePath string) error {
	in, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func (m *FileMover) DeleteFile(filePath string) {
	os.Remove(filePath)
}

func (m *FileMover) FileExists(filePath string) bool {

	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}

func (m *FileMover) MoveFile(sourceFilePath string, destinationFilePath string) {

	if sourceFilePath != destinationFilePath {
		// Can't just use os.Rename, because it doesn't work across devices

		err := m.CopyFile(sourceFilePath, destinationFilePath)

		if err == nil {
			m.DeleteFile(sourceFilePath)
		}
	}
}

func (m *FileMover) MoveFileIntoDirectory(filePath string, directoryName string) string {

	fileName := filepath.Base(filePath)
	archiveFilePath := filepath.Join(directoryName, fileName)
	m.MoveFile(filePath, archiveFilePath)

	return archiveFilePath
}
