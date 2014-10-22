package irc

import (
	"fmt"
	"strconv"
	"strings"
)

type HandlerFunc func(*Connection, *Command)

var internal_handlers = map[string][]HandlerFunc{
	"PING": {handle_ping},
	"MODE": {handle_mode},
	"JOIN": {handle_join},
	"PART": {handle_part},
	"324":  {handle_channel_info},
	"329":  {handle_channel_info},
	"353":  {handle_channel_info},
	"366":  {handle_channel_info},
}

func (c *Connection) Add_handler(cmd string, fn HandlerFunc) {
	if c.handlers == nil {
		c.handlers = make(map[string][]HandlerFunc)
	}
	if _, ok := c.handlers[cmd]; !ok {
		// Create empty array for key
		c.handlers[cmd] = []HandlerFunc{}
	}
	c.handlers[cmd] = append(c.handlers[cmd], fn)
}

func (c *Connection) handle_server_command(cmd *Command) {
	fmt.Println(cmd)
	// Call internal handler first
	if handlers, ok := internal_handlers[cmd.Command]; ok {
		for _, fn := range handlers {
			fn(c, cmd)
		}
	}
	if handlers, ok := c.handlers[cmd.Command]; ok {
		for _, fn := range handlers {
			fn(c, cmd)
		}
	}
}

func handle_ping(c *Connection, cmd *Command) {
	// The server prefers that we respond to the PING command
	c.Send(fmt.Sprintf("PONG :%s\r\n", cmd.Args[0]))
}

func handle_join(c *Connection, cmd *Command) {
	if cmd.Source_nick == c.Nick {
		// Create structure for newly joined channel
		c.Channels[cmd.Args[0]] = &Channel{Name: cmd.Args[0], conn: c}
		// We want to know the current mode of the channel
		c.Send(fmt.Sprintf("MODE :%s\r\n", cmd.Args[0]))
	} else {
		// Somebody has joined a channel that we're in
		// Refresh the list of names
		ch := c.Channels[cmd.Args[0]]
		ch.Refresh_names()
	}
}

func handle_part(c *Connection, cmd *Command) {
	if cmd.Source_nick != c.Nick {
		// Somebody left a channel, so refresh the names
		ch := c.Channels[cmd.Args[0]]
		ch.Refresh_names()
	} else {
		// We left a channel
		delete(c.Channels, cmd.Args[0])
	}
}

func handle_channel_info(c *Connection, cmd *Command) {
	switch cmd.Command {
	case "324":
		// Channel mode
		ch := c.Channels[cmd.Args[1]]
		ch.Set_mode(cmd.Args[2])
	case "329":
		// Channel create time
		// :hitchcock.freenode.net 329 gwebirc #gwebirc 1413696501
		ch := c.Channels[cmd.Args[1]]
		ch.Timestamp, _ = strconv.ParseUint(cmd.Args[2], 10, 64)
	case "353":
		// Channel name list
		// :asimov.freenode.net 353 gwebirc @ #gwebirc :gwebirc @agaffney
		ch := c.Channels[cmd.Args[2]]
		ch.Add_names(strings.Split(cmd.Args[3], " "))
	case "366":
		// End of channel name list
		// :asimov.freenode.net 366 gwebirc #gwebirc :End of /NAMES list.
		ch := c.Channels[cmd.Args[1]]
		ch.Finalize_names()
	}
}

func handle_mode(c *Connection, cmd *Command) {
	if strings.HasPrefix(cmd.Args[0], "#") {
		// Channel
		ch := c.Channels[cmd.Args[0]]
		if len(cmd.Args) == 3 {
			// Mode on user in channel
			ch.Refresh_names()
		} else if len(cmd.Args) == 2 {
			// Channel mode
			ch.Set_mode(cmd.Args[1])
		}
	} else {
		// User
		c.user.Set_mode(cmd.Args[1])
	}
}
