package cmd

import (
	"fmt"
	"io"
	"strings"

	xtext "github.com/lucasepe/x/text"
)

const (
	appName = "cirql"
)

func usage(wri io.Writer) {
	var (
		desc = []string{
			"A simple, privacy-first command-line tool for managing contacts locally.\n",
			"This CLI tool allows you to manage contact information entirely offline using a lightweight SQLite database.",
			"» Contacts can belong to one or more categories, making organization flexible and efficient.",
			"» Full-text search (FTS) is supported for fast, powerful lookups across contact data.",
			"» Import contacts from vCard files, and export one, many, or all contacts as vCards.",
			"» Search for contacts with upcoming birthdays within a specified days range.",
			"Ideal for users who value full control over their data while retaining the freedom to choose whether or not to use cloud services.",
		}

		donateInfo = []string{
			"If you find this tool helpful consider supporting with a donation.",
			"Every bit helps cover development time and fuels future improvements.\n",
			"Your support truly makes a difference — thank you!\n",
			"  * https://www.paypal.com/donate/?hosted_button_id=FV575PVWGXZBY\n",
		}
	)

	fmt.Fprintln(wri)
	fmt.Fprint(wri, "┌─┐┬┬─┐┌─┐ ┬\n")
	fmt.Fprint(wri, "│  │├┬┘│─┼┐│  \n")
	fmt.Fprint(wri, "└─┘┴┴└─└─┘└┴─┘\n")

	fmt.Fprintln(wri)
	for _, el := range desc {
		if el[0] == 194 {
			fmt.Fprintf(wri, "%s\n\n", xtext.Indent(xtext.Wrap(el, 60), "  "))
			continue
		}
		fmt.Fprintf(wri, "%s\n\n", xtext.Wrap(el, 76))
	}
	fmt.Fprintln(wri)

	fmt.Fprint(wri, "USAGE:\n\n")
	fmt.Fprintf(wri, "  %s <command> [FLAGS]\n\n", appName)

	fmt.Fprint(wri, "COMMANDS:\n\n")
	fmt.Fprint(wri, "  rm           Remove a contact.\n")
	fmt.Fprint(wri, "  push         Import contacts from vCard(s).\n")
	fmt.Fprint(wri, "  pull         Show contacts as vCard or generate messages from a template.\n")
	fmt.Fprint(wri, "  help         Display this help and exit.\n")
	fmt.Fprint(wri, "  version      Output version information and exit.\n\n")

	fmt.Fprintf(wri, "» Type '%s <command> --help' for help about a specific command.\n", appName)

	fmt.Fprint(wri, "\n\nSUPPORT:\n\n")
	fmt.Fprint(wri, xtext.Indent(strings.Join(donateInfo, "\n"), "  "))
	fmt.Fprint(wri, "\n\n")

	fmt.Fprintln(wri, "Copyright (c) 2025 Luca Sepe")
}
