package pull

import (
	"fmt"
	"os"
	"strings"

	"github.com/lucasepe/cirql/internal/store"
	getoptutil "github.com/lucasepe/cirql/internal/util/getopt"
	"github.com/lucasepe/cirql/internal/vcards"
	"github.com/lucasepe/x/getopt"
	"github.com/lucasepe/x/text/conv"
)

func Do(args []string) error {
	_, opts, err := getopt.GetOpt(
		args,
		"m:d:g:o:s:t:h",
		[]string{
			"match",
			"days",
			"groups",
			"output",
			"split",
			"template",
			"help",
		},
	)
	if err != nil {
		return err
	}

	match := getoptutil.FindOptVal(opts, []string{"-m", "--match"})
	daysUntilBirth := conv.Int(getoptutil.FindOptVal(opts, []string{"-d", "--days"}), 0)
	split := conv.Int(getoptutil.FindOptVal(opts, []string{"-s", "--split"}), 0)
	output := getoptutil.FindOptVal(opts, []string{"-o", "--output"})
	templateFile := getoptutil.FindOptVal(opts, []string{"-t", "--template"})
	categories := getoptutil.FindOptVal(opts, []string{"-g", "--groups"})

	showHelp := getoptutil.HasOpt(opts, []string{"-h", "--help"})
	if showHelp {
		usage(os.Stderr)
		return nil
	}

	var groups []string
	if categories != "" {
		groups = strings.Split(categories, ",")
	}

	if templateFile == "" {
		if split > 0 {
			fmt.Fprintln(os.Stderr, "Warning: --split has no effect without --template")
		}
		if output != "" {
			fmt.Fprintln(os.Stderr, "Warning: --output has no effect without --template")
		}
	} else {
		if output == "" {
			output, err = os.Getwd()
			if err != nil {
				return err
			}
		}
	}

	db, err := store.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	var (
		handler vcards.CardHandler
		closer  func() error
	)

	if templateFile == "" {
		handler = newDefaultHandler(os.Stdout)
	} else {
		wri, err := newFileRotator(templateFile, output)
		if err != nil {
			return err
		}
		closer = wri.Close

		enc, err := newTemplateEncoder(templateFile)
		if err != nil {
			return err
		}

		handler = newTemplateHandler(wri, enc, split)
	}

	defer func() {
		if closer != nil {
			err := closer()
			if err != nil {
				fmt.Fprintf(os.Stderr, "\n%s\n", err.Error())
				os.Exit(1)
			}
		}
	}()

	tot, err := store.List(db, store.ListOptions{
		Match:          match,
		Categories:     groups,
		DaysUntilBirth: daysUntilBirth,
		Handler:        handler,
	})
	if err != nil {
		return err
	}

	if tot == 0 {
		fmt.Fprintf(os.Stderr, "\nNo results.\n")
		return nil
	}

	fmt.Fprintf(os.Stderr, "\nFound %d total contacts.\n", tot)

	return nil
}
