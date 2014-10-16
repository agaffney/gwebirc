package main

import (
	"fmt"
	"github.com/agaffney/gwebirc/irc"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("Irc_foo() -> %s\n", irc.Irc_foo())
	// just testing the highlighting
	fmt.Println("foo again")
}
