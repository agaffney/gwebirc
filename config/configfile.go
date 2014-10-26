package config

import (
	"encoding/json"
	"fmt"
	"github.com/agaffney/gwebirc/util"
	"io/ioutil"
)

func (c *Config) init() {
	c.Servers = make([]IRCServer, 0)
	// Set some sane defaults for an empty config
	c.config_file = fmt.Sprintf("%s/.gwebirc", util.Get_homedir())
	c.Http.Port = 9002
}

func (c *Config) Parse_config_file() error {
	c.init()
	content, err := ioutil.ReadFile(c.config_file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Write_config_file() error {
	// Indent using two spaces
	content, err := json.MarshalIndent(*c, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.config_file, content, 0600)
	if err != nil {
		return err
	}
	return nil
}
