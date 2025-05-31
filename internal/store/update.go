package store

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lucasepe/cirql/internal/vcards"
)

func Update(db *sql.DB, o vcards.Card) error {
	id, err := ParseUID(vcards.UID(o))
	if err != nil {
		return err
	}
	if id <= 0 {
		return fmt.Errorf("contact ID must be > 0")
	}

	if o.Name() == nil {
		return fmt.Errorf("missing name field in vCard")
	}

	fn, gn := vcards.N(o)
	adr := vcards.ADR(o)
	lat, lon := vcards.GEO(o)
	eml, tel := vcards.EMAIL(o), vcards.TEL(o)
	dob := vcards.BDAY(o)

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE contacts
		SET fn = ?, gn = ?, adr = ?, lat = ?, lon = ?, email = ?, phone = ?, birthday = ?
		WHERE id = ?
	`, fn, gn, adr, lat, lon, eml, tel, dob, id)
	if err != nil {
		return fmt.Errorf("update contact: %w", err)
	}

	// Elimina le categorie esistenti e reinserisci quelle nuove
	_, err = tx.Exec(`DELETE FROM categories WHERE contact_id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete old categories: %w", err)
	}

	if all := o.Categories(); len(all) > 0 {
		for _, cat := range all {
			cat = strings.TrimSpace(cat)
			if cat != "" {
				_, err = tx.Exec(`INSERT INTO categories (contact_id, category) VALUES (?, ?)`, id, cat)
				if err != nil {
					return fmt.Errorf("insert category: %w", err)
				}
			}
		}
	}

	// Aggiorna l'indice FTS
	_, err = tx.Exec(`
		UPDATE contacts_fts
		SET fn = ?, gn = ?, adr = ?, email = ?, phone = ?
		WHERE rowid = ?
	`, fn, gn, adr, eml, tel, id)
	if err != nil {
		return fmt.Errorf("update FTS index: %w", err)
	}

	return tx.Commit()
}
