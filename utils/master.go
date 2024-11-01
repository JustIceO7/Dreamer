package utils

import "github.com/spf13/viper"

// IsMaster checks if given discord uid is master returns bool
func IsMaster(uid string) bool {
	return uid == viper.GetString("master")
}
