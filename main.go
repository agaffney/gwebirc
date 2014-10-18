package main

import (
	"github.com/agaffney/gwebirc/config"
	"github.com/agaffney/gwebirc/irc"
)

func main() {
	conf := &config.Config{}
	conf.Parse_command_line()
	irc := &irc.Connection{Host: "irc.freenode.net", Port: 6667, Tls: false}
	irc.Connect()
}
