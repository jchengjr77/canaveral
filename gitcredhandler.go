package main

import (
	"canaveral/nativestore"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type goodResponse struct {
	Login string
}

// Checks that user auth token corresponds to a valid auth token and matches
// the username passed in
func verifyCreds(usr, secret string) error {
	var goodFill []goodResponse

	usrURL := url + "/user"
	request, reqErr := http.NewRequest("GET", usrURL, nil)
	if reqErr != nil {
		return reqErr
	}
	request.Header.Set("Authorization", "token "+secret)

	response, respErr := http.DefaultClient.Do(request)
	if respErr != nil {
		return respErr
	}
	defer response.Body.Close()

	responseJSONData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return readErr
	}

	start := make([]byte, 1)
	last := make([]byte, 1)
	start[0] = '['
	last[0] = ']'
	toDecode := append(start, append(responseJSONData, last...)...)

	err := json.Unmarshal(toDecode, &goodFill)
	if goodFill[0].Login == "" {
		return errors.New("Failed to authenticate token")
	} else if err == nil {
		if strings.ToLower(usr) == strings.ToLower(goodFill[0].Login) {
			return nil
		}
		return errors.New("Token didn't correspond to username")
	}
	return err
}

// gitAddWrapper wraps the addGitCredsHandler function, taking in a username
// and securely reading the personal auth token
func gitAddWrapper() error {
	fmt.Print("Enter username: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Enter Personal Auth Token: ")
	byteToken, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		token := string(byteToken)
		fmt.Print("\r\n")
		verifyErr := verifyCreds(username, token)
		if verifyErr == nil {
			return addGitCredsHandler(username, token)
		}
		// fmt.Println(verifyErr)
		return verifyErr

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
		fmt.Println("A git personal auth token is required. Please provide one.")
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

func printGitUser() {
	if gitCredsExist() {
		usr, _, err := nativestore.FetchCreds(label, url)
		check(err)
		fmt.Println(usr)
		return
	}
	fmt.Println("-no git username stored-")
}
