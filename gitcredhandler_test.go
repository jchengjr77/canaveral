package main

import (
	"canaveral/nativestore"
	"testing"
)

func TestAddGitCredsHandler(t *testing.T) {
	resNoUsr := captureOutput(
		func() {
			addGitCredsHandler("", "password")
		})
	if resNoUsr != "A git username is required. Please provide one.\n" {
		t.Logf("addGitCredsHandler('', _) output: %s\n", resNoUsr)
		t.Error("func addGitCredsHandler() failed in case of ('', _)\n")
	}
	resNoPass := captureOutput(
		func() {
			addGitCredsHandler("username", "")
		})
	if resNoPass != "A git password is required. Please provide one.\n" {
		t.Logf("addGitCredsHandler('', _) output: %s\n", resNoPass)
		t.Error("func addGitCredsHandler() failed in case of (_, '')\n")
	}
	resValid := captureOutput(
		func() {
			addGitCredsHandler("username", "password")
		})
	if resValid != "Adding git account: username\n" {
		t.Logf("addGitCredsHandler('u', 'p') output: %s\n", resValid)
		t.Error("func addGitCredsHandler() failed on valid input\n")
	}
	fetchUsr, fetchSec, fetchErr := nativestore.FetchCreds("github credentials", "https://api.github.com")
	if fetchErr == nil {
		if fetchUsr != "username" {
			t.Errorf("Added incorrect username. Expected username, found %s", fetchUsr)
		} else if fetchSec != "password" {
			t.Errorf("Added incorrect password. Expected password, found %s", fetchSec)
		}
	} else {
		t.Errorf("Fetch exited on error: %s", fetchErr)
	}
	if remErr := remGitCredsHandler(); remErr != nil {
		t.Errorf("Removed exited on error: %s", remErr)
	}
	manualRem := nativestore.DeleteCreds("github credentials", "https://api.github.com")
	if manualRem == nil {
		t.Error("Manual remove succeeded but should have failed as credentials should have already been removed.\n")
	}
}
