package files

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileProcessor struct {
	Filters []FileProcessingFilter
}

func (f *FileProcessor) WalkPath(path string) {
	filepath.Walk(path, f.processFile)
}

func (f *FileProcessor) processFile(path string, info os.FileInfo, err error) error {

	for _, filter := range f.Filters {

		if filter.ShouldProcessFile(path, info) {
			err = filter.ProcessFile(path, info)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Processing error %s: %s\n", path, err)
				return err
			}
		}
	}

	return err
}
