package main

import (
	"github.com/ttaylorr/go-config/config"
	"github.com/ttaylorr/slack-gitter-bridge/gitter"
	"github.com/ttaylorr/slack/api"
)

func main() {
	config, err := config.New("./config")
	if err != nil {
		panic(err)
	}

	slack := Slack(config)
	slackChannel, _ := config.String("slack.room")

	gitter, err := Gitter(config)
	if err != nil {
		panic(err)
	}

	for {
		PostToSlack(slack, slackChannel, <-gitter.Messages())
	}
}

func Slack(config *config.Configuration) *api.Slack {
	slackApi, _ := config.String("slack.auth")
	slack := api.New(slackApi)

	return slack
}

func Gitter(config *config.Configuration) (*gitter.Stream, error) {
	auth, _ := config.String("gitter.auth")
	room, _ := config.String("gitter.room")

	return gitter.OpenStream(
		auth,
		room,
	)
}

func PostToSlack(slack *api.Slack, channel string, msg *gitter.Message) {
	slack.Request(&api.RequestParam{
		Method: api.ChatPostMessageMethod,
		Parameters: map[string]string{
			"channel":  channel,
			"text":     msg.Text,
			"username": msg.Sender.Username,
			"as_user":  "true",
			"icon_url": msg.Sender.AvatarUrlSmall,
		},
	})

}
