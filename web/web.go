package web

import (
	"github.com/jaredbangs/PodcastHub/config"
	"github.com/jaredbangs/go-repository/boltrepository"
	"html/template"
	"net/http"
)

type Web struct {
	Config config.Configuration
	Repo   *boltrepository.Repository
}

func (w *Web) Start() {

	w.Repo = boltrepository.NewRepository(w.Config.RepositoryFile)

	http.HandleFunc("/", w.handleListFeedsRequest)

	http.ListenAndServe(":8080", nil)
}

func (w *Web) handleListFeedsRequest(rw http.ResponseWriter, r *http.Request) {

	//feed := nil

	t, _ := template.ParseFiles("listFeeds.html")
	t.Execute(rw, nil)
}
