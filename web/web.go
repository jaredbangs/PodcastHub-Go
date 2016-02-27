package web

import (
	"encoding/json"
	"github.com/jaredbangs/PodcastHub/archive"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Web struct {
	Config   config.Configuration
	feedInfo []FeedInfo
	repo     *repositories.FeedRepository
}

type FeedInfo struct {
	Id          string
	LastUpdated time.Time
	Title       string
	Url         string
}

func (w *Web) Start() {

	w.initializeRepo()
	http.HandleFunc("/", w.index)
	http.HandleFunc("/feedArchiveStrategies", w.listFeedArchiveStrategies)
	http.HandleFunc("/feeds", w.listFeeds)
	http.HandleFunc("/feeds/", w.processFeed)
	http.HandleFunc("/items/", w.updateItem)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("web/scripts"))))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func (w *Web) getFeedInfo() []FeedInfo {

	if w.feedInfo == nil {
		w.feedInfo = make([]FeedInfo, 0)

		w.repo.ForEach(func(feed parsing.Feed) {

			feedInfo := FeedInfo{}
			feedInfo.Id = feed.Id
			feedInfo.LastUpdated = feed.LastUpdated
			feedInfo.Title = feed.Channel.Title
			feedInfo.Url = feed.FeedUrl

			w.feedInfo = append(w.feedInfo, feedInfo)
		})
	}
	return w.feedInfo
}

func (w *Web) index(rw http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("web/index.html")

	t.Execute(rw, nil)
}

func (w *Web) initializeRepo() {

	if w.repo == nil {
		w.repo = repositories.NewFeedRepository(w.Config)
	}
}

func (w *Web) listFeedArchiveStrategies(rw http.ResponseWriter, r *http.Request) {

	keys := make([]string, 0, len(archive.AvailableArchiveStrategies))
	for k := range archive.AvailableArchiveStrategies {
		keys = append(keys, k)
	}
	json.NewEncoder(rw).Encode(keys)
}

func (w *Web) listFeeds(rw http.ResponseWriter, r *http.Request) {

	json.NewEncoder(rw).Encode(w.getFeedInfo())
}

func (w *Web) processFeed(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Serve the resource.
		w.showFeed(rw, r)
		break
	case "POST":
		// Create a new record.
	case "PUT":
		// Update an existing record.
		w.updateFeed(rw, r)
		break
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}

}

func (w *Web) showFeed(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/feeds/"):]

	feed, _ := w.repo.ReadById(id)

	json.NewEncoder(rw).Encode(feed)
}

func (w *Web) updateFeed(rw http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var postedFeed parsing.Feed
	err := decoder.Decode(&postedFeed)
	if err != nil {

	} else {
		feed, _ := w.repo.ReadById(postedFeed.Id)

		feed.ArchivePath = postedFeed.ArchivePath
		feed.ArchiveStrategy = postedFeed.ArchiveStrategy

		w.repo.Save(&feed)

		json.NewEncoder(rw).Encode(feed)
	}
}

func (w *Web) updateItem(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		// Create a new record.
	case "PUT":
		// Update an existing record.

		decoder := json.NewDecoder(r.Body)
		var item parsing.Item
		err := decoder.Decode(&item)
		if err != nil {

		} else {
			feed, _ := w.repo.ReadById(item.FeedId)

			feed.UpdateItem(item)

			w.repo.Save(&feed)

			json.NewEncoder(rw).Encode(item)
		}

		break

	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}
