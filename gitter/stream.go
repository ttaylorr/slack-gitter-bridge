package gitter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stream struct {
	room     string
	auth     string
	messages chan *Message
}

func OpenStream(auth, room string) (*Stream, error) {
	s := &Stream{
		room:     room,
		auth:     auth,
		messages: make(chan *Message),
	}

	go s.parseMessages()

	return s, nil
}

func (s *Stream) Messages() <-chan *Message {
	return s.messages
}

func (s *Stream) parseMessages() error {
	req, err := s.request()
	if err != nil {
		return err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		var next = &Message{}
		if err := decoder.Decode(next); err != nil {
			return err
		}

		s.messages <- next
	}

	return nil
}

func (s *Stream) request() (*http.Request, error) {
	url := fmt.Sprintf("https://stream.gitter.im/v1/rooms/%s/chatMessages", s.room)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", s.auth))

	return req, nil
}
