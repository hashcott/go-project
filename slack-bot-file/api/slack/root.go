package api

import (
	"github.com/slack-go/slack"
)

var api *slack.Client

func InitSlackApi(token string) {
	api = slack.New(token)
}
