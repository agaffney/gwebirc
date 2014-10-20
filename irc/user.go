package irc

import (
	"fmt"
)

type User struct {
	conn      *Connection
	name      string
	real_name string
	bitmask   uint
	mode      string
	nick      string
}

func (u *User) Set_mode(mode string) {
	u.mode = merge_modes(u.mode, mode)
	fmt.Printf("User mode is now %s\n", u.mode)
}
