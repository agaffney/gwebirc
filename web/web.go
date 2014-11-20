package web

import (
	"fmt"
	"github.com/agaffney/gwebirc/core"
	"github.com/agaffney/gwebirc/irc"
	"net/http"
)

type WebManager struct {
	Conf *core.Config
	Irc  *irc.IrcManager
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
	// Add our event handler and start the IRC manager
	wm.Irc.Set_event_callback(irc_event_callback)
	wm.Irc.Start()
}

func irc_event_callback(e *irc.Event, c *irc.Connection) {
	fmt.Println(e)
}
