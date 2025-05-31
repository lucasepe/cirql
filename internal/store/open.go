package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucasepe/xdg"

	_ "modernc.org/sqlite"
)

const (
	defaultFilename = "cirql.db"
)

func Open() (*sql.DB, error) {
	dbPath := filepath.Join(xdg.ConfigDir(), "CirQL")
	if err := os.MkdirAll(dbPath, 0755); err != nil {
		return nil, err
	}
	dbPath = filepath.Join(dbPath, defaultFilename)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback() // in caso di errore

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY,
		fn TEXT,
		gn TEXT,
		adr TEXT,
		lat REAL,
		lon REAL,
		email TEXT,
		phone TEXT,
		birthday INTEGER
	);`)
	if err != nil {
		return nil, fmt.Errorf("create contacts table: %w", err)
	}

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		contact_id INTEGER NOT NULL,
		category TEXT NOT NULL COLLATE NOCASE,
		UNIQUE(contact_id, category),
		FOREIGN KEY(contact_id) REFERENCES contacts(id)
	);`)
	if err != nil {
		return nil, fmt.Errorf("create categories table: %w", err)
	}

	_, err = tx.Exec(`
	CREATE VIRTUAL TABLE IF NOT EXISTS contacts_fts USING fts5(
		fn,
		gn,
		adr,
		email,
		phone,
		content='contacts',
		content_rowid='id'
	);`)
	if err != nil {
		return nil, fmt.Errorf("create contacts_fts virtual table: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return db, nil
}
