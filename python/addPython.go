package python

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/jchengjr77/canaveral/lib"
)

func checkToolExists(toolName string) bool {
	if toolName == "" {
		fmt.Println("Cannot pass in a blank toolname")
		return false
	}
	if toolName[0] == '-' {
		fmt.Println("Cannot pass in an option as a toolname")
		return false
	}
	cmd := exec.Command("which", toolName)
	_, err := cmd.Output()
	// When 'which' cannot find the tool, it exits with status 1 (error)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("PROBLEM: %s not found in path\n", toolName)
		return false
	}
	fmt.Printf("%s found in path\n", toolName)
	return true
}

// Creates a conda environment called projName
// Conda must be installed for this to work
// * tested
func createCondaEnv(projName string) error {
	err := exec.Command("conda", "create", "-n", projName).Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("conda environment created: " + projName)
	return nil
}

// This is a very hacky solution but conda has a problem where you can't
// Easily use conda activate in subshells. Because go execute commands run
// in a subshell, it's hard to circumvent this. There are workarounds for
// bash scripts, but none were easily portable to a go function. So for now,
// this creates and deletes a shell file that uses the workarounds.
// ? Look for better solutions
// * tested
func activateAndSetupConda(projName string) error {
	err := lib.CreateFile("tmp.sh")
	defer os.Remove("tmp.sh")
	if err != nil {
		return err
	}
	tmp, err := os.OpenFile("tmp.sh", os.O_APPEND|os.O_WRONLY, 0644)
	defer tmp.Close()
	_, err = tmp.WriteString("eval \"$(conda shell.bash hook)\"\nconda activate " + projName + "\nconda install python\nconda install pip")
	activate := exec.Command("sh", "tmp.sh")
	activate.Stderr = os.Stderr
	activate.Stdout = os.Stdout
	err = activate.Run()
	return err
}

// createInstallSh creates the install_packages.sh file with the shebang and
// example comment showing how to use it to install packages with pip
// * tested
func createInstallSh() (finalErr error) {
	err := lib.CreateFile("install_packages.sh")
	if err != nil {
		return err
	}
	install, err := os.OpenFile("install_packages.sh", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		finalErr = install.Close()
	}()
	_, err = install.WriteString("#!/bin/sh\n\n# Ex: pip install [package]==[version]\n")
	return err
}

// createREADME creates a README.md file, with optional conda setup information
// if conda is true
// * tested
func createREADME(projName string, conda bool) (finalErr error) {
	err := lib.CreateFile("README.md")
	if err != nil {
		return err
	}
	readme, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0644)
	defer func() {
		finalErr = readme.Close()
	}()
	header := "# " + projName
	setup := "## Setup and Installation\n"
	if conda {
		setup += "Create a conda environment\n`conda create " + projName + "`\nActivate conda environment\n`conda activate " + projName + "`\nInstall python\n`conda install python`\nInstall pip\n`conda install pip`\n"
	}
	install := "Install Packages\n`sh install_packages.sh`\n\n"
	canaveral := "###### Created with [Canaveral](https://github.com/jchengjr77/canaveral)"
	_, err = readme.WriteString(header + "\n\n" + setup + install + canaveral)
	return err
}

// AddPythonProj launches a new python project.
// It will create a conda environment if conda is installed,
// create a install_packages.sh shell file for installing requirements
// with pip, create a basic README.md, and a [projName].py file
// * tested
func AddPythonProj(projName string, wsPath string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	// Get workspace path
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	err = os.MkdirAll(string(ws)+"/"+projName, os.ModePerm)
	lib.Check(err)
	// Navigate to canaveral workspace
	err = os.Chdir(string(ws) + "/" + projName)
	lib.Check(err)
	conda := false
	if checkToolExists("conda") {
		fmt.Printf("Found conda. Creating a conda environment called %s\n", projName)
		err = createCondaEnv(projName)
		lib.Check(err)
		err = activateAndSetupConda(projName)
		lib.Check(err)
		conda = true
	} else {
		fmt.Println("Conda not found, skipping conda environment creation")
	}
	err = createInstallSh()
	lib.Check(err)

	err = createREADME(projName, conda)
	lib.Check(err)

	err = lib.CreateFile(projName + ".py")
	lib.Check(err)
	return nil
}
