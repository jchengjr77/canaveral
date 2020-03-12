package main

import (
	"os"
	"testing"
)

func TestAddGitCredsHandler(t *testing.T) {
	os.Setenv("CredentialsTest", "true")
	// Git credentials shouldn't exist
	if gitCredsExist() {
		t.Error("Git credentials exist on entry into test (bad state)")
	}

	// Testing failure with no username
	resNoUsr := captureOutput(
		func() {
			addGitCredsHandler("", "password")
		})
	if resNoUsr != "A git username is required. Please provide one.\n" {
		t.Logf("addGitCredsHandler('', _) output: %s\n", resNoUsr)
		t.Error("func addGitCredsHandler() failed in case of ('', _)\n")
	}

	// Testing failure with no password
	resNoPass := captureOutput(
		func() {
			addGitCredsHandler("username", "")
		})
	if resNoPass != "A git personal auth token is required. Please provide one.\n" {
		t.Logf("addGitCredsHandler('', _) output: %s\n", resNoPass)
		t.Error("func addGitCredsHandler() failed in case of (_, '')\n")
	}
}
