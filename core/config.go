package core

import (
	"fmt"
)

type Config struct {
	config_file string
	Connections []IRCConnection `json:"connections"`
	Http        struct {
		Tls          bool   `json:"tls"`
		Cert         string `json:"cert,omitempty"`
		Key          string `json:"key,omitempty"`
		Cacert       string `json:"ca_cert,omitempty"`
		Port         uint16 `json:"port"`
		Enable_webui bool   `json:"enable_webui"`
	} `json:"http"`
}

type IRCConnection struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Tls          bool   `json:"tls"`
	Tls_Verify   bool   `json:"tls_verify"`
	Auto_connect bool   `json:"auto_connect"`
}

func default_config_file() string {
	default_config_file := fmt.Sprintf("%s/.gwebirc", Get_homedir())
	return default_config_file
}
