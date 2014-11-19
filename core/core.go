package core

import (
	"fmt"
	"github.com/agaffney/gwebirc/config"
	"github.com/agaffney/gwebirc/irc"
	"github.com/agaffney/gwebirc/web"
	"os"
)

func Start() {
	conns := []*irc.Connection{}
	conf := &config.Config{}
	conf.Parse_command_line()
	err := conf.Parse_config_file()
	if err != nil {
		fmt.Printf("failed to parse config file: %s\n", err)
		os.Exit(1)
	}
	//conf.Write_config_file()
	// Setup our HTTP server
	w := &web.Web{Conf: conf, Conns: conns}
	w.Start()
	// Start our IRC connections
	i := irc.Irc{Conf: conf, Conns: conns}
	i.Start()
	// Block indefinitely
	select {}
}
