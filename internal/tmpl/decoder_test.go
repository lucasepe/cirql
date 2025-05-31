package tmpl_test

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/lucasepe/cirql/internal/tmpl"
)

func TestMessageDecoder(t *testing.T) {
	text := `=== email:alice@example.com ===
Hello Alice!

=== email:bob@example.com,tel:+393339998877 ===
Hi Bob!

`
	want := []tmpl.Message{
		{Recipients: []tmpl.Recipient{
			{Type: tmpl.Email, Value: "alice@example.com"},
		}, Body: []byte("Hello Alice!")},
		{Recipients: []tmpl.Recipient{
			{Type: tmpl.Email, Value: "bob@example.com"},
			{Type: tmpl.Phone, Value: "+393339998877"},
		}, Body: []byte("Hi Bob!")},
	}

	var got []tmpl.Message

	handler := func(m tmpl.Message) error {
		got = append(got, m)
		return nil
	}

	dec := tmpl.NewDecoder(strings.NewReader(text))
	if err := dec.Decode(handler); err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %+v, got:   %+v", want, got)
	}
}
