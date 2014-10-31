package core

import (
	"fmt"
	"github.com/agaffney/gwebirc/irc"
)

func irc_start() {
	// Start the configured IRC connections
	for _, conn := range conf.Connections {
		irc := &irc.Connection{Name: conn.Name, Host: conn.Host, Port: conn.Port, Tls: conn.Tls}
		irc_conns = append(irc_conns, irc)
		irc.Init()
		// Add our handlers
		irc.Add_handler("PRIVMSG", handle_msg)
		irc.Add_handler("NOTICE", handle_msg)
		irc.Add_handler("366", handle_366) // 366 means end of NAMES list
		go irc.Start()
	}
}

func handle_366(c *irc.Connection, cmd *irc.Event) {
	channel := cmd.Args[1]
	fmt.Printf("Names list for %s:\n", channel)
	for _, name := range c.Get_channel(channel).Names {
		fmt.Println(name)
	}
}

func handle_msg(c *irc.Connection, cmd *irc.Event) {

}
