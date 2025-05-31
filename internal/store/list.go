package store

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	ftsutil "github.com/lucasepe/cirql/internal/util/fts"
	"github.com/lucasepe/cirql/internal/vcards"
)

type ListOptions struct {
	Match          string
	DaysUntilBirth int
	Categories     []string
	Handler        vcards.CardHandler
}

func List(db *sql.DB, opts ListOptions) (int, error) {
	var (
		rows         *sql.Rows
		err          error
		args         []any
		whereClauses []string
	)

	// Build filters
	if opts.Match != "" {
		ftsQuery := fmt.Sprintf("SELECT rowid FROM contacts_fts WHERE contacts_fts MATCH %q",
			ftsutil.FormatFTSQuery(opts.Match))
		whereClauses = append(whereClauses,
			fmt.Sprintf("contacts.id IN (%s)", ftsQuery))
	}

	if opts.DaysUntilBirth > 0 {
		// Calcola MMDD per oggi e i prossimi N giorni
		today := time.Now()
		mmddList := make([]string, 0, opts.DaysUntilBirth+1)
		for i := 0; i <= opts.DaysUntilBirth; i++ {
			day := today.AddDate(0, 0, i)
			mmddList = append(mmddList, day.Format("0102")) // MMDD
		}

		// Costruisci WHERE IN
		placeholders := strings.Repeat("?,", len(mmddList))
		placeholders = placeholders[:len(placeholders)-1] // rimuovi ultima virgola
		whereClauses = append(whereClauses, fmt.Sprintf(`
			substr(CAST(birthday AS TEXT), 5, 4) IN (%s)
		`, placeholders))

		for _, v := range mmddList {
			args = append(args, v)
		}
	}

	if len(opts.Categories) > 0 {
		placeholders := strings.Repeat("?,", len(opts.Categories))
		placeholders = placeholders[:len(placeholders)-1]

		whereClauses = append(whereClauses, fmt.Sprintf(`
        contacts.id IN (
            SELECT contact_id
            FROM categories
            WHERE category COLLATE NOCASE IN (%s)
        )
    `, placeholders))

		for _, c := range opts.Categories {
			args = append(args, c)
		}
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Final query
	querySQL := `
		SELECT DISTINCT contacts.id, contacts.fn, contacts.gn, 
			   contacts.email, contacts.phone,
		       contacts.adr, contacts.lat, contacts.lon,
		       contacts.birthday
	    FROM contacts ` + whereSQL + `
		ORDER BY contacts.gn, contacts.fn
	`

	rows, err = db.Query(querySQL, args...)
	if err != nil {
		return 0, fmt.Errorf("query contacts: %w", err)
	}
	defer rows.Close()

	rev := time.Now().UTC().Format("20060102T150405Z")

	total := 0
	for rows.Next() {
		con, err := scanContact(rows)
		if err != nil {
			return total, fmt.Errorf("scan contact: %w", err)
		}

		// Fetch categories
		id, err := ParseUID(vcards.UID(con))
		if err != nil {
			return total, err
		}

		categories, err := loadCategories(db, id)
		if err != nil {
			return total, fmt.Errorf("load categories: %w", err)
		}
		if len(categories) > 0 {
			con.SetValue(vcards.FieldCategories, strings.Join(categories, ","))
		}

		// Set card revision
		con.SetValue(vcards.FieldRevision, rev)

		// Handle record
		err = opts.Handler.Handle(con)
		if err != nil {
			return total, err
		}

		total = total + 1
	}

	err = rows.Err()
	return total, err
}

func loadCategories(db *sql.DB, contactID int64) ([]string, error) {
	rows, err := db.Query(`
		SELECT category
		FROM categories
		WHERE contact_id = ? COLLATE NOCASE
		ORDER BY category ASC
	`, contactID)
	if err != nil {
		return nil, fmt.Errorf("query categories: %w", err)
	}
	defer rows.Close()

	var cats []string
	for rows.Next() {
		var cat string
		if err := rows.Scan(&cat); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		cats = append(cats, cat)
	}

	return cats, nil
}

// formatFTSQuery prepara una query compatibile con FTS5
// Aggiunge wildcard '*' per cercare prefissi, rimuove caratteri speciali per evitare errori
func formatFTSQuery(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}

	// Rimuove caratteri che possono rompere la query MATCH
	clean := strings.Map(func(r rune) rune {
		switch r {
		case '"', '\'', ':', '-', '+', '*', '(', ')':
			return -1
		default:
			return r
		}
	}, input)

	// Suddivide per parole e aggiunge '*' per ricerca prefissata
	tokens := strings.Fields(clean)
	for i, token := range tokens {
		tokens[i] = token + "*"
	}

	// Combina i token in una query con spazio (interpretabile come AND implicito in FTS5)
	return strings.Join(tokens, " ")
}
