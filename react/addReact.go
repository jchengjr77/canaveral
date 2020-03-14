package react

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Path to create-react-app executable (local path)
const craPath string = "./node_modules/.bin/create-react-app"

// checkCRAExists looks in the parent path for create-react-app
// If there is no executable in the path, then throws error.
// Else, returns a message with the path of create-react-app
// * tested
func checkCRAExists() bool {
	if !fileExists("." + craPath) {
		fmt.Printf(
			"ERROR: didn't find 'create-react-app' in local path '%s'\n", craPath)
		return false
	}
	fmt.Printf("'create-react-app' executable is in '%s'\n", craPath)
	return true
}

// installCRA() installs create-react-app.
// REQUIRES (soft): create-react-app isn't already installed.
// This is a soft requirement bc npm will just update it if it is.
// ? untested, low priority
func installCRA() {
	fmt.Println(
		"\nLooks like you don't have create-react-app. Let's install it...")
	// Install it locally instead of globally
	installCRA := exec.Command("npm", "install", "create-react-app")
	installCRA.Stdout = os.Stdout
	installCRA.Stderr = os.Stderr
	err := installCRA.Run()
	check(err)
}

// AddReactProj launches a new react project.
// The main mechanism is similar to addProj (in root folder).
// However, create-react-app plays a large role in setup.
// * tested
func AddReactProj(projName string, wsPath string) {
	ws, err := ioutil.ReadFile(wsPath)
	check(err)
	err = os.MkdirAll(string(ws), os.ModePerm)
	check(err)
	if !checkCRAExists() {
		installCRA()
	}
	cmd := exec.Command(craPath, string(ws)+"/"+projName)
	// set correct pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	check(err)
	fmt.Printf("Added React project %s to workspace %s\n", projName, string(ws))
}
