package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// showWorkspaceHandler checks the confDir for the workspace file.
// If such a file exists, it reads its contents and navigates to that path.
// Otherwise, it notifies the user that there is no canaveral workspace set.
// ? untested, low priority
func showWorkspaceHandler() error {
	if !fileExists(usrHome + confDir + wsFName) {
		fmt.Printf("Can't find workspace file in %s\n", usrHome+confDir+wsFName)
		fmt.Println("Please specify a canaveral workspace.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	ws, err := ioutil.ReadFile(usrHome + confDir + wsFName)
	check(err)
	fmt.Printf("\nYour canaveral path: %s\n", ws)
	fmt.Printf("\nCurrent canaveral projects:\n")
	files, err := ioutil.ReadDir(string(ws))
	check(err)
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

// setWorkspaceHandler takes in a new path and writes to the confDir.
// If the workspace file already exists, it overwrites it with the new path.
// Otherwise, it creates the workspace file and writes the path in.
// ? untested, low priority
func setWorkspaceHandler(newWorkspace string) error {
	err := os.MkdirAll(usrHome+confDir, os.ModePerm)
	check(err)
	f, err := os.Create(usrHome + confDir + wsFName)
	// If file exists, truncates
	check(err)
	defer f.Close() // Close the file at the return of this function
	f.WriteString(newWorkspace)
	fmt.Printf("Set canaveral workspace to: %s\n", newWorkspace)
	return nil
}
