package config

import (
	"fmt"
	"github.com/agaffney/gwebirc/util"
)

type Config struct {
	config_file string
	Servers     []IRCServer   `json:"servers"`
	Listeners   []WebListener `json:"listeners"`
}

type IRCServer struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Use_tls      bool   `json:"use_tls"`
	Tls_Verify   bool   `json:"tls_verify"`
	Auto_connect bool   `json:"auto_connect"`
}

type WebListener struct {
	Port    int    `json:"port"`
	Use_tls bool   `json:"use_tls"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
	Cacert  string `json:"cacert"`
}

func default_config_file() string {
	default_config_file := fmt.Sprintf("%s/.gwebirc", util.Get_homedir())
	return default_config_file
}
