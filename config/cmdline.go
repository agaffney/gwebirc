package config

import (
	"flag"
	"fmt"
)

func (c *Config) Parse_command_line() {
	config_file := flag.String("config", "./gwebirc.json", "config file")
	flag.Parse()
	c.config_file = *config_file
	fmt.Printf("config_file = %s\n", c.config_file)
}
