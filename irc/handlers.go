package irc

import (
	"fmt"
	"strconv"
	"strings"
)

type HandlerFunc func(*Connection, *Event)

func (c *Connection) setup_handlers() {
	c.Add_handler("PING", func(c *Connection, e *Event) {
		c.Send("PONG :" + e.Args[0])
	})

	c.Add_handler("JOIN", handle_join)

	c.Add_handler("CTCP_VERSION", func(c *Connection, e *Event) {
		c.CtcpResponse("VERSION", e.Source_nick, "none of your business")
	})

	c.Add_handler("MODE", handle_mode)

	c.Add_handler("PART", handle_part)

	for _, foo := range []string{"324", "329", "353", "366"} {
		c.Add_handler(foo, handle_channel_info)
	}
}

func (c *Connection) Add_handler(code string, fn HandlerFunc) {
	if c.handlers == nil {
		c.handlers = make(map[string][]HandlerFunc)
	}
	if _, ok := c.handlers[code]; !ok {
		// Create empty array for key
		c.handlers[code] = []HandlerFunc{}
	}
	c.handlers[code] = append(c.handlers[code], fn)
}

func (c *Connection) handle_event(e *Event) {
	if handlers, ok := c.handlers[e.Code]; ok {
		for _, fn := range handlers {
			fn(c, e)
		}
	}
}

func handle_join(c *Connection, cmd *Event) {
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

func handle_part(c *Connection, cmd *Event) {
	if cmd.Source_nick != c.Nick {
		// Somebody left a channel, so refresh the names
		ch := c.Channels[cmd.Args[0]]
		ch.Refresh_names()
	} else {
		// We left a channel
		delete(c.Channels, cmd.Args[0])
	}
}

func handle_channel_info(c *Connection, cmd *Event) {
	switch cmd.Code {
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

func handle_mode(c *Connection, cmd *Event) {
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
