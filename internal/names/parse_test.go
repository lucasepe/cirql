package names

import (
	"testing"
)

func TestParseFullName(t *testing.T) {
	tests := []struct {
		input string
		want  Name
	}{
		{
			input: "Dr. Angela Maria De Santis Jr.",
			want: Name{
				Salutation: "Dr.",
				FirstName:  "Angela",
				MiddleName: "Maria",
				LastName:   "De Santis",
				Suffix:     "Jr.",
			},
		},
		{
			input: "Mr John van der Sar",
			want: Name{
				Salutation: "Mr",
				FirstName:  "John",
				MiddleName: "",
				LastName:   "van der Sar",
				Suffix:     "",
			},
		},
		{
			input: "Ms. Anna-Marie Della Rovere PhD",
			want: Name{
				Salutation: "Ms.",
				FirstName:  "Anna-Marie",
				MiddleName: "",
				LastName:   "Della Rovere",
				Suffix:     "PhD",
			},
		},
		{
			input: "Prof Ludwig van Beethoven",
			want: Name{
				Salutation: "Prof",
				FirstName:  "Ludwig",
				MiddleName: "",
				LastName:   "van Beethoven",
				Suffix:     "",
			},
		},
		{
			input: "Captain Jack Sparrow",
			want: Name{
				Salutation: "Captain",
				FirstName:  "Jack",
				MiddleName: "",
				LastName:   "Sparrow",
				Suffix:     "",
			},
		},
		{
			input: "Lady Gaga",
			want: Name{
				Salutation: "Lady",
				FirstName:  "Gaga",
				MiddleName: "",
				LastName:   "",
				Suffix:     "",
			},
		},
		{
			input: "Sir Ian McKellen",
			want: Name{
				Salutation: "Sir",
				FirstName:  "Ian",
				MiddleName: "",
				LastName:   "McKellen",
				Suffix:     "",
			},
		},
		{
			input: "Madam Curie",
			want: Name{
				Salutation: "Madam",
				FirstName:  "Curie",
				MiddleName: "",
				LastName:   "",
				Suffix:     "",
			},
		},
		{
			input: "Rev. Martin Luther King Jr",
			want: Name{
				Salutation: "Rev.",
				FirstName:  "Martin",
				MiddleName: "Luther",
				LastName:   "King",
				Suffix:     "Jr",
			},
		},
		{
			input: "Angela",
			want: Name{
				Salutation: "",
				FirstName:  "Angela",
				MiddleName: "",
				LastName:   "",
				Suffix:     "",
			},
		},
		{
			input: "John Jacob Jingleheimer Schmidt",
			want: Name{
				Salutation: "",
				FirstName:  "John",
				MiddleName: "Jacob Jingleheimer",
				LastName:   "Schmidt",
				Suffix:     "",
			},
		},
		{
			input: "Dr. José Luis Rodríguez Zapatero",
			want: Name{
				Salutation: "Dr.",
				FirstName:  "José",
				MiddleName: "Luis Rodríguez",
				LastName:   "Zapatero",
				Suffix:     "",
			},
		},
		{
			input: "Ms Mary-Kate Olsen",
			want: Name{
				Salutation: "Ms",
				FirstName:  "Mary-Kate",
				MiddleName: "",
				LastName:   "Olsen",
				Suffix:     "",
			},
		},
		{
			input: "Mr. Jean-Luc Picard",
			want: Name{
				Salutation: "Mr.",
				FirstName:  "Jean-Luc",
				MiddleName: "",
				LastName:   "Picard",
				Suffix:     "",
			},
		},
		{
			input: "Sgt. Pepper",
			want: Name{
				Salutation: "Sgt.",
				FirstName:  "Pepper",
				MiddleName: "",
				LastName:   "",
				Suffix:     "",
			},
		},
		{
			input: "  ", // empty/whitespace
			want:  Name{},
		},
		{
			input: "",
			want:  Name{},
		},
	}

	for _, tt := range tests {
		got := ParseFullName(tt.input)
		if got != tt.want {
			t.Errorf("ParseFullName(%q) = %+v; want %+v", tt.input, got, tt.want)
		}
	}
}
