package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	api "slack-bot-file/api/slack"
)

var cmdUpload = &cobra.Command{
	Use:   "upload [FILE_PATHS]",
	Short: "Upload multiple file to all channel Slack.",
	Long:  `Upload multiple file to all channel Slack. You can upload with multiple file to multiple channel.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := api.Upload(args)
		if err != nil {
			log.Fatal("Error occurred:", err)
		}
		fmt.Println("DONE !")
	},
}

func init() {
	cmdRoot.AddCommand(cmdUpload)
}
