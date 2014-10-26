package config

import (
	"flag"
)

func (c *Config) Parse_command_line() {
	config_file := flag.String("config", "", "config file, defaults to ~/.gwebirc")
	flag.Parse()
	if *config_file != "" {
		c.config_file = *config_file
	}
}
