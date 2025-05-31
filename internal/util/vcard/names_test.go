package vcard

import (
	"testing"
)

func TestIsMalformedN(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		name     string
	}{
		{
			name:     "Malformed: SORT-AS in value with colon before semicolon",
			input:    "SORT-AS=Angela De Santis:De Santis;Angela",
			expected: true,
		},
		{
			name:     "Malformed: colon appears before semicolon",
			input:    "something:strange;value",
			expected: true,
		},
		{
			name:     "Well-formed: full structured name",
			input:    "De Santis;Angela;;;",
			expected: false,
		},
		{
			name:     "Well-formed: only family name",
			input:    "Rossi;;;;",
			expected: false,
		},
		{
			name:     "Not malformed: colon present, no semicolon",
			input:    "SORT-AS:some value without semicolon",
			expected: false,
		},
		{
			name:     "Well-formed: empty input",
			input:    "",
			expected: false,
		},
		{
			name:     "Well-formed: only semicolons",
			input:    ";;;;",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsMalformedN(tc.input)
			if got != tc.expected {
				t.Errorf("IsMalformedN(%q) = %v; expected %v", tc.input, got, tc.expected)
			}
		})
	}
}
