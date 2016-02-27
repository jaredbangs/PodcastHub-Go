package archive

import (
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInstantiateDeleteArchive(t *testing.T) {

	strategyType := AvailableArchiveStrategies["deleteFile"]

	a := reflect.New(strategyType).Elem().Interface().(ArchiveStrategy)

	assert.Equal(t, "deleteFile", a.GetName())
}

func TestDeleteArchiveAvailable(t *testing.T) {

	a, _ := NewDeleteFile(&parsing.Feed{ArchiveStrategy: "deleteFile"})

	strategyType := AvailableArchiveStrategies[a.GetName()]

	assert.Equal(t, "deleteFile", strategyType.Elem().Name())
}

func TestDeleteFileArchiveNameWithArchivePath(t *testing.T) {

	a, _ := NewDeleteFile(&parsing.Feed{ArchiveStrategy: "deleteFile"})

	assert.Equal(t, "deleteFile", a.GetName())
}

func TestDeleteFileArchiveNameWithoutArchivePath(t *testing.T) {

	a, _ := NewDeleteFile(&parsing.Feed{})

	assert.Equal(t, "NullArchiveStrategy", a.GetName())
}
