package irc

import (
	"fmt"
)

type Channel struct {
	name string
	mode string
}

func (c *Connection) Join(channel string) {
	c.Send(fmt.Sprintf("JOIN %s\r\n", channel))
}

func (ch *Channel) Set_mode(mode string) {
	ch.mode = mode
	fmt.Printf("Mode for channel %s is now %s\n", ch.name, ch.mode)
}
