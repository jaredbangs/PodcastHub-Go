package actions

import (
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"github.com/satori/go.uuid"

	"io/ioutil"
	"net/http"
)

type Subscribe struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func (s *Subscribe) Subscribe(feedUrl string) {

	if len(feedUrl) > 0 {
		fmt.Println("Getting " + feedUrl)

		response, httpErr := http.Get(feedUrl)

		if httpErr != nil {
			fmt.Println("ERR\t" + httpErr.Error())
		} else {

			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			s.addFeed(feedUrl, body)
		}
	}
}

func (s *Subscribe) addFeed(feedUrl string, content []byte) {

	if s.repo == nil {
		s.repo = repositories.NewFeedRepository(s.Config)
	}

	currentFeed, err := parsing.TryParse(content)

	if err == nil {

		hasFeed, err := s.repo.HasByUrl(feedUrl)

		if err == nil {

			if !hasFeed {
				u1 := uuid.NewV4()
				currentFeed.Id = u1.String()
				currentFeed.FeedUrl = feedUrl
				s.repo.Save(&currentFeed)
			}
		}
	}
}
