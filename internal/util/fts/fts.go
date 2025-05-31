package ftsutil

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// FormatFTSQuery prepara una query compatibile con SQLite FTS5.
// - Normalizza UTF-8 (es. "café" → "cafe")
// - Rimuove caratteri pericolosi
// - Supporta "OR" tra termini
// - Applica '*' per ricerca prefissata
func FormatFTSQuery(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}

	// Normalizza UTF8 rimuovendo accenti (NFD)
	// "Café" → "Café" → "Cafe"
	t := norm.NFD.String(input)
	var sb strings.Builder
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			// ignora i caratteri "mark" (accenti)
			continue
		}
		sb.WriteRune(r)
	}
	normalized := sb.String()

	// Tokenizza e processa OR
	tokens := strings.Fields(normalized)
	var output []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if strings.ToUpper(token) == "OR" && i > 0 && i < len(tokens)-1 {
			// Costruisce "foo* OR bar*"
			prev := strings.Trim(tokenSanitize(tokens[i-1]), "*") + "*"
			next := strings.Trim(tokenSanitize(tokens[i+1]), "*") + "*"
			output = output[:len(output)-1] // rimuove il precedente
			output = append(output, prev+" OR "+next)
			i++ // salta il prossimo
			continue
		}

		output = append(output, tokenSanitize(token)+"*")
	}

	return strings.Join(output, " ")
}

// tokenSanitize rimuove caratteri pericolosi per MATCH
func tokenSanitize(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '"', '\'', ':', '-', '+', '*', '(', ')':
			return -1
		default:
			return r
		}
	}, s)
}
