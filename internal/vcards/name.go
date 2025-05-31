package vcards

import "strings"

// Names returns names of the card.
func (c Card) Names() []*Name {
	ns := c[FieldName]
	if ns == nil {
		return nil
	}

	names := make([]*Name, len(ns))
	for i, n := range ns {
		names[i] = newName(n)
	}
	return names
}

// Name returns the preferred name of the card. If it isn't specified, it
// returns nil.
func (c Card) Name() *Name {
	n := c.Preferred(FieldName)
	if n == nil {
		return nil
	}
	if strings.TrimSpace(n.Value) == "" {
		return nil
	}

	return newName(n)
}

// AddName adds the specified name to the list of names.
func (c Card) AddName(name *Name) {
	c.Add(FieldName, name.field())
}

// SetName replaces the list of names with the single specified name.
func (c Card) SetName(name *Name) {
	c.Set(FieldName, name.field())
}

// Name contains an object's name components.
type Name struct {
	*Field

	FamilyName      string
	GivenName       string
	AdditionalName  string
	HonorificPrefix string
	HonorificSuffix string
}

func newName(field *Field) *Name {
	components := strings.Split(field.Value, ";")
	return &Name{
		field,
		maybeGet(components, 0),
		maybeGet(components, 1),
		maybeGet(components, 2),
		maybeGet(components, 3),
		maybeGet(components, 4),
	}
}

func (n *Name) field() *Field {
	if n.Field == nil {
		n.Field = new(Field)
	}
	n.Field.Value = strings.Join([]string{
		n.FamilyName,
		n.GivenName,
		n.AdditionalName,
		n.HonorificPrefix,
		n.HonorificSuffix,
	}, ";")
	return n.Field
}
