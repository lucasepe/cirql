package vcard_test

import (
	"testing"

	vcardutil "github.com/lucasepe/cirql/internal/util/vcard"
)

func TestIsMalformedBDAY(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"20240429", false},  // data valida
		{"19991231", false},  // data valida
		{"20240229", false},  // bisestile valida
		{"20230229", true},   // non bisestile → errore
		{"20240431", true},   // aprile ha 30 giorni → errore
		{"2024-04-29", true}, // formato sbagliato
		{"abcd1234", true},   // formato invalido
		{"", true},           // stringa vuota
	}

	for _, tt := range tests {
		got := vcardutil.IsMalformedBDAY(tt.input)
		if got != tt.expected {
			t.Errorf("(%q) got %v, expected: %v", tt.input, got, tt.expected)
		}
	}
}
