package rm

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/lucasepe/cirql/internal/store"
	getoptutil "github.com/lucasepe/cirql/internal/util/getopt"
	"github.com/lucasepe/x/getopt"
)

func Do(args []string) error {
	_, opts, err := getopt.GetOpt(
		args,
		"m:h",
		[]string{
			"match",
			"help",
		},
	)
	if err != nil {
		return err
	}

	match := getoptutil.FindOptVal(opts, []string{"-m", "--match"})

	showHelp := (match == "")
	showHelp = showHelp || getoptutil.HasOpt(opts, []string{"-h", "--help"})
	if showHelp {
		usage(os.Stderr)
		return nil
	}

	db, err := store.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	id, err := chooseContactID(db, "Select ID to delete", match)
	if err != nil {
		return err
	}
	if id <= 0 {
		return nil
	}

	return deleteContactByID(db, id, true)
}

func deleteContactByID(db *sql.DB, id int64, prompt bool) error {
	if prompt {
		fmt.Fprintf(os.Stderr, " » Delete contact with ID '%d'? [y/N]: ", id)

		var confirm string
		_, err := fmt.Scanln(&confirm)
		if err != nil {
			return err
		}

		if strings.ToLower(confirm) != "y" {
			fmt.Fprintln(os.Stderr, " » Aborted.")
			return nil
		}
	}

	if err := store.Delete(db, id); err != nil {
		return err
	}

	fmt.Fprint(os.Stdout, "\nContact successfully deleted.\n")

	return nil
}
