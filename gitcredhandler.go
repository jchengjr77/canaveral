package main

import (
	"canaveral/nativestore"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// gitAddWrapper wraps the addGitCredsHandler function, taking in a username
// and securely reading the password
// * tested
func gitAddWrapper() error {
	fmt.Print("Enter username: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		password := string(bytePassword)
		fmt.Print("\r\n")
		return addGitCredsHandler(username, password)
	}
	return err
}

// gitCredsHandler takes in a git username and password and stores them
// ? Implement a no-password version of this perhaps?
// * tested
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
// * tested
func remGitCredsHandler() error {
	fmt.Println("Removing git from canaveral.")
	return nativestore.DeleteCreds(label, url)
}

// Checks whether or not the user has git credentials set
// * tested
func gitCredsExist() bool {
	fmt.Println("Checking whether or not git credentials have been added.")
	_, _, err := nativestore.FetchCreds(label, url)
	return (err == nil)
}
