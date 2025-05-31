// Package vcard implements the vCard format, defined in RFC 6350.
package vcards

import (
	"strings"
)

// Card properties.
const (
	// General Properties
	FieldSource = "SOURCE"
	FieldKind   = "KIND"
	FieldXML    = "XML"

	// Identification Properties
	FieldFormattedName = "FN"
	FieldName          = "N"
	FieldNickname      = "NICKNAME"
	FieldPhoto         = "PHOTO"
	FieldBirthday      = "BDAY"
	FieldAnniversary   = "ANNIVERSARY"
	FieldGender        = "GENDER"

	// Delivery Addressing Properties
	FieldAddress = "ADR"

	// Communications Properties
	FieldTelephone = "TEL"
	FieldEmail     = "EMAIL"
	FieldIMPP      = "IMPP" // Instant Messaging and Presence Protocol
	FieldLanguage  = "LANG"

	// Geographical Properties
	FieldTimezone    = "TZ"
	FieldGeolocation = "GEO"

	// Organizational Properties
	FieldTitle        = "TITLE"
	FieldRole         = "ROLE"
	FieldLogo         = "LOGO"
	FieldOrganization = "ORG"
	FieldMember       = "MEMBER"
	FieldRelated      = "RELATED"

	// Explanatory Properties
	FieldCategories   = "CATEGORIES"
	FieldNote         = "NOTE"
	FieldProductID    = "PRODID"
	FieldRevision     = "REV"
	FieldSound        = "SOUND"
	FieldUID          = "UID"
	FieldClientPIDMap = "CLIENTPIDMAP"
	FieldURL          = "URL"
	FieldVersion      = "VERSION"

	// Security Properties
	FieldKey = "KEY"

	// Calendar Properties
	FieldFreeOrBusyURL      = "FBURL"
	FieldCalendarAddressURI = "CALADRURI"
	FieldCalendarURI        = "CALURI"
)

// Card property parameters.
const (
	ParamLanguage      = "LANGUAGE"
	ParamValue         = "VALUE"
	ParamPreferred     = "PREF"
	ParamAltID         = "ALTID"
	ParamPID           = "PID"
	ParamType          = "TYPE"
	ParamMediaType     = "MEDIATYPE"
	ParamCalendarScale = "CALSCALE"
	ParamSortAs        = "SORT-AS"
	ParamGeolocation   = "GEO"
	ParamTimezone      = "TZ"
)

// Values for ParamType.
const (
	// Generic
	TypeHome = "home"
	TypeWork = "work"

	// For FieldTelephone
	TypeText      = "text"
	TypeVoice     = "voice" // Default
	TypeFax       = "fax"
	TypeCell      = "cell"
	TypeVideo     = "video"
	TypePager     = "pager"
	TypeTextPhone = "textphone"

	// For FieldRelated
	TypeContact      = "contact"
	TypeAcquaintance = "acquaintance"
	TypeFriend       = "friend"
	TypeMet          = "met"
	TypeCoWorker     = "co-worker"
	TypeColleague    = "colleague"
	TypeCoResident   = "co-resident"
	TypeNeighbor     = "neighbor"
	TypeChild        = "child"
	TypeParent       = "parent"
	TypeSibling      = "sibling"
	TypeSpouse       = "spouse"
	TypeKin          = "kin"
	TypeMuse         = "muse"
	TypeCrush        = "crush"
	TypeDate         = "date"
	TypeSweetheart   = "sweetheart"
	TypeMe           = "me"
	TypeAgent        = "agent"
	TypeEmergency    = "emergency"
)

// A field contains a value and some parameters.
type Field struct {
	Value  string
	Params Params
	Group  string
}

// Params is a set of field parameters.
type Params map[string][]string

// Get returns the first value with the key k. It returns an empty string if
// there is no such value.
func (p Params) Get(k string) string {
	values := p[k]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// Add adds the k, v pair to the list of parameters. It appends to any existing
// values.
func (p Params) Add(k, v string) {
	p[k] = append(p[k], v)
}

// Set sets the parameter k to the single value v. It replaces any existing
// value.
func (p Params) Set(k, v string) {
	p[k] = []string{v}
}

// Types returns the field types.
func (p Params) Types() []string {
	types := p[ParamType]
	list := make([]string, len(types))
	for i, t := range types {
		list[i] = strings.ToLower(t)
	}
	return list
}

// HasType returns true if and only if the field have the provided type.
func (p Params) HasType(t string) bool {
	for _, tt := range p[ParamType] {
		if strings.EqualFold(t, tt) {
			return true
		}
	}
	return false
}
