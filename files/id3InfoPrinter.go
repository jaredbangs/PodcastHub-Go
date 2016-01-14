package files

import (
	"fmt"
	"os"
	"strings"
)

type Id3InfoPrinter struct {
}

func (p Id3InfoPrinter) ProcessFile(path string, info os.FileInfo) (bool, error) {

	fmt.Println(path)
	fmt.Println(info.Size())

	reader := TagReader{}

	reader.ReadFileAndPrintFileInfo(path)

	return false, nil
}

func (p Id3InfoPrinter) ShouldProcessFile(path string, info os.FileInfo) bool {
	return strings.HasSuffix(path, ".mp3")
}
