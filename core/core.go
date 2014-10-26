package core

import (
	"fmt"
	"github.com/agaffney/gwebirc/config"
	"github.com/agaffney/gwebirc/irc"
	"net/http"
	"os"
)

var conf *config.Config
var conn []*irc.Connection

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
	if conf.Http.Enable_webui {
		http.Handle("/webui/", http.StripPrefix("/webui/", http.FileServer(http.Dir("./webui"))))
	}
	http.HandleFunc("/api/", api_handler)
	go http.ListenAndServe(fmt.Sprintf(":%d", conf.Http.Port), nil)
	// Start our IRC connections
	conn = make([]*irc.Connection, 1)
	for _, server := range conf.Servers {
		irc := &irc.Connection{Host: server.Host, Port: server.Port, Tls: server.Use_tls}
		conn = append(conn, irc)
		irc.Init()
		irc.Add_handler("366", handle_366)
		go irc.Start()
	}
	// Block indefinitely
	select {}
}

func handle_366(c *irc.Connection, cmd *irc.Event) {
	channel := cmd.Args[1]
	fmt.Printf("Names list for %s:\n", channel)
	for _, name := range c.Channels[channel].Names {
		fmt.Println(name)
	}
}
