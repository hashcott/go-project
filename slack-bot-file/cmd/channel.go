package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"log"
)

var cmdChannel = &cobra.Command{
	Use:   "channel [command]",
	Short: "Management channel id.",
	Long:  `You can create, delete channel id to list.`,
}

var cmdAddChannel = &cobra.Command{
	Use:   "add [LIST_CHANNEL_ID]",
	Short: "Add channel id to list.",
	Long:  "You can add channel id to list channel config.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currentChannelIds := viper.GetStringSlice("channel_ids")

		idx := slices.IndexFunc(currentChannelIds, func(s string) bool {
			return s == args[0]
		})

		if idx != -1 {
			log.Fatal("Error occurred: Channel ID exists !")
		}

		channelIds := make([]string, 0)
		channelIds = append(channelIds, currentChannelIds...)
		channelIds = append(channelIds, args...)
		viper.Set("channel_ids", channelIds)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal("Error occurred: ", err)
		}
		fmt.Println("Add channel successfully !")
	},
}

var cmdRemoveChannel = &cobra.Command{
	Use:   "remove [CHANNEL_ID]",
	Short: "Remove channel id from list.",
	Long:  "You can remove channel id from list channel config. You can remove only one channel id per command.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currentChannelIds := viper.GetStringSlice("channel_ids")
		idx := slices.IndexFunc(currentChannelIds, func(s string) bool {
			return s == args[0]
		})
		currentChannelIds = append(currentChannelIds[:idx], currentChannelIds[idx+1:]...)
		viper.Set("channel_ids", currentChannelIds)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal("Error occurred: ", err)
		}
		fmt.Println("Remove channel successfully !")
	},
}

func init() {
	cmdChannel.AddCommand(cmdAddChannel)
	cmdChannel.AddCommand(cmdRemoveChannel)
	cmdRoot.AddCommand(cmdChannel)
}
