package irc

import (
	"fmt"
)

type Channel struct {
	name      string
	mode      string
	names     []string
	new_names []string
}

func (c *Connection) Join(channel string) {
	c.Send(fmt.Sprintf("JOIN %s\r\n", channel))
}

func (ch *Channel) Set_mode(mode string) {
	ch.mode = merge_modes(ch.mode, mode)
}

func (ch *Channel) Add_names(names []string) {
	// Append to new_names until we call finalize below
	ch.new_names = append(ch.new_names, names...)
}

func (ch *Channel) Finalize_names() {
	ch.names = ch.new_names
	ch.new_names = make([]string, 5)
}
