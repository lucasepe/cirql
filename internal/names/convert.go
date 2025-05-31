package names

import "strings"

func ToVCardN(n Name) string {
	// I campi vanno messi in questo ordine:
	// LastName;FirstName;MiddleName;Salutation;Suffix
	fields := []string{
		escapeVCardValue(n.LastName),
		escapeVCardValue(n.FirstName),
		escapeVCardValue(n.MiddleName),
		escapeVCardValue(n.Salutation),
		escapeVCardValue(n.Suffix),
	}

	return strings.Join(fields, ";")
}

// escapeVCardValue fa escaping dei caratteri speciali (\, ;, ,) nel valore vCard
func escapeVCardValue(s string) string {
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		";", "\\;",
		",", "\\,",
		"\n", "\\n",
	)
	return replacer.Replace(s)
}
