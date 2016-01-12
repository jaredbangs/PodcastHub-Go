package files

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileProcessor struct {
	Filters     []FileProcessingFilter
	pathChanged bool
}

func (f *FileProcessor) WalkPath(path string) {
	filepath.Walk(path, f.processFile)

	for f.pathChanged {
		f.pathChanged = false
		filepath.Walk(path, f.processFile)
	}
}

func (f *FileProcessor) processFile(path string, info os.FileInfo, err error) error {

	for _, filter := range f.Filters {

		if filter.ShouldProcessFile(path, info) {

			pathChanged, err := filter.ProcessFile(path, info)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Processing error %s: %s\n", path, err)
				return err
			} else if pathChanged {
				f.pathChanged = true
				break
			}
		}
	}

	return err
}
