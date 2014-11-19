package core

import (
	"fmt"
	"github.com/agaffney/gwebirc/config"
	"github.com/agaffney/gwebirc/irc"
	"github.com/agaffney/gwebirc/web"
	"os"
)

var conf *config.Config
var irc_conns = []*irc.Connection{}

func Start() {
	conf = &config.Config{}
	conf.Parse_command_line()
	err := conf.Parse_config_file()
	if err != nil {
		fmt.Printf("failed to parse config file: %s\n", err)
		os.Exit(1)
	}
	//conf.Write_config_file()
	// Setup our HTTP server
	w := &web.Web{Conf: conf, Conns: irc_conns}
	w.Start()
	// Start our IRC connections
	irc_start()
	// Block indefinitely
	select {}
}
