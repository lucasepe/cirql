package push

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/lucasepe/cirql/internal/store"
	getoptutil "github.com/lucasepe/cirql/internal/util/getopt"
	ioutil "github.com/lucasepe/cirql/internal/util/io"
	"github.com/lucasepe/cirql/internal/vcards"
	"github.com/lucasepe/x/getopt"
)

func Do(args []string) error {
	extras, opts, err := getopt.GetOpt(args,
		"ho",
		[]string{"help", "override"},
	)
	if err != nil {
		return err
	}

	showHelp := getoptutil.HasOpt(opts, []string{"-h", "--help"})
	if showHelp {
		usage(os.Stderr)
		return nil
	}

	override := getoptutil.HasOpt(opts, []string{"-o"})

	var filename string
	if len(extras) > 0 {
		filename = extras[0]
	}

	src, cleanup, err := ioutil.FileOrStdin(filename)
	if err != nil {
		if errors.Is(err, ioutil.ErrNoInputDetected) {
			fmt.Fprintf(os.Stderr, "warning: %s.\n", err.Error())
			usage(os.Stderr)
			return nil
		}
		return err
	}
	defer cleanup()

	db, err := store.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	stats, err := doPush(src, db, override)
	if err != nil {
		return err
	}

	printReport(os.Stdout, stats)
	return nil
}

func doPush(src io.Reader, dst *sql.DB, override bool) (metrics, error) {
	stats := metrics{}

	dec := vcards.NewDecoder(src)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			stats.failed += 1
			return stats, err
		}

		res, err := store.CreateOrUpdate(dst, card, override)
		if err != nil {
			stats.failed += 1
			return stats, err
		}

		switch res {
		case store.Created:
			stats.created += 1
		case store.Skipped:
			stats.skipped += 1
		case store.Updated:
			stats.updated += 1
		}
	}

	return stats, nil
}
