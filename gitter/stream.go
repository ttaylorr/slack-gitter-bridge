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
	closer   chan bool
}

func OpenStream(auth, room string) (*Stream, error) {
	s := &Stream{
		room:     room,
		auth:     auth,
		messages: make(chan *Message),
		closer:   make(chan bool),
	}

	go s.parseMessages()

	return s, nil
}

func (s *Stream) Messages() <-chan *Message {
	return s.messages
}

func (s *Stream) parseMessages() error {
	resp, rerr := s.openStream()
	if rerr != nil {
		return rerr
	}

	decoder := json.NewDecoder(resp.Body)

	var err error
	for {
		select {
		case <-s.closer:
			break
		default:
			var next = &Message{}
			if derr := decoder.Decode(next); derr != nil {
				err = derr
				break
			}

			s.messages <- next
		}
	}

	return err
}

func (s *Stream) Close() {
	s.closer <- true
}

func (s *Stream) openStream() (*http.Response, error) {
	client := new(http.Client)

	req, err := s.request()
	if err != nil {
		return nil, err
	}

	return client.Do(req)
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
