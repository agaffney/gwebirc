package main

import (
	"fmt"
	"github.com/agaffney/gwebirc/core"
	"github.com/agaffney/gwebirc/irc"
	"github.com/agaffney/gwebirc/web"
	"os"
)

func main() {
	conf := &core.Config{}
	conf.Parse_command_line()
	err := conf.Parse_config_file()
	if err != nil {
		fmt.Printf("failed to parse config file: %s\n", err)
		os.Exit(1)
	}
	//conf.Write_config_file()
	// Create our IRC manager
	i := &irc.IrcManager{Conf: conf}
	// Create the Web manager
	w := &web.WebManager{Conf: conf, Irc: i}
	// Start both the web server and IRC connections
	w.Start()
	// Block indefinitely
	select {}
}
