package react

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// checkCRAExists looks in the path for create-react-app
// If there is no executable in the path, then throws error.
// Else, returns a message with the path of create-react-app
// ! untested
func checkCRAExists() bool {
	path, err := exec.LookPath("create-react-app")
	if err != nil {
		fmt.Printf("ERROR: didn't find 'create-react-app' in path\n")
		return false
	}
	fmt.Printf("'create-react-app' executable is in '%s'\n", path)
	return true
}

// installCRA() installs create-react-app.
// REQUIRES (soft): create-react-app isn't already installed.
// This is a soft requirement bc npm will just update it if it is.
// ! untested
func installCRA() {
	fmt.Println(
		"\nLooks like you don't have create-react-app. Let's install it...\n")
	installCRA := exec.Command("npm", "install", "-g", "create-react-app")
	installCRA.Stdout = os.Stdout
	installCRA.Stderr = os.Stderr
	err := installCRA.Run()
	check(err)
}

// AddReactProj launches a new react project.
// The main mechanism is similar to addProj (in root folder).
// However, create-react-app plays a large role in setup.
// ! untested
func AddReactProj(projName string, wsPath string) {
	ws, err := ioutil.ReadFile(wsPath)
	check(err)
	err = os.Chdir(string(ws))
	check(err)
	if !checkCRAExists() {
		installCRA()
	}
	cmd := exec.Command("create-react-app", projName)
	// set correct pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	check(err)
	fmt.Printf("Added React project %s to workspace %s\n", projName, string(ws))
}
