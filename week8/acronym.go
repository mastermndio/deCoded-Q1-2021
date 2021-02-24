// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re := regexp.MustCompile("_| |-")
	splitS := re.Split(s, -1)
	acronym := []string{}

	for _, word := range splitS {
		if len(word) > 0 {
			acronym = append(acronym, strings.ToUpper(string(word[0])))
		}
	}

	return strings.Join(acronym, "")
}
