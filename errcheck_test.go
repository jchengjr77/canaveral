package main

import (
	"errors"
	"os"
	"os/user"
	"testing"
)

func checkPanic(t *testing.T, f func()) {
	defer func() bool {
		if r := recover(); r == nil {
			t.Log("The code did not panic")
			return false
		}
		t.Log("The code panicked")
		return true
	}()
	f()
}

func TestCheck(t *testing.T) {
	var err1 error = nil
	t.Log("Following check should not panic")
	checkPanic(t, func() {
		check(err1)
	})

	var err2 error = errors.New("Test Error")
	t.Log("Following check should panic")
	checkPanic(t, func() {
		check(err2)
	})
}

func TestFileExists(t *testing.T) {
	tempusr, err := user.Current()
	check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_dir/"
	if fileExists(newPath+testFile) == true {
		t.Errorf("function fileExists() found a non-existent file at: %s",
			tempHome+testFile)
		t.Errorf("Check that no file exists there already.")
	}

	err = os.MkdirAll(newPath, os.ModePerm)
	check(err)
	f, err := os.Create(newPath + testFile)
	check(err)
	defer os.RemoveAll(newPath)
	defer f.Close()
	if fileExists(newPath+testFile) == false {
		t.Errorf("function fileExists() failed to recognize file at: %s",
			tempHome+testFile)
	}

}

func TestDirExists(t *testing.T) {
	tempusr, err := user.Current()
	check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_dir/"
	if dirExists(newPath) == true {
		t.Errorf("function dirExists() found a non-existent dir at: %s",
			newPath)
		t.Errorf("Check that no dir exists there already.")
	}

	err = os.MkdirAll(newPath, os.ModePerm)
	check(err)
	defer os.RemoveAll(newPath)
	if dirExists(newPath) == false {
		t.Errorf("function dirExists() failed to recognize dir at: %s",
			newPath)
	}
}
