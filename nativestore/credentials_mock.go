package nativestore

import (
	"github.com/docker/docker-credential-helpers/credentials"
)

type mockstorage struct {
	creds map[string]*credentials.Credentials
}

func newMockstorage() *mockstorage {
	return &mockstorage{
		creds: make(map[string]*credentials.Credentials),
	}
}

func (m *mockstorage) Add(creds *credentials.Credentials) error {
	m.creds[creds.ServerURL] = creds
	return nil
}

func (m *mockstorage) Delete(serverURL string) error {
	if _, found := m.creds[serverURL]; found {
		delete(m.creds, serverURL)
		return nil
	}
	return credentials.NewErrCredentialsNotFound()
}

func (m *mockstorage) Get(serverURL string) (string, string, error) {
	cred, found := m.creds[serverURL]
	if !found {
		return "", "", credentials.NewErrCredentialsNotFound()
	}
	return cred.Username, cred.Secret, nil
}

var mockStore = newMockstorage()
