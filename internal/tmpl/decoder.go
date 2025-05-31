package tmpl

import (
	"bufio"
	"io"
	"strings"
)

type Decoder interface {
	Decode(cb func(Message) error) error
}

func NewDecoder(r io.Reader) *decoderImpl {
	return &decoderImpl{s: bufio.NewScanner(r)}
}

type decoderImpl struct {
	s *bufio.Scanner
}

func (dec *decoderImpl) Decode(cb func(m Message) error) error {
	var (
		current *Message
	)

	for dec.s.Scan() {
		line := dec.s.Text()

		if strings.HasPrefix(line, "===") && strings.HasSuffix(line, "===") {
			if current != nil && cb != nil {
				err := cb(*current)
				if err != nil {
					return err
				}
			}

			// Start a new message
			trimmed := strings.TrimSuffix(strings.TrimPrefix(line, "==="), "===")
			recipients, err := Split(strings.TrimSpace(trimmed), ",")
			if err != nil {
				return err
			}

			current = &Message{
				Recipients: recipients,
				Body:       []byte{},
			}
			continue
		}

		if current != nil {
			current.Body = append(current.Body, line...)
		}
	}

	// handle last message if exists
	if current != nil && cb != nil {
		err := cb(*current)
		if err != nil {
			return err
		}
	}

	return dec.s.Err()
}
