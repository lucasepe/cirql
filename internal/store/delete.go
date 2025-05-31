package store

import (
	"database/sql"
	"fmt"
)

// Delete elimina un contatto e aggiorna l'indice FTS
func Delete(db *sql.DB, contactID int64) error {
	// Inizia la transazione
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Elimina le categorie associate al contatto
	_, err = tx.Exec(`DELETE FROM categories WHERE contact_id = ?`, contactID)
	if err != nil {
		return fmt.Errorf("delete categories: %w", err)
	}

	// Elimina il contatto dalla tabella contacts
	_, err = tx.Exec(`DELETE FROM contacts WHERE id = ?`, contactID)
	if err != nil {
		return fmt.Errorf("delete contact: %w", err)
	}

	// Elimina l'indice FTS per il contatto
	_, err = tx.Exec(`DELETE FROM contacts_fts WHERE rowid = ?`, contactID)
	if err != nil {
		return fmt.Errorf("delete from FTS index: %w", err)
	}

	// Commetti la transazione
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
