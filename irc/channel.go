package irc

type Channel struct {
	Name      string
	Timestamp uint64
	Mode      string
	Names     []string
	new_names []string
}

func (ch *Channel) Set_mode(mode string) {
	ch.Mode = merge_modes(ch.Mode, mode)
}

func (ch *Channel) Add_names(names []string) {
	// Append to new_names until we call finalize below
	ch.new_names = append(ch.new_names, names...)
}

func (ch *Channel) Finalize_names() {
	ch.Names = ch.new_names
	ch.new_names = make([]string, 5)
}
