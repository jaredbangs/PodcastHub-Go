package repositories

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
)

type LegacyFeedRepository struct {
	Config     config.Configuration
	bucketName string
	boltRepo   *boltrepository.Repository
}

func NewLegacyFeedRepository(config config.Configuration) *LegacyFeedRepository {
	r := &LegacyFeedRepository{
		Config:     config,
		bucketName: "Feeds",
	}
	r.initializeUnderlyingRepository()
	return r
}

func (r *LegacyFeedRepository) ForEach(action func(string, parsing.Feed)) {

	r.boltRepo.ForEach(r.bucketName, func(key string, val interface{}) {

		action(key, val.(parsing.Feed))
	})
}

func (r *LegacyFeedRepository) GetAllKeys() []string {

	allKeys := make([]string, 1)

	r.ForEach(func(key string, feed parsing.Feed) {
		allKeys = append(allKeys, key)
	})

	return allKeys
}

func (r *LegacyFeedRepository) Read(keyName string) (feed parsing.Feed, err error) {

	obj, err := r.boltRepo.Read(r.bucketName, keyName)

	feed = obj.(parsing.Feed)

	return feed, err
}

func (r *LegacyFeedRepository) Save(keyName string, feed *parsing.Feed) {
	r.boltRepo.Save(r.bucketName, keyName, feed)
}

func (r *LegacyFeedRepository) initializeUnderlyingRepository() {

	if r.boltRepo == nil {
		r.boltRepo = boltrepository.NewRepository(r.Config.RepositoryFile)

		r.boltRepo.GetObject = func(val []byte) interface{} {
			feed := &parsing.Feed{}
			r.boltRepo.Deserialize(val, &feed)
			return *feed
		}
	}
}
