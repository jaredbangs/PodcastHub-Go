package actions

import (
	"bufio"
	"fmt"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/repositories"
	"github.com/satori/go.uuid"
	"os"
	"strings"
)

type Reorg struct {
	Config     config.Configuration
	legacyRepo *repositories.LegacyFeedRepository
	repo       *repositories.FeedRepository
}

func (u *Reorg) CopyFromOldToNew() error {

	if u.repo == nil {
		u.repo = repositories.NewFeedRepository(u.Config)
		u.legacyRepo = repositories.NewLegacyFeedRepository(u.Config)
	}

	file, err := os.Open(u.Config.SubscriptionFile)
	defer file.Close()

	if err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			u.updateFeed(scanner.Text())
		}
	}

	return err
}

func (u *Reorg) updateFeed(feedUrl string) {

	if len(feedUrl) > 0 {
		if !strings.HasPrefix(feedUrl, "#") {
			fmt.Println("Processing " + feedUrl)

			feedRecord, _ := u.legacyRepo.Read(feedUrl)

			if feedRecord.Id == "" {
				u1 := uuid.NewV4()
				feedRecord.Id = u1.String()
			}

			if feedRecord.FeedUrl == "" {
				feedRecord.FeedUrl = feedUrl
			}

			u.repo.Save(&feedRecord)
		}
	}
}
