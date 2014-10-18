package config

import (
	"fmt"
	"os/user"
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
	usr, _ := user.Current()
	default_config_file := fmt.Sprintf("%s/.gwebirc", usr.HomeDir)
	return default_config_file
}
