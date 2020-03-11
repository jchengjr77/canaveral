package nativestore

import (
	"os"
	"testing"
)

// test configuration directory location
func TestSetGetCreds(t *testing.T) {
	os.Setenv("CredentialsTest", "true")
	label := "test credentials"
	url := "https://api.github.com"
	SetCreds(label, url, "username", "password")
	user, secret, err := FetchCreds(label, url)
	if err == nil {
		if user != "username" {
			t.Errorf("Expecting username, got %s", user)
		}
		if secret != "password" {
			t.Errorf("Expecting password, got %s", secret)
		}
	} else {
		t.Errorf("Failed to fetch, got error: %s", err)
	}
	if delErr := DeleteCreds(label, url); delErr != nil {
		t.Errorf("Delete failed on error: %s", delErr)
	}
	if reUsr, reSec, refetchErr := FetchCreds(label, url); refetchErr == nil {
		t.Errorf("Delete failed silently, still found username %s and secret %s", reUsr, reSec)
	}
}
