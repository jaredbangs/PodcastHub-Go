package archive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullArchiveAvailable(t *testing.T) {

	a := &NullArchiveStrategy{}

	strategyType := AvailableArchiveStrategies[a.GetName()]

	assert.Equal(t, "NullArchiveStrategy", strategyType.Elem().Name())
}

func TestNullArchiveName(t *testing.T) {

	a := &NullArchiveStrategy{}
	assert.Equal(t, "NullArchiveStrategy", a.GetName())
}
