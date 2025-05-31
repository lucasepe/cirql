package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lucasepe/cirql/internal/vcards"
)

var _ vcards.CardHandler = (*sqliteHandler)(nil)

type sqliteHandler struct {
	db *sql.DB
}

func (h *sqliteHandler) Handle(c vcards.Card) error {
	if c.Name() == nil {
		return fmt.Errorf("missing name field in vCard")
	}

	fn, gn := c.Name().FamilyName, c.Name().GivenName
	id, err := Lookup(h.db, fn, gn)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if id <= 0 {
		err = Create(h.db, c)
	} else {
		c.SetValue(vcards.FieldUID, strconv.FormatInt(id, 10))
		err = Update(h.db, c)
	}

	return err
}

const (
	uidPrefix = "urn:cirql:contacts:"
)

type scanner interface {
	Scan(dest ...any) error
}

func scanContact(s scanner) (vcards.Card, error) {
	var (
		id       int64
		fn, gn   string
		eml, tel string
		adr      string
		lat, lon float64
		dob      int
	)

	res := make(vcards.Card)
	res.SetValue(vcards.FieldVersion, "3.0")

	err := s.Scan(&id, &fn, &gn, &eml, &tel, &adr, &lat, &lon, &dob)
	if err != nil {
		return res, err
	}

	res.SetValue(vcards.FieldUID, FormatUID(id))
	res.SetValue(vcards.FieldFormattedName, fmt.Sprintf("%s %s", gn, fn))
	res.SetValue(vcards.FieldName, fmt.Sprintf("%s;%s;;;", fn, gn))

	if eml != "" {
		res.SetValue(vcards.FieldEmail, eml)
	}
	if tel != "" {
		res.SetValue(vcards.FieldTelephone, tel)
	}

	if adr != "" {
		res.SetValue(vcards.FieldAddress, adr)
	}

	if dob > 0 {
		res.SetValue(vcards.FieldBirthday, strconv.Itoa(dob))
	}

	if lat > 0 && lon > 0 {
		res.SetValue(vcards.FieldGeolocation,
			fmt.Sprintf("%.6f;%.6f", lat, lon))
	}

	return res, nil
}

func FormatUID(id int64) string {
	return fmt.Sprintf("%s%d", uidPrefix, id)
}

func ParseUID(uid string) (int64, error) {
	if !strings.HasPrefix(uid, uidPrefix) {
		return 0, errors.New("invalid UID: missing prefix")
	}

	rowIDStr := strings.TrimPrefix(uid, uidPrefix)
	rowID, err := strconv.ParseInt(rowIDStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid UID: rowID not a number")
	}

	return rowID, nil
}
