package irc

import (
	"fmt"
	"github.com/agaffney/gwebirc/core"
	"github.com/agaffney/gwebirc/types"
)

type IrcManager struct {
	Conf   *core.Config
	Conns  []*Connection
	Events chan *types.IrcEvent
}

func (im *IrcManager) Start() {
	// Initialize some things
	im.Events = make(chan *types.IrcEvent, 100)
	// Start the configured IRC connections
	for _, conn := range im.Conf.Connections {
		i := &Connection{Name: conn.Name, Host: conn.Host, Port: conn.Port, Tls: conn.Tls, manager: im}
		im.Conns = append(im.Conns, i)
		i.Init()
		// Add our handlers
		i.Add_handler("PRIVMSG", handle_msg)
		i.Add_handler("NOTICE", handle_msg)
		i.Add_handler("366", handle_366) // 366 means end of NAMES list
		go i.Start()
	}
}

func handle_366(c *Connection, cmd *types.IrcEvent) {
	channel := cmd.Args[1]
	fmt.Printf("Names list for %s:\n", channel)
	for _, name := range c.Get_channel(channel).Names {
		fmt.Println(name)
	}
}

func handle_msg(c *Connection, cmd *types.IrcEvent) {

}
