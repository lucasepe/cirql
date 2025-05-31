package rm

import (
	"fmt"
	"io"
)

const (
	appName = "cirql"
	cmdName = "rm"
)

func usage(wri io.Writer) {
	fmt.Fprint(wri, "\nRemove a contact.\n\n")

	fmt.Fprint(wri, "\nUSAGE:\n\n")
	fmt.Fprintf(wri, "  %s %s [FLAGS]\n\n", appName, cmdName)

	fmt.Fprint(wri, "FLAGS:\n\n")
	fmt.Fprintln(wri, "  -m, --match   <string>    Search query to filter contacts by name, email, etc.")
	fmt.Fprintln(wri)

	fmt.Fprint(wri, "EXAMPLES:\n\n")
	fmt.Fprint(wri, " » Remove a contact by filter (with confirmation):\n\n")
	fmt.Fprintf(wri, "     %s %s --match 'pinco'\n\n", appName, cmdName)
}
