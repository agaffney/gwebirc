package irc

import (
	"fmt"
	"strings"
)

type Event struct {
	Raw         string
	Source      string
	Source_nick string
	Code        string
	Args        []string
}

func (e *Event) parse() {
	e.Raw = strings.TrimRight(e.Raw, " \r\n")
	pieces := strings.Split(e.Raw, " ")
	if strings.HasPrefix(pieces[0], ":") {
		// Event includes a source
		e.Source = pieces[0][1:]
		pieces = pieces[1:]
		source_parts := strings.Split(e.Source, "!")
		if len(source_parts) > 1 {
			e.Source_nick = source_parts[0]
		}
	}
	// Grab the command
	e.Code = pieces[0]
	pieces = pieces[1:]
	// Grab the rest of the args
	for idx, arg := range pieces {
		if strings.HasPrefix(arg, ":") {
			// Strip off leading colon
			pieces[idx] = pieces[idx][1:]
			e.Args = append(e.Args, strings.Join(pieces[idx:], " "))
			break
		} else {
			e.Args = append(e.Args, arg)
		}
	}
	// Do extra parsing for CTCP
	if e.Code == "PRIVMSG" && e.Args[1][0] == '\x01' {
		// Strip off surrounding \x01 from message
		e.Args[1] = e.Args[1][1 : len(e.Args[1])-1]
		ctcp_args := strings.SplitN(e.Args[1], " ", 2)
		e.Code = fmt.Sprintf("CTCP_%s", ctcp_args[0])
		if len(ctcp_args) > 1 {
			e.Args[1] = ctcp_args[1]
		} else {
			e.Args = e.Args[0:1]
		}
	}
}
