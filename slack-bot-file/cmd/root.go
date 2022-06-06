package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"slack-bot-file/utils"
)

var cmdRoot = &cobra.Command{
	Use:   "slack-bot",
	Short: "Upload file to channel Slack.",
	Long:  `Upload file to channel Slack. You can upload with multiple file to multiple channel.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Slack-File-Bot")
	},
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	_, err := utils.LoadConfig("config")
	if err != nil {
		log.Fatal("Can't load config: ", err)
	}
}
