package vcard

import (
	"strings"
)

// IsMalformedN controlla se il campo "N" contiene un valore malformato.
func IsMalformedN(in string) bool {
	// Se ha ":" prima del primo ";" (errore tipico)
	if idxColon := strings.Index(in, ":"); idxColon != -1 {
		idxSemi := strings.Index(in, ";")
		if idxSemi != -1 && idxColon < idxSemi {
			return true
		}
	}

	return false
}
