package files

import (
	"os"
)

type FileProcessingFilter interface {
	ProcessFile(path string, info os.FileInfo) (pathChanged bool, err error)
	ShouldProcessFile(path string, info os.FileInfo) bool
}
