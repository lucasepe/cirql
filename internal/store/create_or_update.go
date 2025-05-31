package store

import (
	"database/sql"
	"fmt"

	"github.com/lucasepe/cirql/internal/names"
	"github.com/lucasepe/cirql/internal/vcards"
)

type Result int

const (
	Failed Result = iota
	Created
	Updated
	Skipped
)

func CreateOrUpdate(db *sql.DB, c vcards.Card, override bool) (Result, error) {
	fullName := vcards.FN(c)
	if fullName == "" {
		return Failed, fmt.Errorf("invalid vcard: missing mandatory field value FN")
	}

	if fn, gn := vcards.N(c); fn == "" && gn == "" {
		c.Set(vcards.FieldName, &vcards.Field{
			Value: names.ToVCardN(names.ParseFullName(fullName)),
		})
	}

	fn, gn := vcards.N(c)
	if fn == "" && gn == "" {
		return Failed,
			fmt.Errorf("invalid vcard %q: unable to derive field value N", fullName)
	}

	id, err := Lookup(db, fn, gn)
	if err != nil && err != sql.ErrNoRows {
		return Failed, err
	}

	if id <= 0 {
		err = Create(db, c)
		if err != nil {
			return Failed, err
		}
		return Created, nil
	}

	if !override {
		return Skipped, nil
	}

	c.SetValue(vcards.FieldUID, FormatUID(id))
	if err = Update(db, c); err != nil {
		return Failed, err
	}

	return Updated, nil
}
