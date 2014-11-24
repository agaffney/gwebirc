package irc

import (
	"fmt"
)

const (
	CH_TYPE_SERVER  string = "server"
	CH_TYPE_CHANNEL string = "channel"
	CH_TYPE_USER    string = "user"
)

type Channel struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Timestamp uint64   `json:"timestamp,omitempty"`
	Mode      string   `json:"mode,omitempty"`
	Names     []string `json:"names,omitempty"`
	conn      *Connection
	new_names []string
}

func (ch *Channel) Set_mode(mode string) {
	ch.Mode = merge_modes(ch.Mode, mode)
}

func (ch *Channel) Refresh_names() {
	ch.conn.Send(fmt.Sprintf("NAMES %s\r\n", ch.Name))
}

func (ch *Channel) Add_names(names []string) {
	// Append to new_names until we call finalize below
	ch.new_names = append(ch.new_names, names...)
}

func (ch *Channel) Finalize_names() {
	ch.Names = ch.new_names
	ch.new_names = make([]string, 0)
}

func (ch *Channel) Part() {
	ch.conn.Part(ch.Name)
}
