package irc

import (
	"fmt"
	"strings"
)

var handlers = map[string]func(*Connection, *Command){
	"PING": handle_ping,
	"MODE": handle_mode,
	"JOIN": handle_join,
}

func (c *Connection) handle_server_command(cmd *Command) {
	if fn, ok := handlers[cmd.command]; ok {
		fn(c, cmd)
	} else {
		fmt.Printf("No handler for command '%s'\n", cmd.command)
	}
}

func handle_ping(c *Connection, cmd *Command) {
	// The server prefers that we respond to the PING command
	c.Send(fmt.Sprintf("PONG :%s\r\n", cmd.args[0]))
}

func handle_join(c *Connection, cmd *Command) {
	// Create structure for newly joined channel
	c.channels[cmd.args[0]] = Channel{name: cmd.args[0]}
}

func handle_mode(c *Connection, cmd *Command) {
	if strings.HasPrefix(cmd.args[0], "#") {
		// Channel
		ch := c.channels[cmd.args[0]]
		ch.Set_mode(cmd.args[1])
	} else {
		// User
		c.mode = cmd.args[1]
		fmt.Printf("User mode is now %s\n", c.mode)
	}
}
