package gitter

import "time"

type Message struct {
	Id          string      `json:"id"`
	Text        string      `json:"text"`
	Html        string      `json:"html"`
	Sent        time.Time   `json:"sent"`
	EditedAt    time.Time   `json:"editatedAt"`
	Sender      User        `json:"fromUser"`
	Unread      bool        `json:"unread"`
	ReaderCount int         `json:"readBy"`
	Urls        []string    `json:"urls"`
	Mentions    []string    `json:"mentions"`
	Issues      []string    `json:"issues"`
	Meta        interface{} `json:"meta"`
}
