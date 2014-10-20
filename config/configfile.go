package config

import (
	"encoding/json"
	"io/ioutil"
)

func (c *Config) Parse_config_file() error {
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
	content, err := json.Marshal(*c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.config_file, content, 0600)
	if err != nil {
		return err
	}
	return nil
}
