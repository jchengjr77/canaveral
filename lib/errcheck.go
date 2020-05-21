package lib

import (
	"fmt"
	"os"
	"os/exec"
)

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

// CreateFile creates a file with the name filename
// ? untested (low priority)
func CreateFile(filename string) error {
	f, err := os.Create(filename)
	defer f.Close()
	return err
}

// CheckToolExists uses the 'which' command to find a specific tool.
// It then parses the output of the command, and checks if
// 'which' found the toolname in the path or not.
// * tested
func CheckToolExists(toolName string) bool {
	if toolName == "" {
		fmt.Println("Cannot pass in a blank toolname")
		return false
	}
	if toolName[0] == '-' {
		fmt.Println("Cannot pass in an option as a toolname")
		return false
	}
	cmd := exec.Command("which", toolName)
	_, err := cmd.Output()
	// When 'which' cannot find the tool, it exits with status 1 (error)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("PROBLEM: %s not found in path\n", toolName)
		return false
	}
	fmt.Printf("%s found in path\n", toolName)
	return true
}
