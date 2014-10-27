package core

import (
	"encoding/json"
	"github.com/agaffney/gwebirc/irc"
	"net/http"
	"strings"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request, []string){
	"connections": handle_connections,
}

func api_handler(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimPrefix(r.URL.Path, "/api/")
	url_parts := strings.Split(url, "/")
	if fn, ok := handlers[url_parts[0]]; ok {
		fn(w, r, url_parts)
	} else {
		// Return a 404
		http.NotFound(w, r)
	}
}

func handle_connections(w http.ResponseWriter, r *http.Request, params []string) {
	j := json.NewEncoder(w)
	switch len(params) {
	case 1:
		// List connections
		j.Encode(irc_conns)
	case 2:
		// List specific connection by name
		found_conn := false
		for _, conn := range irc_conns {
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
		for _, c := range irc_conns {
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
	// Send explicit 200 if we end up here
	w.WriteHeader(http.StatusOK)
}
