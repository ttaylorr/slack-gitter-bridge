package main

import (
	"github.com/ttaylorr/go-config/config"
	"github.com/ttaylorr/slack-gitter-bridge/bridge"
	"github.com/ttaylorr/slack-gitter-bridge/gitter"
	"github.com/ttaylorr/slack/api"
)

func main() {
	config, err := config.New("./config")
	if err != nil {
		panic(err)
	}

	slack := Slack(config)
	gitter, _ := Gitter(config)
	slackChannel, _ := config.String("slack.room")

	bridge := bridge.New(slack, slackChannel, gitter)

	bridge.Open()
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
