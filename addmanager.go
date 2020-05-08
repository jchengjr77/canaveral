package main

import (
	"canaveral/react"
	"fmt"
	"io/ioutil"
	"os"
)

// addProj takes in a project name and adds it to the workspace.
// Requires that the workspace exists.
// Requires that the project name is non-empty
// * tested
func addProj(projName string, wsPath string) {
	ws, err := ioutil.ReadFile(wsPath)
	check(err)
	err = os.MkdirAll(string(ws)+"/"+projName, os.ModePerm)
	check(err)
	fmt.Printf("Added project %s to workspace %s\n", projName, string(ws))
}

// ? Incomplete functionality
// addProjectHandler takes in a project name and initializes a new project.
// If the project name is empty, it prompts the user to enter a name.
// Vanilla behavior includes generating a directory labeled the project name.
// Initializes all boilerplate code for specified project type.
// * tested
func addProjectHandler(projName string, projType string) error {
	if projName == "" {
		fmt.Println("Please provide a project name.")
		fmt.Println("(For more info, 'canaveral --help')")
		return nil
	} else if !fileExists(usrHome + confDir + wsFName) {
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs a workspace to add projects to.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	if projType == "react" {
		fmt.Println("Creating React project...")
		react.AddReactProj(projName, usrHome+confDir+wsFName)
	} else {
		fmt.Println("Creating generic project...")
		addProj(projName, usrHome+confDir+wsFName)
	}
	return nil
}
