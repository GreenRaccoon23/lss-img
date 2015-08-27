package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Helper function for filepath.Walk().
// Evaluate matches based on the basename of the file.
func walkFnName(path string, file os.FileInfo, err error) error {
	// Do a workaround for filepath package bug.
	if _, err = os.Stat(path); err != nil {
		return nil
	}

	// Quit if the file isn't an image.
	if !isImg(path) {
		return nil
	}

	// IF the current file is a directory,
	//     THEN skip it.
	if file.IsDir() {
		return nil
	}

	// Evaluate whether the filename is one which the user
	//   wants to be printed to the console.
	target := file.Name()
	if isReject(target) {
		return nil
	}

	// Print the filename to the console.
	genOutput(path)
	return nil
}

// Helper function for filepath.Walk().
// Evaluate matches based on the full path of the file.
func walkFnPath(path string, file os.FileInfo, err error) error {
	// Do a workaround for filepath package bug.
	if _, err = os.Stat(path); err != nil {
		return nil
	}

	// Quit if the file isn't an image.
	if !isImg(path) {
		return nil
	}

	// Evaluate whether the filename is one which the user
	//   wants to be printed to the console.
	target := path
	if isReject(target) {
		return nil
	}

	// Print the path to the console.
	genOutput(path)
	return nil
}

// Evaluate whether a file is a match
//   based on user specified patterns and exclusions.
func isReject(name string) bool {
	// Quit if the name doesn't matche all of the user-specified patterns.
	if !isMatch(name) {
		return true
	}

	// Quit early if user didn't specify any exclusions.
	if !doExclude {
		return false
	}

	// Evaluate whether the name matches any of the user-specified exclusions.
	if isExclusion(name) {
		return true
	}
	return false
}

// Evaluate whether a file is a match
//   based on user specified patterns.
// The file must match ALL patterns.
func isMatch(name string) bool {
	for _, p := range patterns {
		if !strings.Contains(name, p) {
			return false
		}
	}
	return true
}

// Evaluate whether a file is a match
//   based on user specified exclusions.
// The file must NOT match ANY exclusion patterns.
func isExclusion(s string) bool {
	if isSlcInStr(s, exclusions) {
		return true
	}
	return false
}

// Output the filename or path to the console.
// Write it to the log file if user requested it.
func genOutput(path string) {
	output := fmtOutput(path)
	defer progress(output)
	defer func() { total++ }()

	if !doWrite {
		return
	}
	write(output, logFile)
}

// Format the filename to be printed based on user specifications.
func fmtOutput(path string) string {
	if doRelative {
		return relPath(root, path)
	}
	if doBase {
		return filepath.Base(path)
	}
	return path
}
