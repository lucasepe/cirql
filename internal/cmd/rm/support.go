package rm

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/lucasepe/cirql/internal/store"
	"github.com/lucasepe/cirql/internal/vcards"
)

func chooseContactID(db *sql.DB, title, query string) (int64, error) {
	fmt.Fprintln(os.Stderr)

	tot, err := store.List(db, store.ListOptions{
		Match:   query,
		Handler: &printCardHandler{},
	})
	if err != nil {
		return -1, err
	}

	if tot == 0 {
		fmt.Fprintln(os.Stderr, "No contacts found.")
		return -1, nil
	}

	fmt.Fprintf(os.Stderr, "\n » %s (or press ENTER to cancel): ", title)

	var input string
	fmt.Scanln(&input)

	if input == "" || input == "0" {
		fmt.Fprintln(os.Stderr, " » Cancelled.")
		return -1, nil
	}

	return strconv.ParseInt(input, 10, 64)
}

var _ vcards.CardHandler = (*printCardHandler)(nil)

type printCardHandler struct{}

func (h *printCardHandler) Handle(c vcards.Card) error {
	id, err := store.ParseUID(vcards.UID(c))
	if err != nil {
		return err
	}

	if c.Name() == nil {
		return nil
	}

	fn := fmt.Sprintf("%s %s", c.Name().GivenName, c.Name().FamilyName)
	fmt.Fprintf(os.Stderr, "[%d] %s\n", id, fn)

	return nil
}
