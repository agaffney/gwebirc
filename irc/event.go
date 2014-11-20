package irc

import (
	"fmt"
	"strings"
)

type Event struct {
	Source      string   `json:"source"`
	Source_nick string   `json:"source_nick"`
	Code        string   `json:"code"`
	Args        []string `json:"args"`
	Msg         string   `json:"msg"`
	Channel     string   `json:"channel"`
	Raw         string   `json:"-"`
}

func (e *Event) parse() {
	line := strings.TrimRight(e.Raw, " \r\n")
	// The event includes a source if it begins with a colon
	if line[0] == ':' {
		idx := strings.Index(line, " ")
		e.Source = line[1:idx]
		line = line[idx+1:]
		source_parts := strings.SplitN(e.Source, "!", 2)
		if len(source_parts) > 1 {
			e.Source_nick = source_parts[0]
		}
	}
	// Check for a message at the end prefixed with a colon
	idx := strings.Index(line, ":")
	if idx >= 0 {
		e.Msg = line[idx+1:]
		// Remove from the space (we assume) before the colon to the end of the string
		line = line[0 : idx-1]
	}
	// Split the remainder on space
	pieces := strings.Split(line, " ")
	// Grab the event code
	e.Code = pieces[0]
	// Everything else goes to args
	e.Args = pieces[1:]
	// Do extra parsing for CTCP
	if e.Code == "PRIVMSG" && e.Msg[0] == '\x01' {
		// Strip off surrounding \x01 from message
		e.Msg = e.Msg[1 : len(e.Msg)-1]
		ctcp_args := strings.SplitN(e.Msg, " ", 2)
		e.Code = fmt.Sprintf("CTCP_%s", ctcp_args[0])
		if len(ctcp_args) > 1 {
			e.Msg = ctcp_args[1]
		}
	}
}
