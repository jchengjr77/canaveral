package main

import (
	"canaveral/lib"
	"canaveral/node"
	"canaveral/react"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// initRepo initializes a git repo if the -g flag is set
// and git credentials are stored in the users native storage
// ? untested
func initRepo(projName string) {
	if projName == "" {
		fmt.Println("Please provide a repo name.")
		fmt.Println("(For more info, 'canaveral --help')")
	}
	ws, err := ioutil.ReadFile(usrHome + confDir + wsFName)
	lib.Check(err)
	os.Chdir(string(ws))
	createRepo := exec.Command("git", "init", projName)
	createRepo.Stdout = os.Stdout
	createRepo.Stdin = os.Stdin
	err = createRepo.Run()
	lib.Check(err)
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
		return nil
	} else if !lib.FileExists(usrHome + confDir + wsFName) {
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs a workspace to add projects to.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	if projType == "react" {
		fmt.Println("Creating React project...")
		react.AddReactProj(projName, usrHome+confDir+wsFName)
		if init {
			initRepo(projName)
		}
	} else if projType == "node" {
		fmt.Println("Creating Node project...")
		node.AddNodeProj(projName, usrHome+confDir+wsFName)
		if init {
			initRepo(projName)
		}
	} else {
		fmt.Println("Creating generic project...")
		addProj(projName, usrHome+confDir+wsFName)
		if init {
			initRepo(projName)
		}
	}
	return nil
}
