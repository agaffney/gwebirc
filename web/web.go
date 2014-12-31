package web

import (
	"fmt"
	"github.com/agaffney/gwebirc/core"
	"github.com/agaffney/gwebirc/irc"
	"github.com/agaffney/gwebirc/types"
	"net/http"
)

type WebManager struct {
	Conf   *core.Config
	Irc    *irc.IrcManager
	Events []*types.Event
}

func (wm *WebManager) Start() {
	// Configure and start the web server
	if wm.Conf.Http.Enable_webui {
		http.Handle("/webui/", http.StripPrefix("/webui/", http.FileServer(http.Dir("./webui"))))
	}
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		api_handler(wm, w, r)
	})
	go http.ListenAndServe(fmt.Sprintf(":%d", wm.Conf.Http.Port), nil)
	go wm.event_poller()
}

func (wm *WebManager) event_poller() {
	// Read events from channel infinitely
	for e := range wm.Irc.Events {
		fmt.Println(e)
		wm.Events = append(wm.Events, e)
	}
}
