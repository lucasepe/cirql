// Package vcard implements the vCard format, defined in RFC 6350.
package vcards

import (
	"strconv"
	"strings"
	"time"
)

// MIME type and file extension for VCard, defined in RFC 6350 section 10.1.
const (
	MIMEType  = "text/vcard"
	Extension = "vcf"
)

const timestampLayout = "20060102T150405Z"

// Kind is an object's kind.
type Kind string

// Values for FieldKind.
const (
	KindIndividual   Kind = "individual"
	KindGroup        Kind = "group"
	KindOrganization Kind = "org"
	KindLocation     Kind = "location"
)

// A Card is an address book entry.
type Card map[string][]*Field

// Get returns the first field of the card for the given property. If there is
// no such field, it returns nil.
func (c Card) Get(k string) *Field {
	fields := c[k]
	if len(fields) == 0 {
		return nil
	}
	return fields[0]
}

// Add adds the k, f pair to the list of fields. It appends to any existing
// fields.
func (c Card) Add(k string, f *Field) {
	c[k] = append(c[k], f)
}

// Set sets the key k to the single field f. It replaces any existing field.
func (c Card) Set(k string, f *Field) {
	c[k] = []*Field{f}
}

// Preferred returns the preferred field of the card for the given property.
func (c Card) Preferred(k string) *Field {
	fields := c[k]
	if len(fields) == 0 {
		return nil
	}

	field := fields[0]
	min := 100
	for _, f := range fields {
		n := 100
		if pref := f.Params.Get(ParamPreferred); pref != "" {
			n, _ = strconv.Atoi(pref)
		} else if f.Params.HasType("pref") {
			// Apple Contacts adds "pref" to the TYPE param
			n = 1
		}

		if n < min {
			min = n
			field = f
		}
	}
	return field
}

// Value returns the first field value of the card for the given property. If
// there is no such field, it returns an empty string.
func (c Card) Value(k string) string {
	f := c.Get(k)
	if f == nil {
		return ""
	}
	return f.Value
}

// AddValue adds the k, v pair to the list of field values. It appends to any
// existing values.
func (c Card) AddValue(k, v string) {
	c.Add(k, &Field{Value: v})
}

// SetValue sets the field k to the single value v. It replaces any existing
// value.
func (c Card) SetValue(k, v string) {
	c.Set(k, &Field{Value: v})
}

// PreferredValue returns the preferred field value of the card.
func (c Card) PreferredValue(k string) string {
	f := c.Preferred(k)
	if f == nil {
		return ""
	}
	return f.Value
}

// Values returns a list of values for a given property.
func (c Card) Values(k string) []string {
	fields := c[k]
	if fields == nil {
		return nil
	}

	values := make([]string, len(fields))
	for i, f := range fields {
		values[i] = f.Value
	}
	return values
}

// Kind returns the kind of the object represented by this card. If it isn't
// specified, it returns the default: KindIndividual.
func (c Card) Kind() Kind {
	kind := strings.ToLower(c.Value(FieldKind))
	if kind == "" {
		return KindIndividual
	}
	return Kind(kind)
}

// SetKind sets the kind of the object represented by this card.
func (c Card) SetKind(kind Kind) {
	c.SetValue(FieldKind, string(kind))
}

// FormattedNames returns formatted names of the card. The length of the result
// is always greater or equal to 1.
func (c Card) FormattedNames() []*Field {
	fns := c[FieldFormattedName]
	if len(fns) == 0 {
		return []*Field{{Value: ""}}
	}
	return fns
}

// Categories returns category information about the card, also known as "tags".
func (c Card) Categories() []string {
	return strings.Split(c.PreferredValue(FieldCategories), ",")
}

// SetCategories sets category information about the card.
func (c Card) SetCategories(categories []string) {
	c.SetValue(FieldCategories, strings.Join(categories, ","))
}

// Revision returns revision information about the current card.
func (c Card) Revision() (time.Time, error) {
	rev := c.Value(FieldRevision)
	if rev == "" {
		return time.Time{}, nil
	}
	return time.Parse(timestampLayout, rev)
}

// SetRevision sets revision information about the current card.
func (c Card) SetRevision(t time.Time) {
	c.SetValue(FieldRevision, t.Format(timestampLayout))
}

func maybeGet(l []string, i int) string {
	if i < len(l) {
		return strings.TrimSpace(l[i])
	}
	return ""
}
