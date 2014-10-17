package main

import (
	"fmt"
	"github.com/agaffney/gwebirc/irc"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("Irc_foo() -> %s\n", irc.Irc_foo())
	irc := &irc.Connection{Host: "irc.freenode.net", Port: 6667, Tls: false}
	irc.Connect()
}
