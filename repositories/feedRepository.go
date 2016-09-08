package repositories

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/go-repository/boltrepository"
	"log"
	"time"
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

func (r *FeedRepository) Clone() {

	log.Println("Opening target database")

	target := boltrepository.NewRepository(r.Config.RepositoryCloneTargetFile)
	target.GetObject = func(val []byte) interface{} {
		feed := &parsing.Feed{}
		target.Deserialize(val, &feed)
		return *feed
	}

	targetIds := boltrepository.NewRepository(r.Config.RepositoryCloneTargetFile)
	targetIds.GetObject = func(val []byte) interface{} {
		id := ""
		targetIds.Deserialize(val, &id)
		return id
	}

	r.feedsByIdBoltRepo.ForEach(r.feedsByIdBucketName, func(key string, val interface{}) {

		feed := val.(parsing.Feed)

		targetHasItem, _ := target.HasItem(r.feedsByIdBucketName, feed.Id)

		if !targetHasItem {
			target.Save(r.feedsByIdBucketName, feed.Id, feed)
			targetIds.Save(r.idsByUrlBucketName, feed.FeedUrl, feed.Id)
			log.Println("Target adding feed " + feed.FeedUrl)
		} else {
			log.Println("Target already contains feed " + feed.FeedUrl)
		}
	})
}

func (r *FeedRepository) DeleteById(id string) (err error) {

	err = r.feedsByIdBoltRepo.Delete(r.feedsByIdBucketName, id)

	return err
}

func (r *FeedRepository) DeleteByUrl(url string) (err error) {

	id, err := r.idsByUrlBoltRepo.Read(r.idsByUrlBucketName, url)

	if id == nil {
		err = fmt.Errorf("url %q not found", url)
	} else {
		err = r.DeleteById(id.(string))
		if err == nil {
			err = r.idsByUrlBoltRepo.Delete(r.idsByUrlBucketName, url)
		}
	}
	return err
}

func (r *FeedRepository) ForEach(action func(parsing.Feed)) {

	r.feedsByIdBoltRepo.ForEach(r.feedsByIdBucketName, func(key string, val interface{}) {

		feed := val.(parsing.Feed)
		action(feed)
	})
}

func (r *FeedRepository) GetAllIds() []string {

	allKeys := make([]string, 1)

	r.ForEach(func(feed parsing.Feed) {
		allKeys = append(allKeys, feed.Id)
	})

	return allKeys
}

func (r *FeedRepository) GetAllUrls() []string {

	allKeys := make([]string, 1)

	r.ForEach(func(feed parsing.Feed) {
		allKeys = append(allKeys, feed.FeedUrl)
	})

	return allKeys
}

func (r *FeedRepository) HasById(id string) (has bool, err error) {

	has, err = r.feedsByIdBoltRepo.HasItem(r.feedsByIdBucketName, id)

	return has, err
}

func (r *FeedRepository) HasByUrl(url string) (has bool, err error) {

	has, err = r.idsByUrlBoltRepo.HasItem(r.idsByUrlBucketName, url)

	return has, err
}

func (r *FeedRepository) ReadById(id string) (feed parsing.Feed, err error) {

	obj, err := r.feedsByIdBoltRepo.Read(r.feedsByIdBucketName, id)

	feed = obj.(parsing.Feed)

	return feed, err
}

func (r *FeedRepository) ReadByUrl(url string) (feed parsing.Feed, err error) {

	id, err := r.idsByUrlBoltRepo.Read(r.idsByUrlBucketName, url)

	if id == nil {
		err = fmt.Errorf("url %q not found", url)
	} else {
		feed, err = r.ReadById(id.(string))
	}
	return feed, err
}

func (r *FeedRepository) Save(feed *parsing.Feed) {
	feed.LastUpdated = time.Now()
	r.feedsByIdBoltRepo.Save(r.feedsByIdBucketName, feed.Id, feed)
	r.idsByUrlBoltRepo.Save(r.idsByUrlBucketName, feed.FeedUrl, feed.Id)
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
