package pull

import (
	"fmt"
	"io"
)

const (
	appName = "cirql"
	cmdName = "pull"
)

func usage(wri io.Writer) {
	fmt.Fprint(wri, "\nSearch for contacts and display them in vCard format.\n\n")

	fmt.Fprint(wri, "\nUSAGE:\n\n")
	fmt.Fprintf(wri, "  %s %s [FLAGS]\n\n", appName, cmdName)

	fmt.Fprint(wri, "FLAGS:\n\n")
	fmt.Fprint(wri, "  -m, --match    <string>    Search query to filter contacts by name, email, etc.\n\n")

	fmt.Fprint(wri, "  -g, --groups   <list>      Filter contacts by category (accepts a comma-separated list).\n")
	fmt.Fprint(wri, "                               * example: \"family,friends,work\"\n")
	fmt.Fprint(wri, "                               * matches contacts having at least one of the\n")
	fmt.Fprint(wri, "                                 specified categories\n\n")

	fmt.Fprint(wri, "  -d, --days     <number>    Only include contacts with a birthday in the next N days.\n\n")
	fmt.Fprint(wri, "  -s, --split    <number>    Write every N contacts to a separate file instead of stdout.\n")
	fmt.Fprint(wri, "                               * only applies when using --template flag\n")
	fmt.Fprint(wri, "                               * files are saved in the --output directory using\n")
	fmt.Fprint(wri, "                                 the template filename as base\n\n")
	fmt.Fprint(wri, "  -o, --output    <dir>      Directory where generated custom message files will be saved.\n")
	fmt.Fprint(wri, "                               * only used with --split\n")
	fmt.Fprint(wri, "                               * defaults to current directory\n\n")
	fmt.Fprint(wri, "  -t, --template  <file>     Template file for the message (provide file path).\n\n")
	fmt.Fprint(wri, "  -h, --help                 Display this help and exit.\n\n")
	fmt.Fprintln(wri)

	fmt.Fprint(wri, "EXAMPLES:\n\n")
	fmt.Fprint(wri, " » Search for contacts with a specific surname:\n\n")
	fmt.Fprintf(wri, "     %s %s -m 'rossi'\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Search for contacts by email domain:\n\n")
	fmt.Fprintf(wri, "     %s %s -m '@example.com'\n\n", appName, cmdName)
	fmt.Fprint(wri, " » List contacts with upcoming birthdays in the next 7 days:\n\n")
	fmt.Fprintf(wri, "     %s %s --days 7\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Search contacts named 'paolo' with a birthday in the next 15 days:\n\n")
	fmt.Fprintf(wri, "     %s %s -m \"paolo\" --days 15\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Filter contacts in the 'family' or 'work' categories:\n\n")
	fmt.Fprintf(wri, "     %s %s --groups \"family,work\"\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Find contacts in 'friends' category with birthdays in the next 10 days:\n\n")
	fmt.Fprintf(wri, "     %s %s -g friends -d 10\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Generate messages using a custom template file:\n\n")
	fmt.Fprintf(wri, "     %s %s --template ./templates/birthday_message.txt\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Generate a message for contacts named 'claudio' using a custom template:\n\n")
	fmt.Fprintf(wri, "     %s %s -m \"claudio\" -t ./templates/draft.txt\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Generate a separate file for every 10 contacts:\n\n")
	fmt.Fprintf(wri, "     %s %s --split 10\n\n", appName, cmdName)
	fmt.Fprint(wri, " » Generate a separate file every 15 contacts using a custom template:\n\n")
	fmt.Fprintf(wri, "     %s %s -t ./templates/discount.txt -s 15\n\n", appName, cmdName)
	fmt.Fprint(wri, "   This will produce (in the current directory) files like:\n\n")
	fmt.Fprint(wri, "      discount_001.txt, discount_002.txt, etc.\n")
}
