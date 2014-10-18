package main

import (
	"github.com/agaffney/gwebirc/irc"
)

func main() {
	irc := &irc.Connection{Host: "irc.freenode.net", Port: 6667, Tls: false}
	irc.Connect()
}
