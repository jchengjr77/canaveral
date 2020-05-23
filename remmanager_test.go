package main

import (
	"bytes"
	"canaveral/lib"
	"os"
	"os/user"
	"testing"
)

func TestConfirmDelete(t *testing.T) {
	origOut := lib.RedirOut()
	defer func() {
		lib.ResetOut(origOut) // reset it
	}()
	var stdin bytes.Buffer // testable io
	stdin.WriteByte(byte('y'))
	res, err := confirmDelete("testProj", &stdin)
	lib.Check(err)
	if !res {
		t.Errorf("func confirmDelete() did not return true when fed 'y'")
	}
	stdin.WriteByte(byte('n'))
	res, err = confirmDelete("testProj", &stdin)
	lib.Check(err)
	if res {
		t.Errorf("func confirmDelete() did not return false when fed 'n'")
	}
	stdin.Write([]byte("foo"))
	res, err = confirmDelete("testProj", &stdin)
	lib.Check(err)
	if res {
		t.Errorf("func confirmDelete() did not return false when fed 'foo'")
	}
}

func TestTryRemProj(t *testing.T) {
	tempusr, err := user.Current()
	lib.Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_ws/"
	err = os.MkdirAll(newPath, os.ModePerm)
	lib.Check(err)
	f, err := os.Create(newPath + "testProj")
	defer os.RemoveAll(newPath)
	defer f.Close()
	wsF, err := os.Create(tempHome + "/tempWSPath")
	defer os.Remove(tempHome + "/tempWSPath")
	defer wsF.Close()
	wsF.Write([]byte(newPath))
	res, err := tryRemProj("testProjFoo", tempHome+"/tempWSPath")
	// should be false
	lib.Check(err)
	if res {
		t.Logf("Path: %s\n", newPath)
		t.Errorf("func tryRemProj() returned true. Should be false.")
	}
	// omitted success case, as it is partially tested by confirmDelete
}

func TestRemProjectHandler(t *testing.T) {
	res := lib.CaptureOutput(func() {
		remProjectHandler("")
	})
	if res !=
		"Cannot remove an unspecified project. Please provide the project name.\n" {
		t.Logf("remProjectHandler('') output: %s\n", res)
		t.Error("func remProjectHandler() failed in case of ''")
	}
	// omitted case 2 and 3, as they are tested in the aggregate
}
