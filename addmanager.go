package main

import (
	"canaveral/git"
	"canaveral/lib"
	"canaveral/node"
	"canaveral/python"
	"canaveral/react"
	"canaveral/reactnative"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Wraps the init repo to perform init in correct directory
// * wraps tested function
func createAndInit(projName string) error {
	ws, err := ioutil.ReadFile(usrHome + confDir + wsFName)
	lib.Check(err)
	os.Chdir(string(ws) + "/" + projName)
	return git.InitRepo("", "")
}

// addProj takes in a project name and adds it to the workspace.
// Requires that the workspace exists.
// Requires that the project name is non-empty
// * tested
func addProj(projName string, wsPath string) {
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	err = os.MkdirAll(string(ws)+"/"+projName, os.ModePerm)
	lib.Check(err)
	fmt.Printf("Added project %s to workspace %s\n", projName, string(ws))
}

// ? Incomplete functionality
// addProjectHandler takes in a project name and initializes a new project.
// If the project name is empty, it prompts the user to enter a name.
// Vanilla behavior includes generating a directory labeled the project name.
// Initializes all boilerplate code for specified project type.
// * tested
func addProjectHandler(projName string, projType string, init bool) error {
	if projName == "" {
		fmt.Println("Please provide a project name.")
		fmt.Println("(For more info, 'canaveral --help')")
		return errors.New("No project name specified")
	} else if !lib.FileExists(usrHome + confDir + wsFName) {
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs a workspace to add projects to.")
		fmt.Println("(For help, type 'canaveral --help')")
		return errors.New("Unspecified workspace")
	}
	if projType == "react" {
		fmt.Println("Creating React project...")
		react.AddReactProj(projName, usrHome+confDir+wsFName)
	} else if projType == "node" {
		fmt.Println("Creating Node project...")
		node.AddNodeProj(projName, usrHome+confDir+wsFName)
	} else if projType == "python" {
		fmt.Println("Creating Python project...")
		python.AddPythonProj(projName, usrHome+confDir+wsFName)
	} else if projType == "reactnative" {
		fmt.Println("Creating React Native project...")
		reactnative.AddReactNativeProj(projName, usrHome+confDir+wsFName)
	} else if projType == "" {
		fmt.Println("Creating generic project...")
		addProj(projName, usrHome+confDir+wsFName)
	} else {
		fmt.Printf("Sorry, we do not support the project type '%s'\n", projType)
		return errors.New("Unsupported project type specified")
	}
	if init {
		createAndInit(projName)
	}
	return nil
}
