package irc

import (
	"strings"
)

type Command struct {
	Source      string
	Source_nick string
	Command     string
	Args        []string
}

func parse_command(input string) *Command {
	cmd := &Command{}
	input = strings.TrimRight(input, " \r\n")
	pieces := strings.Split(input, " ")
	if strings.HasPrefix(pieces[0], ":") {
		// Command includes a source
		cmd.Source = pieces[0][1:]
		pieces = pieces[1:]
		source_parts := strings.Split(cmd.Source, "!")
		if len(source_parts) > 1 {
			cmd.Source_nick = source_parts[0]
		}
	}
	// Grab the command
	cmd.Command = pieces[0]
	pieces = pieces[1:]
	// Grab the rest of the args
	for idx, arg := range pieces {
		if strings.HasPrefix(arg, ":") {
			// Strip off leading colon
			pieces[idx] = pieces[idx][1:]
			cmd.Args = append(cmd.Args, strings.Join(pieces[idx:], " "))
			break
		} else {
			cmd.Args = append(cmd.Args, arg)
		}
	}
	return cmd
}
