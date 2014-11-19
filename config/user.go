package config

import (
	"os/user"
)

func get_user_object() *user.User {
	u, _ := user.Current()
	return u
}

func Get_username() string {
	return get_user_object().Username
}

func Get_homedir() string {
	return get_user_object().HomeDir
}
