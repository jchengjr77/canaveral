package github

import (
	"os"
	"testing"

	"github.com/jchengjr77/canaveral/lib"
)

func TestAddGHCredsHandler(t *testing.T) {
	err := os.Setenv("CredentialsTest", "true")
	lib.Check(err)
	// Github credentials shouldn't exist
	if ghCredsExist() {
		t.Error("Github credentials exist on entry into test (bad state)")
	}

	// Testing failure with no username
	resNoUsr := lib.CaptureOutput(
		func() {
			addGHCredsHandler("", "password")
		})
	if resNoUsr != "A github username is required. Please provide one.\n" {
		t.Logf("addGHCredsHandler('', _) output: %s\n", resNoUsr)
		t.Error("func addGHCredsHandler() failed in case of ('', _)\n")
	}

	// Testing failure with no password
	resNoPass := lib.CaptureOutput(
		func() {
			addGHCredsHandler("username", "")
		})
	if resNoPass != "A github personal auth token is required. Please provide one.\n" {
		t.Logf("addGHCredsHandler('', _) output: %s\n", resNoPass)
		t.Error("func addGHCredsHandler() failed in case of (_, '')\n")
	}
}
