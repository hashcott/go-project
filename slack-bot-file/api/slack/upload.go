package api

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func Upload(filePaths []string) error {
	for _, s := range filePaths {
		params := slack.FileUploadParameters{
			Channels: viper.GetStringSlice("channel_ids"),
			File:     s,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s, Url: %s ==> Uploaded successfully. \n", file.Name, file.URL)
	}
	return nil
}
