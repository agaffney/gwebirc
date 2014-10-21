package irc

import (
	"fmt"
	"strconv"
	"strings"
)

var handlers = map[string]func(*Connection, *Command){
	"PING": handle_ping,
	"MODE": handle_mode,
	"JOIN": handle_join,
	"324":  handle_channel_info,
	"329":  handle_channel_info,
	"353":  handle_channel_info,
	"366":  handle_channel_info,
}

func (c *Connection) handle_server_command(cmd *Command) {
	if fn, ok := handlers[cmd.command]; ok {
		fn(c, cmd)
	} else {
		// Just print the unhandled commands for now
		fmt.Println(cmd)
	}
}

func handle_ping(c *Connection, cmd *Command) {
	// The server prefers that we respond to the PING command
	c.Send(fmt.Sprintf("PONG :%s\r\n", cmd.args[0]))
}

func handle_join(c *Connection, cmd *Command) {
	// Create structure for newly joined channel
	c.channels[cmd.args[0]] = &Channel{Name: cmd.args[0]}
	// We want to know the current mode of the channel
	c.Send(fmt.Sprintf("MODE :%s\r\n", cmd.args[0]))
}

func handle_channel_info(c *Connection, cmd *Command) {
	switch cmd.command {
	case "324":
		// Channel mode
		ch := c.channels[cmd.args[1]]
		ch.Set_mode(cmd.args[2])
	case "329":
		// Channel create time
		// :hitchcock.freenode.net 329 gwebirc #gwebirc 1413696501
		ch := c.channels[cmd.args[1]]
		ch.Timestamp, _ = strconv.ParseUint(cmd.args[2], 10, 64)
	case "353":
		// Channel name list
		// :asimov.freenode.net 353 gwebirc @ #gwebirc :gwebirc @agaffney
		ch := c.channels[cmd.args[2]]
		ch.Add_names(strings.Split(cmd.args[3], " "))
	case "366":
		// End of channel name list
		// :asimov.freenode.net 366 gwebirc #gwebirc :End of /NAMES list.
		ch := c.channels[cmd.args[1]]
		ch.Finalize_names()
	}
}

func handle_mode(c *Connection, cmd *Command) {
	if strings.HasPrefix(cmd.args[0], "#") {
		// Channel
		ch := c.channels[cmd.args[0]]
		if len(cmd.args) == 3 {
			// Mode on user in channel

		} else if len(cmd.args) == 2 {
			// Channel mode
			ch.Set_mode(cmd.args[1])
		}
	} else {
		// User
		c.user.Set_mode(cmd.args[1])
	}
}
