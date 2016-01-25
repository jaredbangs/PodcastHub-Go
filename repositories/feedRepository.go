package repositories

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
)

type FeedRepository struct {
	Config              config.Configuration
	feedsByIdBoltRepo   *boltrepository.Repository
	feedsByIdBucketName string
	idsByUrlBoltRepo    *boltrepository.Repository
	idsByUrlBucketName  string
}

func NewFeedRepository(config config.Configuration) *FeedRepository {

	r := &FeedRepository{
		Config:              config,
		feedsByIdBucketName: "FeedsById",
		idsByUrlBucketName:  "IdsByUrl",
	}

	r.initializeUnderlyingRepository()

	return r
}

func (r *FeedRepository) ForEach(action func(string, parsing.Feed)) {

	r.feedsByIdBoltRepo.ForEach(r.feedsByIdBucketName, func(key string, val interface{}) {

		feed := val.(parsing.Feed)
		action(feed.FeedUrl, feed)
	})
}

func (r *FeedRepository) GetAllKeys() []string {

	allKeys := make([]string, 1)

	r.ForEach(func(key string, feed parsing.Feed) {
		allKeys = append(allKeys, key)
	})

	return allKeys
}

func (r *FeedRepository) ReadById(id string) (feed parsing.Feed, err error) {

	obj, err := r.feedsByIdBoltRepo.Read(r.feedsByIdBucketName, id)

	feed = obj.(parsing.Feed)

	return feed, err
}

func (r *FeedRepository) ReadByUrl(url string) (feed parsing.Feed, err error) {

	id, err := r.idsByUrlBoltRepo.Read(r.idsByUrlBucketName, url)

	feed, err = r.ReadById(id.(string))

	return feed, err
}

func (r *FeedRepository) Save(id string, feed *parsing.Feed) {
	r.feedsByIdBoltRepo.Save(r.feedsByIdBucketName, id, feed)
	r.idsByUrlBoltRepo.Save(r.idsByUrlBucketName, feed.FeedUrl, id)
}

func (r *FeedRepository) initializeUnderlyingRepository() {

	if r.feedsByIdBoltRepo == nil {
		r.feedsByIdBoltRepo = boltrepository.NewRepository(r.Config.RepositoryFile)

		r.feedsByIdBoltRepo.GetObject = func(val []byte) interface{} {
			feed := &parsing.Feed{}
			r.feedsByIdBoltRepo.Deserialize(val, &feed)
			return *feed
		}
	}

	if r.idsByUrlBoltRepo == nil {
		r.idsByUrlBoltRepo = boltrepository.NewRepository(r.Config.RepositoryFile)

		r.idsByUrlBoltRepo.GetObject = func(val []byte) interface{} {
			id := ""
			r.idsByUrlBoltRepo.Deserialize(val, &id)
			return id
		}
	}
}
