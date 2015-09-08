package bridge

import (
	"sync"

	"github.com/ttaylorr/slack-gitter-bridge/gitter"
	slack "github.com/ttaylorr/slack/api"
)

type Bridge struct {
	Slack     *slack.Slack
	SlackRoom string
	Gitter    *gitter.Stream
}

func New(slack *slack.Slack, slackRoom string, gitter *gitter.Stream) *Bridge {
	return &Bridge{
		Slack:     slack,
		SlackRoom: slackRoom,
		Gitter:    gitter,
	}
}

func (b *Bridge) Open() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go b.handleSlack(wg)
	go b.handleGitter(wg)

	wg.Wait()
}

func (b *Bridge) handleSlack(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// TODO(ttaylorr): implement Slack RTM in github.com/ttaylorr/slack
	}
}

func (b *Bridge) handleGitter(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		msg := <-b.Gitter.Messages()

		b.Slack.Request(&slack.RequestParam{
			Method: slack.ChatPostMessageMethod,
			Parameters: map[string]string{
				"channel":  b.SlackRoom,
				"text":     msg.Text,
				"username": msg.Sender.Username,
				"as_user":  "true",
				"icon_url": msg.Sender.AvatarUrlSmall,
			},
		})
	}
}
