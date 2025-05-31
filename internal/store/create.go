package store

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lucasepe/cirql/internal/vcards"
)

func Create(db *sql.DB, o vcards.Card) error {
	if o.Name() == nil {
		return fmt.Errorf("missing name field in vCard")
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	fn, gn := vcards.N(o)
	adr := vcards.ADR(o)
	lat, lon := vcards.GEO(o)
	eml, tel := vcards.EMAIL(o), vcards.TEL(o)
	dob := vcards.BDAY(o)

	// Insert the new contact
	res, err := tx.Exec(`
		INSERT INTO contacts (fn, gn, adr, lat, lon, email, phone, birthday)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, fn, gn, adr, lat, lon, eml, tel, dob)
	if err != nil {
		return fmt.Errorf("insert contact: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert id: %w", err)
	}

	// Insert categories
	if all := o.Categories(); len(all) > 0 {
		for _, cat := range all {
			cat = strings.ToLower(strings.TrimSpace(cat))
			if cat != "" {
				_, err = tx.Exec(`INSERT OR IGNORE INTO categories (contact_id, category) VALUES (?, ?)`, id, cat)
				if err != nil {
					return fmt.Errorf("insert category: %w", err)
				}
			}
		}
	}

	// Add to Full-Text Search (FTS)
	_, err = tx.Exec(`
		INSERT INTO contacts_fts(fn, gn, adr, email, phone)
		VALUES (?, ?, ?, ?, ?)
	`, fn, gn, adr, eml, tel)
	if err != nil {
		return fmt.Errorf("insert into FTS index: %w", err)
	}

	return tx.Commit()
}
