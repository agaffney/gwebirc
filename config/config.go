package config

import (
	"fmt"
	"github.com/agaffney/gwebirc/util"
)

type Config struct {
	config_file string
	Servers     []Server
}

type Server struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Use_tls      bool   `json:"use_tls"`
	Tls_Verify   bool   `json:"tls_verify"`
	Auto_connect bool   `json:"auto_connect"`
}

func default_config_file() string {
	default_config_file := fmt.Sprintf("%s/.gwebirc", util.Get_homedir())
	return default_config_file
}
