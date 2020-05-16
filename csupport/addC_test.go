package csupport

import (
	"canaveral/lib"
	"os"
	"os/exec"
	"os/user"
	"testing"
)

func TestCreateMainFile(t *testing.T) {
	testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_C_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists\n")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)
	if lib.FileExists(testProjName + ".c") {
		t.Errorf("Bad State, main file already exists\n")
		return
	}
	err = createMainFile(testProjName)
	if err != nil {
		t.Errorf("Create main file failed with error: %s\n", err.Error())
		return
	}
	if !lib.FileExists(testProjName + ".c") {
		t.Errorf("Failed to actually make main file\n")
		return
	}
	err = exec.Command("gcc", testProjName+".c", "-o", "test").Run()
	if err != nil {
		t.Errorf("Failed to compile with error: %s\n", err.Error())
		return
	}
	outByte, err := exec.Command("./test").Output()
	out := string(outByte)
	if out != "Canaveral is awesome!\n" {
		t.Errorf("AddCProj failed to make correct first C file\n")
		return
	}
}

func TestCreateMakeFile(t *testing.T) {
	testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_C_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists\n")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)
	if lib.FileExists("Makefile") {
		t.Errorf("Bad State, Makefile already exists\n")
		return
	}
	err = createMakeFile(testProjName)
	if err != nil {
		t.Errorf("Create Makefile failed with %s\n", err.Error())
		return
	}
	if !lib.FileExists("Makefile") {
		t.Errorf("Failed to actually create Makefile\n")
		return
	}
	err = createMainFile(testProjName)
	if err != nil {
		t.Errorf("Creating main file failed with error: %s\n", err.Error())
		return
	}
	err = exec.Command("make").Run()
	if err != nil {
		t.Errorf("Make failed with error: %s\n", err.Error())
		return
	}
	outByte, err := exec.Command("./" + testProjName).Output()
	out := string(outByte)
	if out != "Canaveral is awesome!\n" {
		t.Errorf("AddCProj failed to make correct first C file\n")
		return
	}
}

func TestCreateReadme(t *testing.T) {
	testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_C_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists\n")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)
	if lib.FileExists("README.md") {
		t.Errorf("Bad State, readme already exists\n")
		return
	}
	err = createREADME(testProjName)
	if err != nil {
		t.Errorf("Create readme failed with %s\n", err.Error())
		return
	}
	if !lib.FileExists("README.md") {
		t.Errorf("Create readme failed to create a readme\n")
		return
	}
	err = os.Remove("README.md")
	lib.Check(err)
}

func TestAddCProj(t *testing.T) {
	testProjName := "testproj"
	tempusr, err := user.Current()
	lib.Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_C_test_ws"
	wsPath := tempHome + "/tmpcnavrlws_C"
	f, err := os.Create(wsPath)
	lib.Check(err)
	defer func() {
		f.Close()
		os.Remove(wsPath)
	}()
	f.WriteString(newPath)
	err = os.Chdir("../")
	lib.Check(err)
	dir, err := os.Getwd()
	lib.Check(err)
	t.Logf("\nCurrent Dir: %s\n", dir)
	AddCProj(testProjName, wsPath)
	if !lib.FileExists("Makefile") {
		t.Errorf("AddCProj failed to create makefile\n")
		return
	}
	if !lib.FileExists(testProjName + ".c") {
		t.Errorf("AddCProj failed to create main .c file\n")
		return
	}
	if !lib.FileExists(testProjName + ".h") {
		t.Errorf("AddCProj failed to create main .h file\n")
		return
	}
	err = exec.Command("make").Run()
	if err != nil {
		t.Errorf("Make failed with error: %s\n", err.Error())
		return
	}
	outByte, err := exec.Command("./" + testProjName).Output()
	out := string(outByte)
	if out != "Canaveral is awesome!\n" {
		t.Errorf("AddCProj failed to make correct first C file\n")
		return
	}
	os.RemoveAll(newPath)
}
