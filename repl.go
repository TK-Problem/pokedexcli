package main

import "strings"

// cleanInput lowercases the text and splits it into words on any run of
// whitespace. Leading and trailing whitespace is discarded.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
