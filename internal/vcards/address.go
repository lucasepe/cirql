package vcards

import "strings"

// Addresses returns addresses of the card.
func (c Card) Addresses() []*Address {
	adrs := c[FieldAddress]
	if adrs == nil {
		return nil
	}

	addresses := make([]*Address, len(adrs))
	for i, adr := range adrs {
		addresses[i] = newAddress(adr)
	}
	return addresses
}

// Address returns the preferred address of the card. If it isn't specified, it
// returns nil.
func (c Card) Address() *Address {
	adr := c.Preferred(FieldAddress)
	if adr == nil {
		return nil
	}
	return newAddress(adr)
}

// AddAddress adds an address to the list of addresses.
func (c Card) AddAddress(address *Address) {
	c.Add(FieldAddress, address.field())
}

// SetAddress replaces the list of addresses with the single specified address.
func (c Card) SetAddress(address *Address) {
	c.Set(FieldAddress, address.field())
}

// An Address is a delivery address.
type Address struct {
	*Field

	PostOfficeBox   string
	ExtendedAddress string // e.g., apartment or suite number
	StreetAddress   string
	Locality        string // e.g., city
	Region          string // e.g., state or province
	PostalCode      string
	Country         string
}

func newAddress(field *Field) *Address {
	components := strings.Split(field.Value, ";")
	return &Address{
		field,
		maybeGet(components, 0),
		maybeGet(components, 1),
		maybeGet(components, 2),
		maybeGet(components, 3),
		maybeGet(components, 4),
		maybeGet(components, 5),
		maybeGet(components, 6),
	}
}

func (a *Address) field() *Field {
	if a.Field == nil {
		a.Field = new(Field)
	}
	a.Field.Value = strings.Join([]string{
		a.PostOfficeBox,
		a.ExtendedAddress,
		a.StreetAddress,
		a.Locality,
		a.Region,
		a.PostalCode,
		a.Country,
	}, ";")
	return a.Field
}
