package config

import (
	"flag"
)

func (c *Config) Parse_command_line() {
	config_file := flag.String("config", default_config_file(), "config file")
	flag.Parse()
	c.config_file = *config_file
}
