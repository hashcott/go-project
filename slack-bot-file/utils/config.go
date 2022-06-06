package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	SlackBotToken string   `yaml:"slack_bot_token"`
	ChannelIds    []string `yaml:"channel_ids"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	// Config name for file env
	viper.SetConfigName("app")

	// Config type file env .yaml
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&config)
	return
}
