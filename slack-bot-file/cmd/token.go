package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var cmdToken = &cobra.Command{
	Use:   "token [TOKEN_ACCESS]",
	Short: "Set config token access.",
	Long:  `Set OAuth Tokens to access bot Slack.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("SLACK_BOT_TOKEN", args)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal("Error occurred: ", err)
		}
		fmt.Println("Add token successfully !")
	},
}

func init() {
	cmdRoot.AddCommand(cmdToken)
}
