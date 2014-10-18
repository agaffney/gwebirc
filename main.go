package main

import (
	"fmt"
	"github.com/agaffney/gwebirc/config"
	"github.com/agaffney/gwebirc/irc"
	"os"
)

func main() {
	conf := &config.Config{}
	conf.Parse_command_line()
	err := conf.Parse_config_file()
	if err != nil {
		fmt.Printf("failed to parse config file: %s\n", err)
		os.Exit(1)
	}
	irc := &irc.Connection{Host: "irc.freenode.net", Port: 6667, Tls: false}
	irc.Connect()
}
