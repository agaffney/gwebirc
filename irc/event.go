package irc

import (
	"fmt"
	"github.com/agaffney/gwebirc/types"
	"strings"
)

func parse_event(raw string) *types.Event {
	e := &types.Event{Type: "irc"}
	e.Init()
	e.Attribs["_raw"] = raw
	line := strings.TrimRight(raw, " \r\n")
	// The event includes a source if it begins with a colon
	if line[0] == ':' {
		idx := strings.Index(line, " ")
		e.Attribs["source"] = line[1:idx]
		line = line[idx+1:]
		source_parts := strings.SplitN(e.Attribs["source"], "!", 2)
		if len(source_parts) > 1 {
			e.Attribs["source_nick"] = source_parts[0]
		}
	}
	// Check for a message at the end prefixed with a colon
	idx := strings.Index(line, ":")
	if idx >= 0 {
		e.Attribs["msg"] = line[idx+1:]
		// Remove from the space (we assume) before the colon to the end of the string
		line = line[0 : idx-1]
	}
	// Split the remainder on space
	pieces := strings.Split(line, " ")
	// Grab the event code
	e.Attribs["code"] = pieces[0]
	// Everything else goes to args
	e.Args = pieces[1:]
	// Do extra parsing for CTCP
	if e.Attribs["code"] == "PRIVMSG" && e.Attribs["msg"][0] == '\x01' {
		// Strip off surrounding \x01 from message
		msg := e.Attribs["msg"]
		msg = msg[1 : len(msg)-1]
		ctcp_args := strings.SplitN(msg, " ", 2)
		e.Attribs["code"] = fmt.Sprintf("CTCP_%s", ctcp_args[0])
		if len(ctcp_args) > 1 {
			e.Attribs["msg"] = ctcp_args[1]
		}
	}
	return e
}
