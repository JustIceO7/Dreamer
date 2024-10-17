package config

import (
	"os"

	"github.com/spf13/viper"
)

func initDefaults() {
	viper.SetDefault("discord.token", os.Getenv("discord_token"))
	viper.SetDefault("discord.app.id", os.Getenv("discord_app_id"))
	viper.SetDefault("master", os.Getenv("master"))

	viper.SetDefault("port", "8080")
	viper.SetDefault("sd.url", "http://localhost:7860/")
	viper.SetDefault("queue", "")
	viper.SetDefault("checkpoints", "")
	viper.SetDefault("loras", "")

	viper.SetDefault("verified", map[string]bool{})

}
