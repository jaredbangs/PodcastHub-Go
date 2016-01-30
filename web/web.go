package web

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"html/template"
	"log"
	"net/http"
)

type Web struct {
	Config config.Configuration
	repo   *repositories.FeedRepository
}

func (w *Web) Start() {

	w.initializeRepo()

	http.HandleFunc("/", w.handleListFeedsRequest)
	http.HandleFunc("/show/", w.showFeed)

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func (w *Web) handleListFeedsRequest(rw http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("web/listFeeds.html")

	w.repo.ForEach(func(feed parsing.Feed) {
		if feed.Channel.Title != "" {
			t.Execute(rw, feed)
		}
	})
}

func (w *Web) initializeRepo() {

	if w.repo == nil {
		w.repo = repositories.NewFeedRepository(w.Config)
	}
}

func (w *Web) showFeed(rw http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("web/showFeed.html")

	id := r.URL.Path[len("/show/"):]

	feed, err := w.repo.ReadById(id)
	if err == nil {
		t.Execute(rw, feed)
	}
}
