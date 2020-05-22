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

func TestAddMake(t *testing.T) {
	// testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_C_ws"

	// Check state
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
		t.Errorf("Bad state, Makefile already exists\n")
		return
	}

	// Setup state
	lib.CreateFile("Makefile")
	lib.CreateFile("test.c")
	lib.CreateFile("test.h")
	mainfile, err := os.OpenFile("test.c", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := mainfile.Close()
		lib.Check(err)
	}()

	// Write test.c file
	include := "#include <stdio.h>"
	main := "int main() {\n\tprintf(\"Canaveral is awesome!" + `\n` + "\");\n\treturn 0;\n}\n"
	_, err = mainfile.WriteString(include + "\n\n" + main)
	if err != nil {
		t.Errorf("Write test.c file failed with error: " + err.Error() + "\n")
		return
	}

	// Write Makefile
	make, err := os.OpenFile("Makefile", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := make.Close()
		lib.Check(err)
	}()
	topComment := "# Basic makefile generated by Canaveral"
	CC := "CC = gcc"
	CFLAGS := "CFLAGS = -g -Og -Wall -std=c99 -I."
	DEPS := "DEPS = test.h"
	oFiles := "%.o: %.c $(DEPS)\n\t$(CC) -c -o $@ $< $(CFLAGS)"
	outfile := "test: test.o\n\t$(CC) -o $@ $^ $(CFLAGS)"
	PHONY := ".PHONY: clean"
	clean := "clean:\n\trm -f *.o *~"

	_, err = make.WriteString(topComment + "\n\n" + CC + "\n" + CFLAGS + "\n" + DEPS + "\n\n" + oFiles + "\n\n" + outfile + "\n\n" + PHONY + "\n" + clean)
	if err != nil {
		t.Errorf("Write Makefile failed with error: " + err.Error() + "\n")
		return
	}

	// Make should work
	err = exec.Command("make").Run()
	if err != nil {
		t.Errorf("make failed with error " + err.Error() + "\n")
		return
	}
	out, err := exec.Command("./test").Output()
	if string(out) != "Canaveral is awesome!\n" {
		t.Errorf("make compiled but file had wrong output\n")
		return
	}
	if err != nil {
		t.Errorf("running ./test failed with error " + err.Error() + "\n")
		return
	}

	// Create test2.c and test2.h
	lib.CreateFile("test2.c")
	lib.CreateFile("test2.h")
	mainfile2, err := os.OpenFile("test2.c", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := mainfile2.Close()
		lib.Check(err)
	}()
	mainfile2h, err := os.OpenFile("test2.h", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := mainfile2h.Close()
		lib.Check(err)
	}()

	// Write test2.c and test2.h
	include2 := "#include <stdio.h>"
	main2 := "int printCan() {\n\tprintf(\"Canaveral\");\n\treturn 0;\n}\n"
	_, err = mainfile2.WriteString(include2 + "\n\n" + main2)
	if err != nil {
		t.Errorf("Write test2.c file failed with error: " + err.Error() + "\n")
		return
	}

	test2h := "int printCan(void);\n"
	_, err = mainfile2h.WriteString(test2h)
	if err != nil {
		t.Errorf("Write test2.h file failed with error: " + err.Error() + "\n")
		return
	}

	// Delete and rewrite test.c
	os.Remove("test.c")
	lib.CreateFile("test.c")
	mainfilec, err := os.OpenFile("test.c", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := mainfilec.Close()
		lib.Check(err)
	}()

	// Write test.c file
	includec := "#include <stdio.h>\n#include \"test2.h\""
	mainc := "int main() {\n\tprintCan();\n\tprintf(\" is awesome!" + `\n` + "\");\n\treturn 0;\n}\n"
	_, err = mainfilec.WriteString(includec + "\n\n" + mainc)
	if err != nil {
		t.Errorf("rewrite test.c file failed with error: " + err.Error() + "\n")
		return
	}

	// Add test2.c to makefile
	err = AddToMake("test2.c", "test")
	if err != nil {
		t.Errorf("Expected AddToMake to pass, instead got: " + err.Error() + "\n")
		return
	}

	// Make should work
	err = exec.Command("make").Run()
	if err != nil {
		t.Errorf("make failed with error " + err.Error() + "\n")
		return
	}
	out, err = exec.Command("./test").Output()
	if string(out) != "Canaveral is awesome!\n" {
		t.Errorf("make compiled but file had wrong output\n")
		return
	}
	if err != nil {
		t.Errorf("running ./test failed with error " + err.Error() + "\n")
		return
	}

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
