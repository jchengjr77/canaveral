package main

import (
	"bufio"
	"canaveral/lib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// confirmDelete listens for user confirmation and returns a boolean
// Takes in an input channel to increase testability
// * tested
func confirmDelete(projName string, stdin io.Reader) bool {
	fmt.Printf("Are you sure you want to delete %s? ('y' or 'n')\n>", projName)
	reader := bufio.NewReader(stdin)
	response, err := reader.ReadByte()
	lib.Check(err)
	return (response == 'y')
}

// tryRemProj tries to delete a project of specified name.
// if the project exists, it will delete it and return true.
// else, it will return false or throw an error.
// * tested
func tryRemProj(projName string, wsPath string) bool {
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	files, err := ioutil.ReadDir(string(ws))
	lib.Check(err)
	for _, file := range files {
		if file.Name() == projName {
			confirm := confirmDelete(projName, os.Stdin)
			if !confirm {
				fmt.Println("Cancelling project deletion.")
				return true
			}
			err = os.RemoveAll(string(ws) + "/" + projName)
			lib.Check(err)
			fmt.Printf("Removed Project: %s\n", string(ws)+"/"+projName)
			return true
		}
	}
	return false
}

// remProjectHandler deletes a project from the canaveral workspace.
// Input: project name (string)
// Behavior: If project is found, prompt deletion of project.
//		if deletion confirmed -> delete project entirely
// 		if deletion cancelled -> exit
// 	if project is not found, exit.
// * tested
func remProjectHandler(projName string) error {
	if projName == "" {
		fmt.Println("Cannot remove an unspecified project. Please provide the project name.")
		return nil
	} else if !lib.FileExists(usrHome + confDir + wsFName) {
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs a workspace to remove projects from.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	wsPath := usrHome + confDir + wsFName
	if tryRemProj(projName, wsPath) {
		return nil
	}
	fmt.Printf("Could not find project %s in canaveral workspace.\n", projName)
	return nil
}
