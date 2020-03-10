package main

import "fmt"

// ? Incomplete functionality
// addProjectHandler takes in a project name and initializes a new project.
// If the project name is empty, it prompts the user to enter a name.
// Vanilla behavior includes generating a directory labeled the project name.
// Initializes all boilerplate code for specified project type.
// ! untested
func addProjectHandler(projName string) error {
	if projName == "" {
		fmt.Println("Please provide a project name.")
		fmt.Println("(For more info, 'canaveral --help')")
	} else {
		fmt.Printf("Adding Project: %s\n", projName)
	}
	return nil
}
