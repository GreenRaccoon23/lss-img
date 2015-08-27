package main

import (
	"bytes"
	"path/filepath"
	"strings"
)

// Concatenate strings.
func concat(slc ...string) string {
	b := bytes.NewBuffer(nil)
	defer b.Reset()
	for _, s := range slc {
		b.WriteString(s)
	}
	return b.String()
}

// Concatenate elements in a slice, adding a separator between each element.
func concatBy(slc []string, separator string) string {
	var separated []string
	for _, s := range slc {
		separated = append(separated, s, separator)
	}
	return concat(separated...)
}

// Make sure a directory name includes its full path.
// Append a "/" to a string if it doesn't have one already.
func fmtDir(dir *string) {
	if *dir == "" {
		return
	}
	*dir = filepath.Clean(*dir)
	*dir, _ = filepath.Abs(*dir)
	s := string(filepath.Separator)
	if !strings.HasSuffix(*dir, s) {
		*dir = concat(*dir, s)
	}
	return
}

// Remove elements in a slice (if they exist).
// Only remove EXACT matches.
func filter(slc []string, args ...string) (filtered []string) {
	sediment := strain(slc, args...)
	for _, s := range slc {
		if !slcContains(sediment, s) {
			filtered = append(filtered, s)
		}
	}
	return
}

// Find matching elements in a slice.
func strain(slc []string, args ...string) (sediment []string) {
	for _, s := range args {
		if slcContains(slc, s) {
			sediment = append(sediment, s)
		}
	}
	return
}

// Check whether a slice contains a string.
// Only return true if an element in the slice EXACTLY matches the string.
// If testing for more than one string,
//   return true if ANY of them match an element in the slice.
func slcContains(slc []string, args ...string) bool {
	for _, s := range slc {
		for _, a := range args {
			if s == a {
				return true
			}
		}
	}
	return false
}

func isSlcInStr(s string, slc []string) bool {
	for _, a := range slc {
		if strings.Contains(s, a) {
			return true
		}
	}
	return false
}

// Check whether the string form of a byte matches
//   at least one of the letters passed.
func isByteLetter(b uint8, args ...string) bool {
	letter := string(b)
	for _, a := range args {
		if a == letter {
			return true
		}
	}
	return false
}

// Check whether the last letter in a string matches
//   at least one of the letters passed.
func isFirstLetter(s string, args ...string) bool {
	first := string(s[0])
	for _, a := range args {
		if first == a {
			return true
		}
	}
	return false
}

// Check whether the last letter in a string matches
//   at least one of the letters passed.
func isLastLetter(s string, args ...string) bool {
	lastLetter := string(s[len(s)-1])
	for _, z := range args {
		if lastLetter == z {
			return true
		}
	}
	return false
}

// Return true if at least one bool value is true.
func anyTrue(args ...bool) bool {
	for _, a := range args {
		if a {
			return true
		}
	}
	return false
}
