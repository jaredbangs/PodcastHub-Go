package archive

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestDeleteFile(t *testing.T) {

	filePath := "testFiles/testFileToDelete"
	createFileForTesting(filePath)

	m := &FileMover{}
	assert.Equal(t, true, m.FileExists(filePath))
	m.DeleteFile(filePath)

	assert.Equal(t, false, m.FileExists(filePath))
}

func TestFileDoesNotExist(t *testing.T) {

	m := &FileMover{}

	assert.Equal(t, false, m.FileExists("AAAfileMover.go"))
}

func TestFileExists(t *testing.T) {

	m := &FileMover{}

	assert.Equal(t, true, m.FileExists("fileMover.go"))
}

func TestMoveFile(t *testing.T) {

	filePath := "testFiles/testFileToMove"
	createFileForTesting(filePath)

	m := &FileMover{}
	assert.Equal(t, true, m.FileExists(filePath))

	destinationFilePath := "testFilesDestination/testFileToMove2"
	m.MoveFile(filePath, destinationFilePath)
	assert.Equal(t, false, m.FileExists(filePath))
	assert.Equal(t, true, m.FileExists(destinationFilePath))

	m.DeleteFile(destinationFilePath)
	m.DeleteFile(filePath)
}

func TestMoveFileIntoDirectory(t *testing.T) {

	filePath := "testFiles/testFileToMove"
	createFileForTesting(filePath)

	m := &FileMover{}
	assert.Equal(t, true, m.FileExists(filePath))

	destinationDirectory := "testFilesDestination"
	m.MoveFileIntoDirectory(filePath, destinationDirectory)
	assert.Equal(t, false, m.FileExists(filePath))

	destinationFilePath := filepath.Join(destinationDirectory, "testFileToMove")
	assert.Equal(t, true, m.FileExists(destinationFilePath))

	m.DeleteFile(destinationFilePath)
	m.DeleteFile(filePath)
}

func TestMoveFileIntoDirectoryWithSlash(t *testing.T) {

	filePath := "testFiles/testFileToMove"
	createFileForTesting(filePath)

	m := &FileMover{}
	assert.Equal(t, true, m.FileExists(filePath))

	destinationDirectory := "testFilesDestination/"
	resultPath := m.MoveFileIntoDirectory(filePath, destinationDirectory)
	assert.Equal(t, false, m.FileExists(filePath))

	destinationFilePath := filepath.Join(destinationDirectory, "testFileToMove")
	assert.Equal(t, destinationFilePath, resultPath)
	assert.Equal(t, true, m.FileExists(destinationFilePath))

	m.DeleteFile(destinationFilePath)
	m.DeleteFile(filePath)
}

func TestMoveFileIntoDirectoryWithSpace(t *testing.T) {

	filePath := "testFiles/testFileToMove"
	createFileForTesting(filePath)

	m := &FileMover{}
	assert.Equal(t, true, m.FileExists(filePath))

	destinationDirectory := "testFiles Destination Space"
	resultPath := m.MoveFileIntoDirectory(filePath, destinationDirectory)
	assert.Equal(t, false, m.FileExists(filePath))

	destinationFilePath := filepath.Join(destinationDirectory, "testFileToMove")
	assert.Equal(t, destinationFilePath, resultPath)
	assert.Equal(t, true, m.FileExists(destinationFilePath))

	m.DeleteFile(destinationFilePath)
	m.DeleteFile(filePath)
}

func createFileForTesting(filePath string) {

	d1 := []byte("hello")
	ioutil.WriteFile(filePath, d1, 0644)
}
