package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var (
	Red      = color.New(color.FgRed)
	Blue     = color.New(color.FgBlue)
	Green    = color.New(color.FgGreen)
	Magenta  = color.New(color.FgMagenta)
	White    = color.New(color.FgWhite)
	Black    = color.New(color.FgBlack)
	BRed     = color.New(color.FgRed, color.Bold)
	BBlue    = color.New(color.FgBlue, color.Bold)
	BGreen   = color.New(color.FgGreen, color.Bold)
	BMagenta = color.New(color.FgMagenta, color.Bold)
	BWhite   = color.New(color.Bold, color.FgWhite)
	BBlack   = color.New(color.Bold, color.FgBlack)
)

// Print help and exit with a status code.
func help(status int) {
	defer os.Exit(status)
	fmt.Printf(
		"%v%v%v\n",
		`Usage: lss-img [options] <patterns-to-match>
Options:
   -d "`, root, `":
        Start search under a specific directory
    -x:
        Patterns to exclude from matches, separated by commas
    -f:
        Find matches based on the full path of files
          (by default, only the basenames of files are checked for matches)
    -v:
        Display slightly more output
    -r:
        Display the relative path to <path>
          (the full path is displayed by default)
    -b
        Display the basename only
          (the full path is displayed by default)
    -c
        Colorize output
    -w:
        Write output to file
    -h:
        Print this help`,
	)
}

// Print a line break.
func lineBreak(s string) {
	line := strings.Repeat(s, 79)
	fmt.Println(line)
}

// Print a thin line break (of '-' characters).
func line(c *color.Color) {
	c.Set()
	defer colorUnset()
	lineBreak("-")
}

// Print a thick line break (of '=' characters).
func boldLine(c *color.Color) {
	c.Set()
	defer colorUnset()
	lineBreak("=")
}

// Print a found match.
func progress(output string) {
	if doColor == false {
		fmt.Println(output)
		return
	}
	defer colorUnset()
	dir, filename := filepath.Split(output)
	Blue.Printf("%v/", dir)
	Green.Println(filename)
}

// Print the number of found matches.
func report() {
	if !doVerbose {
		return
	}
	results := fmt.Sprintf("Matches found: %d", total)
	if !doColor {
		lineBreak("-")
		fmt.Println(results)
		return
	}
	defer colorUnset()
	line(BWhite)
	BGreen.Println(results)
}

// Exit if an error occurs.
func chkerr(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}

// Reset console color to normal.
func colorUnset() {
	color.Unset()
}
