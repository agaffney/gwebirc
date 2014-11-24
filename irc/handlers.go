package irc

import (
	"fmt"
	"strconv"
	"strings"
)

type HandlerFunc func(*Connection, *Event)

func (c *Connection) setup_handlers() {
	c.Add_handler("PING", func(c *Connection, e *Event) {
		c.Send("PONG :" + e.Msg)
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
	for _, x := range e.Args {
		if x[0] == '#' {
			// This looks like it was targetted at a channel
			e.Channel = x
			break
		}
	}
	if e.Channel == "" {
		if e.Source_nick == "" {
			// This is a server message
			e.Channel = c.Channels[0].Name
		} else {
			// This looks like it was targetted at us directly
			ch := c.Get_channel(e.Source_nick)
			if ch == nil {
				c.Channels = append(c.Channels, &Channel{Type: CH_TYPE_USER, Name: e.Source_nick, conn: c})
			}
			e.Channel = e.Source_nick
		}
	}
	e.Connection = c.Name
	// Call the event callback, if defined
	if c.manager.callback != nil {
		c.manager.callback(e, c)
	}
}

func handle_join(c *Connection, cmd *Event) {
	if cmd.Source_nick == c.Nick {
		// Create structure for newly joined channel
		c.Channels = append(c.Channels, &Channel{Type: CH_TYPE_CHANNEL, Name: cmd.Args[0], conn: c})
		// We want to know the current mode of the channel
		c.Send(fmt.Sprintf("MODE :%s\r\n", cmd.Args[0]))
	} else {
		// Somebody has joined a channel that we're in
		// Refresh the list of names
		ch := c.Get_channel(cmd.Args[0])
		ch.Refresh_names()
	}
}

func handle_part(c *Connection, cmd *Event) {
	if cmd.Source_nick != c.Nick {
		// Somebody left a channel, so refresh the names
		ch := c.Get_channel(cmd.Args[0])
		ch.Refresh_names()
	} else {
		// We left a channel
		for idx, ch := range c.Channels {
			if ch.Name == cmd.Args[0] {
				c.Channels = append(c.Channels[:idx], c.Channels[idx+1:]...)
				break
			}
		}
	}
}

func handle_channel_info(c *Connection, cmd *Event) {
	switch cmd.Code {
	case "324":
		// Channel mode
		ch := c.Get_channel(cmd.Args[1])
		ch.Set_mode(cmd.Args[2])
	case "329":
		// Channel create time
		// :hitchcock.freenode.net 329 gwebirc #gwebirc 1413696501
		ch := c.Get_channel(cmd.Args[1])
		ch.Timestamp, _ = strconv.ParseUint(cmd.Args[2], 10, 64)
	case "353":
		// Channel name list
		// :asimov.freenode.net 353 gwebirc @ #gwebirc :gwebirc @agaffney
		ch := c.Get_channel(cmd.Args[2])
		ch.Add_names(strings.Split(cmd.Msg, " "))
	case "366":
		// End of channel name list
		// :asimov.freenode.net 366 gwebirc #gwebirc :End of /NAMES list.
		ch := c.Get_channel(cmd.Args[1])
		ch.Finalize_names()
	}
}

func handle_mode(c *Connection, cmd *Event) {
	if strings.HasPrefix(cmd.Args[0], "#") {
		// Channel
		ch := c.Get_channel(cmd.Args[0])
		if len(cmd.Args) == 3 {
			// Mode on user in channel
			// Refresh the names list in case someone got voice/op
			ch.Refresh_names()
		} else if len(cmd.Args) == 2 {
			// Channel mode
			ch.Set_mode(cmd.Args[1])
		}
	} else {
		// User
		c.user.Set_mode(cmd.Msg)
	}
}
