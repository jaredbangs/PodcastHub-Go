package web

import (
	"encoding/json"
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/PodcastHub/parsing"
	"github.com/jaredbangs/PodcastHub/repositories"
	"html/template"
	"log"
	"net/http"
)

type Web struct {
	Config   config.Configuration
	feedInfo []FeedInfo
	repo     *repositories.FeedRepository
}

type FeedInfo struct {
	Id    string
	Title string
	Url   string
}

func (w *Web) Start() {

	w.initializeRepo()
	http.HandleFunc("/", w.index)
	http.HandleFunc("/feeds", w.listFeeds)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("web/scripts"))))
	http.HandleFunc("/show/", w.showFeed)

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func (w *Web) getFeedInfo() []FeedInfo {

	if w.feedInfo == nil {
		w.feedInfo = make([]FeedInfo, 0)

		w.repo.ForEach(func(feed parsing.Feed) {

			feedInfo := FeedInfo{}
			feedInfo.Id = feed.Id
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

func (w *Web) listFeeds(rw http.ResponseWriter, r *http.Request) {

	json.NewEncoder(rw).Encode(w.getFeedInfo())
}

func (w *Web) showFeed(rw http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("web/showFeed.html")

	id := r.URL.Path[len("/show/"):]

	feed, err := w.repo.ReadById(id)
	if err == nil {
		t.Execute(rw, feed)
	}
}
