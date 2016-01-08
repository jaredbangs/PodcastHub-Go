package repositories

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestSaveRoundTrip(t *testing.T) {

	filePath := randomString(10) + ".boltdb"

	config := &config.Configuration{RepositoryFile: filePath}

	repository := NewFeedRepository(*config)

	ch := &parsing.Channel{Title: "Hello"}
	sample := &parsing.Feed{Channel: *ch}

	repository.Save("Item1", sample)

	deserialized, _ := repository.Read("Item1")

	assert.Equal(t, "Hello", deserialized.Channel.Title, "Title should match")

	os.Remove(filePath)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
