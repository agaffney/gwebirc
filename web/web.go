package web

import (
	"fmt"
	"github.com/agaffney/gwebirc/config"
	"github.com/agaffney/gwebirc/irc"
	"net/http"
)

type Web struct {
	Conf  *config.Config
	Conns []*irc.Connection
}

func (w *Web) Start() {
	if w.Conf.Http.Enable_webui {
		http.Handle("/webui/", http.StripPrefix("/webui/", http.FileServer(http.Dir("./webui"))))
	}
	http.HandleFunc("/api/", func(wr http.ResponseWriter, r *http.Request) {
		api_handler(w, wr, r)
	})
	go http.ListenAndServe(fmt.Sprintf(":%d", w.Conf.Http.Port), nil)
}
