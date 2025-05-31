package tmpl

import (
	"fmt"
	"strings"
)

type RecipientType string

const (
	Email RecipientType = "email"
	Phone               = "phone"
)

type Recipient struct {
	Type  RecipientType
	Value string
}

type Message struct {
	Recipients []Recipient
	Body       []byte
}

func (r *Recipient) MarshalText() ([]byte, error) {
	switch r.Type {
	case Email:
		return fmt.Appendf(nil, "email:%s", r.Value), nil
	case Phone:
		return fmt.Appendf(nil, "tel:%s", r.Value), nil
	default:
		return nil, fmt.Errorf("unknown recipient type: %q", r.Type)
	}
}

func (r *Recipient) UnmarshalText(text []byte) error {
	s := string(text)
	if strings.HasPrefix(s, "email:") {
		r.Type = Email
		r.Value = strings.TrimPrefix(s, "email:")
	} else if strings.HasPrefix(s, "tel:") {
		r.Type = Phone
		r.Value = strings.TrimPrefix(s, "tel:")
	} else {
		return fmt.Errorf("invalid recipient format: %q", s)
	}
	return nil
}
