package main

import "os"

// check takes in an error and verifies if it is nil.
// If the error is not nil, it will terminate the program.
// * tested
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
// * tested
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// dirExists checks if a dir exists and is not a file before we
// try using it to prevent further errors.
// * tested
func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
