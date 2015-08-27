package main

import (
	"os"
	"strings"
)

var (
	//Flags with '-' switches
	root        string = pwd()
	exclude     string
	doMatchPath bool
	doVerbose   bool
	doColor     bool
	doRelative  bool
	doBase      bool
	doWrite     bool

	//Flags without '-' switches
	patterns   []string
	exclusions []string

	doExclude bool
)

// Check if user requested help.
func chkHelp() {
	if len(os.Args) < 2 {
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "-h", "--h", "--help", "-help":
		help(0)
	}

}

// Parse user flags.
func flags() {
	sFlags := map[string]*string{
		"d": &root,
		"x": &exclude,
	}
	bFlags := map[string]*bool{
		"f": &doMatchPath,
		"v": &doVerbose,
		"r": &doRelative,
		"b": &doBase,
		"c": &doColor,
		"w": &doWrite,
	}

	for i, f := range os.Args {
		if len(f) == 0 {
			continue
		}
		if isByteLetter(f[0], "-") == false {
			continue
		}

		for _, r := range f[1:] {
			s := string(r)
			boolParse(bFlags, s)
			strParse(sFlags, i, s)
		}
	}

	args := flagFilter(sFlags)
	patterns = args
}

// Further modify global variables based on user flags.
func flagsEval() {
	fmtDir(&root)
	// The user can separate patterns by commas instead of spaces.
	if len(patterns) == 1 {
		patterns = strings.Split(patterns[0], ",")
	}

	if exclude != "" {
		doExclude = true
		exclusions = strings.Split(exclude, ",")
	}
}

// Check for bool flags ('-b').
func boolParse(m map[string]*bool, f string) {
	for s, b := range m {
		if s != f {
			continue
		}
		*b = true
	}
}

// Check for string flags ('-s "string"').
func strParse(m map[string]*string, i int, f string) {
	for s, t := range m {
		if s != f {
			continue
		}
		*t = argNext(i)
	}
}

// Return the commandline line argument after the number specified.
func argNext(i int) string {
	if len(os.Args) <= i {
		help(1)
	}
	return os.Args[i+1]
}

// Filter out the flagged arguments from the commandline parameters.
// Return any leftover arguments.
func flagFilter(m map[string]*string) (filtered []string) {
	if len(os.Args) < 2 {
		return
	}

	strFlags := strFlags(m)

	for _, a := range os.Args[1:] {
		if isFirstLetter(a, "-") {
			continue
		}
		if slcContains(strFlags, a) {
			continue
		}
		filtered = append(filtered, a)
	}
	return
}

// Make a list of all string flags, not including their switches ('-s').
func strFlags(m map[string]*string) (slc []string) {
	for _, v := range m {
		slc = append(slc, *v)
	}
	return
}
