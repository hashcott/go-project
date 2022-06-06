package utils

import (
	"fmt"
	"github.com/spf13/viper"
	api "slack-bot-file/api/slack"
)

type Config struct {
	SlackBotToken string   `mapstructure:"slack_bot_token"`
	ChannelIds    []string `mapstructure:"channel_ids"`
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

	if err == nil {
		api.InitSlackApi(config.SlackBotToken)
	}

	return
}
