package main

import (
	"canaveral/nativestore"
	"fmt"
)

var label string = "github credentials"
var url string = "https://api.github.com"

// gitCredsHandler takes in a git username and password and stores them
// ? Implement a no-password version of this perhaps?
func addGitCredsHandler(username, secret string) error {
	if username == "" {
		fmt.Println("A git username is required. Please provide one.")
		return nil
	} else if secret == "" {
		fmt.Println("A git password is required. Please provide one.")
		return nil
	} else {
		fmt.Printf("Adding git account: %s\n", username)
		return nativestore.SetCreds(label, url, username, secret)
	}
}

// Removes git credentials from native storage
func remGitCredsHandler() error {
	fmt.Println("Removing git from canaveral.")
	return nativestore.DeleteCreds(label, url)
}

// Checks whether or not the user has git credentials set
func gitCredsExist() bool {
	fmt.Println("Checking whether or not git credentials have been added.")
	_, _, err := nativestore.FetchCreds(label, url)
	return (err == nil)
}
