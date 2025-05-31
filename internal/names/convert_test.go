package names

import (
	"testing"
)

func TestToVCardN(t *testing.T) {
	tests := []struct {
		name     string
		input    Name
		expected string
	}{
		{
			name:     "Empty struct",
			input:    Name{},
			expected: ";;;;",
		},
		{
			name: "Simple name",
			input: Name{
				FirstName: "Mario",
				LastName:  "Rossi",
			},
			expected: "Rossi;Mario;;;",
		},
		{
			name: "Full name with middle, salutation and suffix",
			input: Name{
				Salutation: "Dr",
				FirstName:  "Angela",
				MiddleName: "Maria",
				LastName:   "De Santis",
				Suffix:     "Jr",
			},
			expected: "De Santis;Angela;Maria;Dr;Jr",
		},
		{
			name: "Name with commas and semicolons",
			input: Name{
				FirstName:  "Luigi;",
				MiddleName: "G.,",
				LastName:   "Bianchi\\",
				Salutation: "Mr.",
				Suffix:     "III",
			},
			expected: "Bianchi\\\\;Luigi\\;;G.\\,;Mr.;III",
		},
		{
			name: "Name with newlines",
			input: Name{
				FirstName:  "Anna\nMaria",
				LastName:   "Verdi",
				Salutation: "Ms",
			},
			expected: "Verdi;Anna\\nMaria;;Ms;",
		},
		{
			name: "Only suffix",
			input: Name{
				Suffix: "PhD",
			},
			expected: ";;;;PhD",
		},
		{
			name: "Only salutation",
			input: Name{
				Salutation: "Sir",
			},
			expected: ";;;Sir;",
		},
		{
			name: "All fields empty strings",
			input: Name{
				Salutation: "",
				FirstName:  "",
				MiddleName: "",
				LastName:   "",
				Suffix:     "",
			},
			expected: ";;;;",
		},
		{
			name: "Complex surname joiners and accents",
			input: Name{
				FirstName:  "José",
				MiddleName: "Luis",
				LastName:   "de la Cruz",
			},
			expected: "de la Cruz;José;Luis;;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToVCardN(tt.input)
			if got != tt.expected {
				t.Errorf("ToVCardN() = %q, want %q", got, tt.expected)
			}
		})
	}
}
