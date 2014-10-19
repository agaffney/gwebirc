package irc

import (
	"fmt"
)

func (c *Connection) read_from_server() {
	for {
		str, err := c.readbuf.ReadString('\n')
		if len(str) > 0 {
			fmt.Print(str)
			c.parse_command(str)
		}
		if err != nil {
			break
		}
	}

}
