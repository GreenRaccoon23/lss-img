package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	logFile *os.File
	total   int
)

func init() {
	chkHelp()
	flags()
	flagsEval()
}

func main() {
	if doColor {
		defer colorUnset()
	}
	if doWrite {
		genLog()
		defer logFile.Close()
	}

	err := search()
	chkerr(err)
	report()
}

// Make a file in which to record matches.
func genLog() {
	pattern := concatBy(patterns, ",")
	fileName := concat("lss-img_", pattern)

	remove(fileName)
	logFile = create(fileName)

	header := fmt.Sprintf(
		"Search results for:\n    \"%v\"\nunder:\n    \"%v\"\n",
		pattern,
		root,
	)
	write(header, logFile)
}

// Search the filesystem.
func search() (err error) {
	if doMatchPath {
		err = filepath.Walk(root, walkFnPath)
	} else {
		err = filepath.Walk(root, walkFnName)
	}
	return
}
