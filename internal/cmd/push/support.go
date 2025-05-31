package push

import (
	"fmt"
	"io"
	"text/tabwriter"
)

type metrics struct {
	created int
	updated int
	skipped int
	failed  int
}

func printReport(wri io.Writer, m metrics) {
	total := m.created + m.updated + m.skipped + m.failed
	fmt.Fprintf(wri, "\n[%d] cards processed:\n\n", total)

	w := tabwriter.NewWriter(wri, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Action\tCount\tPercent")

	writeLine := func(label string, count int) {
		percent := 0.0
		if total > 0 {
			percent = float64(count) / float64(total) * 100
		}
		fmt.Fprintf(w, "%s\t%d\t%.1f%%\n", label, count, percent)
	}

	writeLine("Created", m.created)
	writeLine("Updated", m.updated)
	writeLine("Skipped", m.skipped)
	writeLine("Failed", m.failed)

	w.Flush()
}
