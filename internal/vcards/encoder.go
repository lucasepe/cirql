package vcards

import (
	"errors"
	"io"
	"sort"
	"strings"
)

// An Encoder formats cards.
type Encoder struct {
	w io.Writer
}

// NewEncoder creates a new Encoder that writes cards to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

// Encode formats a card. The card must have a FieldVersion field.
func (enc *Encoder) Encode(c Card) error {
	begin := "BEGIN:VCARD\r\n"
	if _, err := io.WriteString(enc.w, begin); err != nil {
		return err
	}

	version := c.Get(FieldVersion)
	if version == nil {
		return errors.New("vcard: VERSION field missing")
	}

	priorityOrder := []string{
		FieldVersion,
		FieldRevision,
		FieldUID,
		FieldFormattedName,
		FieldName, FieldBirthday,
		FieldEmail,
		FieldTelephone,
		FieldAddress,
		FieldGeolocation,
	}

	printFunc := func(c Card, k string) error {
		fields := c[k]

		for _, f := range fields {
			_, err := io.WriteString(enc.w, formatLine(k, f)+"\n")
			if err != nil {
				return err
			}
		}

		return nil
	}

	// 1. Stampa i campi prioritari in ordine
	for _, key := range priorityOrder {
		err := printFunc(c, key)
		if err != nil {
			return err
		}
	}

	// 2. Stampa i campi rimanenti (ordinati alfabeticamente)
	seen := make(map[string]bool)
	for _, k := range priorityOrder {
		seen[k] = true
	}

	var keys []string
	for k := range c {
		if !seen[k] {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		err := printFunc(c, k)
		if err != nil {
			return err
		}
	}

	end := "END:VCARD\r\n"
	_, err := io.WriteString(enc.w, end)
	return err
}

func formatLine(key string, field *Field) string {
	var s string

	if field.Group != "" {
		s += field.Group + "."
	}
	s += key

	var keys []string
	for k := range field.Params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, pk := range keys {
		for _, pv := range field.Params[pk] {
			s += ";" + formatParam(pk, pv)
		}
	}

	s += ":" + formatValue(field.Value)
	return s
}

func formatParam(k, v string) string {
	return k + "=" + formatValue(v)
}

var valueFormatter = strings.NewReplacer("\\", "\\\\", "\n", "\\n", ",", "\\,")

func formatValue(v string) string {
	return valueFormatter.Replace(v)
}
