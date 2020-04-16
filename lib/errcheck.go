package lib

import "os"

// Check takes in an error and verifies if it is nil.
// If the error is not nil, it will terminate the program.
// * tested
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
// * tested
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a dir exists and is not a file before we
// try using it to prevent further errors.
// * tested
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
