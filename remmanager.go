package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// confirmDelete listens for user confirmation and returns a boolean
// ! untested
func confirmDelete(projName string) bool {
	fmt.Printf("Are you sure you want to delete %s? ('y' or 'n')\n>", projName)
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadByte()
	check(err)
	return (response == 'y')
}

// tryRemProj tries to delete a project of specified name.
// if the project exists, it will delete it and return true.
// else, it will return false or throw an error.
// ! untested
func tryRemProj(projName string) bool {
	ws, err := ioutil.ReadFile(usrHome + confDir + wsFName)
	check(err)
	files, err := ioutil.ReadDir(string(ws))
	check(err)
	for _, file := range files {
		if file.Name() == projName {
			confirm := confirmDelete(projName)
			if !confirm {
				fmt.Println("Cancelling project deletion.")
				return true
			}
			err = os.RemoveAll(string(ws) + "/" + projName)
			check(err)
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
// ! untested
func remProjectHandler(projName string) error {
	if projName == "" {
		fmt.Println("Cannot remove an unspecified project. Please provide the project name.")
		return nil
	} else if !fileExists(usrHome + confDir + wsFName) {
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs a workspace to remove projects from.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	if tryRemProj(projName) {
		return nil
	}
	fmt.Printf("Could not find project %s in canaveral workspace.\n", projName)
	return nil
}
