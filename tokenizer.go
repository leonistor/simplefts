package main

import (
	"strings"
	"unicode"
)

// tokenize returns a slice of tokens for the given text
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// split on any character that is not a letter or a number
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
