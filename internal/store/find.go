package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lucasepe/cirql/internal/vcards"
)

func FindByID(db *sql.DB, id int64) (vcards.Card, error) {
	row := db.QueryRow(`
		SELECT id, fn, gn, email, phone, adr, lat, lon, birthday
		FROM contacts
		WHERE id = ?
	`, id)

	con, err := scanContact(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return con, fmt.Errorf("contact with ID %d not found", id)
		}
		return con, fmt.Errorf("query contact: %w", err)
	}

	// Fetch categories
	categories, err := loadCategories(db, id)
	if err != nil {
		return con, fmt.Errorf("load categories: %w", err)
	}
	if len(categories) > 0 {
		con.SetValue(vcards.FieldCategories, strings.Join(categories, ","))
	}

	return con, nil
}

func Lookup(db *sql.DB, familyName, givenName string) (id int64, err error) {
	row := db.QueryRow(`
			SELECT id FROM contacts
			WHERE LOWER(fn) = LOWER(?) AND LOWER(gn) = LOWER(?)
			LIMIT 1
		`, familyName, givenName)

	err = row.Scan(&id)
	return
}
