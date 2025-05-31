package vcards

import (
	"strconv"
	"strings"
)

func UID(card Card) string {
	field := card.Get(FieldUID)
	if field == nil {
		return ""
	}

	return field.Value
}

func FN(card Card) string {
	field := card.Get(FieldFormattedName)
	if field == nil {
		return ""
	}

	return strings.TrimSpace(field.Value)
}

func N(card Card) (fn, gn string) {
	field := card.Get(FieldName)
	if field == nil {
		return
	}

	if field.Value == "" {
		return
	}

	// Malformed: colon ':' appears before semicolon ';'.
	idxColon := strings.Index(field.Value, ":")
	if idxColon != -1 {
		idxSemi := strings.Index(field.Value, ";")
		if idxSemi != -1 && idxColon < idxSemi {
			return
		}
	}

	components := strings.Split(field.Value, ";")

	fn = maybeGet(components, 0)
	gn = maybeGet(components, 1)

	return
}

func EMAIL(card Card) string {
	return card.PreferredValue(FieldEmail)
}

func TEL(card Card) string {
	return card.PreferredValue(FieldTelephone)
}

func BDAY(card Card) int {
	field := card.Get(FieldBirthday)
	if field == nil {
		return 0
	}

	val, _ := strconv.Atoi(field.Value)
	return val
}

func GEO(card Card) (lat float64, lon float64) {
	field := card.Get(FieldGeolocation)
	if field == nil {
		return
	}

	components := strings.Split(field.Value, ";")

	lat, _ = strconv.ParseFloat(maybeGet(components, 0), 64)
	lon, _ = strconv.ParseFloat(maybeGet(components, 1), 64)

	return
}

func ADR(c Card) string {
	field := c.Get(FieldAddress)
	if field == nil {
		return ""
	}

	return field.Value
}
