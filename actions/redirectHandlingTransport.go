package actions

import (
	"log"
	"net/http"
)

type RedirectHandlingTransport struct {
	RedirectedTo string
	Transport    http.RoundTripper
}

func (l RedirectHandlingTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := l.Transport
	if t == nil {
		t = http.DefaultTransport
	}
	resp, err = t.RoundTrip(req)
	if err != nil {
		return
	}
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		log.Println("Request for", req.URL, "redirected with status", resp.StatusCode)
		for k, v := range resp.Header {
			log.Println("redirect header key:", k, "value:", v)
		}
		l.RedirectedTo = resp.Header.Get("Location")
	}
	return
}
