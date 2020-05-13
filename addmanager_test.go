package main

import (
	"canaveral/lib"
	"os"
	"os/user"
	"testing"
)

func TestAddProj(t *testing.T) {
	origOut := lib.RedirOut()
	defer func() {
		lib.ResetOut(origOut)
	}()
	testProjName := "testProj"
	tempusr, err := user.Current()
	lib.Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_ws"
	wsPath := tempHome + "/tmpcnavrlws"
	f, err := os.Create(wsPath)
	lib.Check(err)
	defer os.Remove(wsPath)
	defer f.Close()
	f.WriteString(newPath)
	addProj(testProjName, wsPath)
	if !lib.DirExists(newPath + "/" + testProjName) {
		t.Errorf("func addProj() failed to create ws at path: %s\n",
			newPath+"/"+testProjName)
		return
	}
	os.RemoveAll(newPath)
}

func TestAddProjecthandler(t *testing.T) {
	res := lib.CaptureOutput(func() {
		addProjectHandler("", "default", false)
	})
	want := "Please provide a project name.\n(For more info, 'canaveral --help')\n"
	if res != want {
		t.Logf("Ideal: %s\n", want)
		t.Logf("Result: %s\n", res)
		t.Error("func addProjectHandler() returned wrong on empty string.")
	}
	// fileExists() and addProj() are independently tested.
}
