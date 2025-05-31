package names

import (
	"strings"
)

type Name struct {
	Salutation string
	FirstName  string
	MiddleName string
	LastName   string
	Suffix     string
}

func ParseFullName(fullName string) Name {
	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return Name{}
	}

	parts := strings.Fields(fullName)
	n := len(parts)
	if n == 0 {
		return Name{}
	}

	res := Name{}

	// Detect salutation (prefix)
	if isSalutation(parts[0]) {
		res.Salutation = parts[0] // normalizeSalutation(parts[0])
		parts = parts[1:]
		n--
		if n == 0 {
			return res
		}
	}

	// Assign FirstName
	res.FirstName = parts[0]

	// Handle suffix at the end
	if n > 1 {
		lastWordClean := cleanWord(parts[n-1])
		if _, ok := suffixMap[lastWordClean]; ok {
			res.Suffix = parts[n-1]
			parts = parts[:n-1]
			n--
		}
	}

	// Process surname (possibly composed)
	if n == 1 {
		// Only first name, no middle or last name
		return res
	}

	i := n - 1
	surnameParts := []string{parts[i]}
	i--

	for i > 0 {
		w := cleanWord(parts[i])
		if surnameJoiners[w] {
			surnameParts = append([]string{parts[i]}, surnameParts...)
			i--
		} else {
			break
		}
	}

	// Middle name(s) are whatever is left between first name and surname
	if i >= 1 {
		res.MiddleName = strings.Join(parts[1:i+1], " ")
	} else {
		res.MiddleName = ""
	}

	res.LastName = strings.Join(surnameParts, " ")

	return res
}

var salutationMap = map[string]struct{}{
	"mr":      {},
	"mrs":     {},
	"ms":      {},
	"dr":      {},
	"prof":    {},
	"capt":    {},
	"captain": {},
	"lt":      {},
	"col":     {},
	"gen":     {},
	"sgt":     {},
	"rev":     {},
	"sir":     {},
	"madam":   {},
	"lord":    {},
	"lady":    {},
}

var suffixMap = map[string]struct{}{
	"jr":  {},
	"sr":  {},
	"ii":  {},
	"iii": {},
	"phd": {},
	"md":  {},
	"esq": {},
}

var surnameJoiners = map[string]bool{
	"de": true, "del": true, "della": true, "di": true,
	"van": true, "von": true, "der": true, "den": true, "da": true,
	"la": true, "le": true, "du": true, "st.": true, "st": true,
}

// cleanWord normalizes the word (lowercase, trims trailing dots)
func cleanWord(s string) string {
	s = strings.ToLower(s)
	return strings.TrimSuffix(s, ".")
}

// isSalutation controlla se la parola è un titolo riconosciuto.
// Rimuove eventuale punto finale e fa match case-insensitive.
func isSalutation(word string) bool {
	w := cleanWord(word)
	_, ok := salutationMap[w]
	return ok
}
