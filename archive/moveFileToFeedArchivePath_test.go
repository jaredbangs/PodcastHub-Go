package archive

import (
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetMoveArchiveStrategyByName(t *testing.T) {

	a, _ := GetArchiveStrategyByName("moveFileToFeedArchivePath")

	assert.Equal(t, "moveFileToFeedArchivePath", a.GetName())
}

func TestInstantiateMoveArchive(t *testing.T) {

	strategyType := AvailableArchiveStrategies["moveFileToFeedArchivePath"]

	a := reflect.New(strategyType).Elem().Interface().(ArchiveStrategy)

	assert.Equal(t, "moveFileToFeedArchivePath", a.GetName())
}

func TestMoveArchiveAvailable(t *testing.T) {

	a, _ := NewMoveFileToFeedArchivePath(&parsing.Feed{ArchivePath: "/", ArchiveStrategy: "moveFileToFeedArchivePath"})

	strategyType := AvailableArchiveStrategies[a.GetName()]

	assert.Equal(t, "moveFileToFeedArchivePath", strategyType.Elem().Name())
}

func TestMoveFileArchiveNameWithArchivePath(t *testing.T) {

	a, _ := NewMoveFileToFeedArchivePath(&parsing.Feed{ArchivePath: "/", ArchiveStrategy: "moveFileToFeedArchivePath"})

	assert.Equal(t, "moveFileToFeedArchivePath", a.GetName())
}

func TestMoveFileArchiveNameWithoutArchivePath(t *testing.T) {

	a, _ := NewMoveFileToFeedArchivePath(&parsing.Feed{ArchiveStrategy: "moveFileToFeedArchivePath"})

	assert.Equal(t, "NullArchiveStrategy", a.GetName())
}

func TestMoveFileArchiveNameWithoutArchiveStrategy(t *testing.T) {

	a, _ := NewMoveFileToFeedArchivePath(&parsing.Feed{ArchivePath: "/"})

	assert.Equal(t, "NullArchiveStrategy", a.GetName())
}
