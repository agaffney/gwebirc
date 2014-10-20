package irc

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
}
