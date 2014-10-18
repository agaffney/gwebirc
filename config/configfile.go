package config

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(c)
	return nil
}
