// Package nativestore implementation taken from
// https://flowerinthenight.com/blog/2017/10/30/nativestore
package nativestore

import (
	"os"

	"github.com/docker/docker-credential-helpers/credentials"
)

// SetCreds sets creentials for the user in their native storage.
func SetCreds(label, url, username, secret string) error {
	creds := credentials.Credentials{
		ServerURL: url,
		Username:  username,
		Secret:    secret,
	}
	credentials.SetCredsLabel(label)
	if os.Getenv("CredentialsTest") == "true" {
		return mockStore.Add(&creds)
	}
	return store.Add(&creds)
}

// FetchCreds gets user credentials from the native storage, looking up by URL.
func FetchCreds(label, url string) (string, string, error) {
	credentials.SetCredsLabel(label)
	if os.Getenv("CredentialsTest") == "true" {
		return mockStore.Get(url)
	}
	return store.Get(url)
}

// DeleteCreds deletes credentials from the native storage, looking up by URL.
func DeleteCreds(label, url string) error {
	credentials.SetCredsLabel(label)
	if os.Getenv("CredentialsTest") == "true" {
		return mockStore.Delete(url)
	}
	return store.Delete(url)
}
