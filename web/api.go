package web

import (
	"encoding/json"
	"github.com/agaffney/gwebirc/irc"
	"net/http"
	"strconv"
	"strings"
)

var handlers = map[string]func(*WebManager, http.ResponseWriter, *http.Request, []string){
	"connections": handle_connections,
	"events":      handle_events,
}

func api_handler(wm *WebManager, w http.ResponseWriter, r *http.Request) {
	url := strings.TrimPrefix(r.URL.Path, "/api/")
	url_parts := strings.Split(url, "/")
	if fn, ok := handlers[url_parts[0]]; ok {
		fn(wm, w, r, url_parts)
	} else {
		// Return a 404
		http.NotFound(w, r)
	}
}

func handle_events(wm *WebManager, w http.ResponseWriter, r *http.Request, params []string) {
	j := json.NewEncoder(w)
	var events []*irc.Event
	switch len(params) {
	case 1:
		events = wm.Events
	case 2:
		// We can cheat a bit here, since we know the implementation details
		id, _ := strconv.ParseUint(params[1], 10, 64)
		if int(id) >= len(wm.Events) {
			// Implement long polling by waiting on event
		}
		// Grab all events starting at the index after the provided event ID
		events = wm.Events[id:]
	default:
		http.Error(w, "Unexpected number of arguments", 400)
		return
	}
	j.Encode(events)
}

func handle_connections(wm *WebManager, w http.ResponseWriter, r *http.Request, params []string) {
	j := json.NewEncoder(w)
	switch len(params) {
	case 1:
		// List connections
		j.Encode(wm.Irc.Conns)
	case 2:
		// List specific connection by name
		found_conn := false
		for _, conn := range wm.Irc.Conns {
			if params[1] == conn.Name {
				j.Encode(conn)
				found_conn = true
				break
			}
		}
		if !found_conn {
			http.NotFound(w, r)
			return
		}
	case 3:
		fallthrough
	case 4:
		var conn *irc.Connection
		for _, c := range wm.Irc.Conns {
			if params[1] == c.Name {
				conn = c
				break
			}
		}
		if conn == nil {
			http.NotFound(w, r)
			return
		}
		switch params[2] {
		case "join":
			conn.Join(params[3])
		case "part":
			conn.Part(params[3])
		case "privmsg":
			target := params[3]
			msg := r.PostFormValue("msg")
			conn.Privmsg(target, msg)
		default:
			http.Error(w, "No such method '"+params[2]+"' for connection", 400)
			return
		}
	default:
		http.NotFound(w, r)
		return
	}
}
