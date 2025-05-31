package push

import (
	"os"
)

func Example_printReport() {
	m := metrics{
		created: 10,
		updated: 20,
		skipped: 8,
		failed:  4,
	}

	printReport(os.Stdout, m)

	// Output:
	//
	// [42] cards processed:
	//
	// Action   Count  Percent
	// Created  10     23.8%
	// Updated  20     47.6%
	// Skipped  8      19.0%
	// Failed   4      9.5%
}
