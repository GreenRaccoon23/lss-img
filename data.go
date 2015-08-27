package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Get the present working directory.
// Exit if an error occurs.
func pwd() string {
	pwd, err := os.Getwd()
	chkerr(err)
	return pwd
}

// Check whether a file is an image.
func isImg(path string) bool {
	ext := filepath.Ext(path)
	if len(ext) < 2 {
		return false
	}
	// Also check for compressed images.
	compressions := []string{".gz", ".bz2"}
	for _, c := range compressions {
		if strings.ToLower(ext) == c {
			ext = concat(ext, c[1:])
		}
	}

	// List of image file types (non-exhaustive, obviously)
	types := []string{
		"svg",
		"jpg", "jpeg",
		"png",
		"gif",
		"tiff", "tif",
		"rif",
		"pdf",
		"eps", "ai",
		"psd",
		"cpt", "psp",
		"xpm", "xbm", "xmc", "bitmap", "xwd",
		"bpg",
		"webp",
		"hdr",
		"xcf", "xcfbz2", "xcfgz", "gbr", "gih", "pat",
		"pix", "matte", "mask", "alpha", "als",
		"fli", "flc",
		"dcm", "dicom",
		"fit", "fits",
		"cel",
		"ico",
		"mng",
		"ora",
		"ppm", "pgm", "pbm", "pnm",
		"ps",
		"sgi", "rgb", "rgba", "bw", "icon",
		"im1", "im8", "im24", "im32", "rs", "ras",
		"tga",
		"pcx", "pcc",
		"xps",
		"odg",
	}

	// First check whether the file extension matches
	//   one in the list of image file extensions.
	for _, t := range types {
		if strings.ToLower(ext[1:]) == t {
			return true
		}
	}

	// If no match was found, check for a file signature as a last resort.
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return false
	}

	if imgType(file) != "" {
		return true
	}

	return false
}

// Find the image type of a file based on file signatures.
// Only works for png, jpg, gif, and bmp images.
func imgType(file *os.File) string {
	bytes := make([]byte, 4)
	n, _ := file.ReadAt(bytes, 0)
	if n < 4 {
		return ""
	}
	if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		return "png"
	}
	if bytes[0] == 0xFF && bytes[1] == 0xD8 {
		return "jpg"
	}
	if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 && bytes[3] == 0x38 {
		return "gif"
	}
	if bytes[0] == 0x42 && bytes[1] == 0x4D {
		return "bmp"
	}
	return ""
}

// Basically filepath.Rel()
//   but quicker/less sophisticated and without the error.
func relPath(relDir, fullPath string) string {
	return strings.Replace(fullPath, relDir, "", 1)
}

// Write a string to a file.
func write(s string, file *os.File) {
	t := concat(s, "\n")
	b := []byte(t)
	_, err := file.Write(b)
	chkerr(err)
}

// Remove a file if it exists.
func remove(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}
}

// Create a file.
// Exit if an error occurs.
func create(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	chkerr(err)
	return file
}
