package irc

import (
	"fmt"
)

const (
	CH_TYPE_SERVER uint8 = iota
	CH_TYPE_CHANNEL
	CH_TYPE_USER
)

// String representations of the above types
var CH_TYPES []string = []string{"server", "channel", "user"}

type Channel struct {
	Type      uint8    `json:"type"`
	Name      string   `json:"name"`
	Timestamp uint64   `json:"timestamp"`
	Mode      string   `json:"mode"`
	Names     []string `json:"names"`
	conn      *Connection
	new_names []string
	events    []*Event
}

func (ch *Channel) Add_event(e *Event) {
	if ch.events == nil {
		ch.events = make([]*Event, 0)
	}
	ch.events = append(ch.events, e)
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
