package irc

import (
	"fmt"
	"strings"
)

type Command struct {
	source  string
	command string
	args    []string
}

func (c *Connection) parse_command(input string) {
	cmd := Command{}
	input = strings.TrimRight(input, "\r\n")
	pieces := strings.Split(input, " ")
	if strings.HasPrefix(pieces[0], ":") {
		// Command includes a source
		cmd.source = pieces[0][1:]
		pieces = pieces[1:]
	}
	// Grab the command
	cmd.command = pieces[0]
	pieces = pieces[1:]
	// Grab the rest of the args
	for idx, arg := range pieces {
		if strings.HasPrefix(arg, ":") {
			// Strip off leading colon
			pieces[idx] = pieces[idx][1:]
			cmd.args = append(cmd.args, strings.Join(pieces[idx:], " "))
			break
		} else {
			cmd.args = append(cmd.args, arg)
		}
	}
	fmt.Println(cmd)
}
