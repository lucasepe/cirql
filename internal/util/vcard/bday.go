package vcard

import (
	"time"
)

func IsMalformedBDAY(dateStr string) bool {
	const layout = "20060102"
	_, err := time.Parse(layout, dateStr)
	if err != nil {
		return true
	}
	return false
}
