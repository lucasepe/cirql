package vcards

import "strings"

// Sex is an object's biological sex.
type Sex string

const (
	SexUnspecified Sex = ""
	SexFemale      Sex = "F"
	SexMale        Sex = "M"
	SexOther       Sex = "O"
	SexNone        Sex = "N"
	SexUnknown     Sex = "U"
)

// Gender returns this card's gender.
func (c Card) Gender() (sex Sex, identity string) {
	v := c.Value(FieldGender)
	parts := strings.SplitN(v, ";", 2)
	return Sex(strings.ToUpper(parts[0])), maybeGet(parts, 1)
}

// SetGender sets this card's gender.
func (c Card) SetGender(sex Sex, identity string) {
	v := string(sex)
	if identity != "" {
		v += ";" + identity
	}
	c.SetValue(FieldGender, v)
}
