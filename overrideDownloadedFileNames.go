package main

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/files"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	filepath.Walk(os.Args[1], processFile)
}

func processFile(path string, info os.FileInfo, err error) error {

	if strings.HasSuffix(path, ".mp3") {

		fmt.Println(path)
		fmt.Println(info.Size())

		reader := files.TagReader{}

		reader.PrintFileInfo(path)
	}

	return err
}
