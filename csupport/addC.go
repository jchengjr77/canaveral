package csupport

import (
	"canaveral/lib"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// createMainFile creates a basic projName.c file that uses stdlib.h
// to print "Canaveral is awesome!" in the main function
// * tested
func createMainFile(projName string) error {
	err := lib.CreateFile(projName + ".c")
	if err != nil {
		return err
	}
	mainfile, err := os.OpenFile(projName+".c", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := mainfile.Close()
		lib.Check(err)
	}()

	include := "#include <stdio.h>"
	main := "int main() {\n\tprintf(\"Canaveral is awesome!" + `\n` + "\");\n\treturn 0;\n}\n"

	_, err = mainfile.WriteString(include + "\n\n" + main)
	return err
}

// createREADME creates a README.md file
// * tested
func createREADME(projName string) error {
	err := lib.CreateFile("README.md")
	if err != nil {
		return err
	}
	readme, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := readme.Close()
		lib.Check(err)
	}()
	header := "# " + projName
	making := "## Compiling\nTo compile " + projName + "simply download it and run `make`. To clean old builds, run `make clean`"
	canaveral := "###### Created with [Canaveral](https://github.com/jchengjr77/canaveral)"
	_, err = readme.WriteString(header + "\n\n" + making + "\n\n" + canaveral)
	return err
}

// createMakeFile creates a make file that compiles the projName.c file into
// a binary called projName
// * tested
func createMakeFile(projName string) error {
	err := lib.CreateFile("Makefile")
	if err != nil {
		return err
	}

	make, err := os.OpenFile("Makefile", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		err := make.Close()
		lib.Check(err)
	}()
	topComment := "# Basic makefile generated by Canaveral"
	CC := "CC = gcc"
	CFLAGS := "CFLAGS = -g -Og -Wall -std=c99 -I."
	DEPS := "DEPS = " + projName + ".h"
	oFiles := "%.o: %.c $(DEPS)\n\t$(CC) -c -o $@ $< $(CFLAGS)"
	outfile := projName + ": " + projName + ".o\n\t$(CC) -o $@ $^ $(CFLAGS)"
	PHONY := ".PHONY: clean"
	clean := "clean:\n\trm -f *.o *~"

	_, err = make.WriteString(topComment + "\n\n" + CC + "\n" + CFLAGS + "\n" + DEPS + "\n\n" + oFiles + "\n\n" + outfile + "\n\n" + PHONY + "\n" + clean)
	return err
}

// AddToMake adds a dependency to a Makefile conforming to the style of the one
// created by canaveral. It adds [filename].h to the DEPS definition and
// [filename].o to the make definition for the binary file [binaryAddTo] line of
// the form [binaryAddTo]: [file].o
// ! untested
func AddToMake(filename, binaryAddTo string) error {
	if !lib.FileExists(filename) {
		return errors.New("File not found: " + filename)
	}
	if len(filename) < 2 || filename[len(filename)-2:] != ".c" {
		return errors.New("Currently addmake only supports adding simple rules for compiling additional c files")
	}
	if !lib.FileExists("Makefile") {
		return errors.New("Cannot find Makefile")
	}
	makefile, err := ioutil.ReadFile("Makefile")
	if err != nil {
		return err
	}

	addedDependency := false
	addedOutfile := false

	lines := strings.Split(string(makefile), "\n")
	for i, line := range lines {
		if len(line) >= 4 && line[0:4] == "DEPS" {
			addedDependency = true
			lines[i] = line + " " + filename[:len(filename)-2] + ".h"
		}
		if len(line) >= len(binaryAddTo) && line[:len(binaryAddTo)] == binaryAddTo {
			addedOutfile = true
			lines[i] = line + " " + filename[:len(filename)-2] + ".o"
		}
	}

	if addedDependency && addedOutfile {
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile("Makefile", []byte(output), 0644)
		return err
	}
	if !addedDependency {
		return errors.New("Currently addmake only supports very simple Makefile additions, and the found Makefile doesn't conform to the necessary standards. Missing line: DEPS [file].h")
	}
	return errors.New("Currently addmake only supports very simple Makefile additions, and the found Makefile doesn't conform to the necessary standards. Missing line: " + binaryAddTo + ": [file].o")
}

// AddCProj launches a new C project.
// The main mechanism is similar to addProj (in root folder).
// It will create a Makefile, projName.c, and a projName.h file
// * tested
func AddCProj(projName string, wsPath string) {
	// Get workspace path
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	err = os.MkdirAll(string(ws)+"/"+projName, os.ModePerm)
	lib.Check(err)
	// Navigate to canaveral workspace
	err = os.Chdir(string(ws) + "/" + projName)
	lib.Check(err)

	err = createMainFile(projName)
	lib.Check(err)
	err = lib.CreateFile(projName + ".h")
	lib.Check(err)
	err = createMakeFile(projName)
	lib.Check(err)
	err = createREADME(projName)
	lib.Check(err)

	fmt.Printf("Added C Project %s to workspace %s\n", projName, string(ws))
}
