package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimExtraPartsFromFileNameShouldNotRemoveAnythingElse(t *testing.T) {

	d := Download{}

	original := "otm010816pod.mp3"

	adjusted := d.TrimExtraPartsFromFileName(original)

	assert.Equal(t, original, adjusted)
}

func TestTrimExtraPartsFromFileNameShouldRemoveQueryParams(t *testing.T) {

	d := Download{}

	original := "otm010816pod.mp3?downloadId=568e0077ccc09d0e_GopMLehT_00000001JM1"

	adjusted := d.TrimExtraPartsFromFileName(original)

	assert.Equal(t, "otm010816pod.mp3", adjusted)
}
