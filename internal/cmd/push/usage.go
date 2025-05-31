package push

import (
	"fmt"
	"io"
)

const (
	appName = "cirql"
	cmdName = "push"
)

func usage(wri io.Writer) {
	fmt.Fprint(wri, "\nImport contacts from vCard(s).\n\n")

	fmt.Fprint(wri, "\nUSAGE:\n\n")
	fmt.Fprintf(wri, "  %s %s [FLAGS]\n\n", appName, cmdName)

	fmt.Fprint(wri, "FLAGS:\n\n")
	fmt.Fprintln(wri, "  -o, --override      Update existing entries")
	fmt.Fprintln(wri, "  -h, --help          Display this help and exit.")
	fmt.Fprintln(wri)

	fmt.Fprint(wri, "EXAMPLES:\n\n")
	fmt.Fprint(wri, " » Import contacts from a vCard file (skip existing):\n\n")
	fmt.Fprintf(wri, "     %s %s /path/to/my-contacts.vcf\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Import contacts from a vCard file (override existing):\n\n")
	fmt.Fprintf(wri, "     %s %s -o /path/to/my-contacts.vcf\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Import contacts providing input via pipe:\n\n")
	fmt.Fprintf(wri, "     cat /path/to/my-contacts.vcf | %s %s\n\n", appName, cmdName)
}
