package react

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"

	"github.com/jchengjr77/canaveral/lib"
)

// checkCRAExists looks in the parent for create-react-app
// If there is no executable in the path, then throws error.
// Else, returns a message with the path of create-react-app
// * tested
func checkCRAExists(craPath string) bool {
	if !lib.FileExists(craPath) {
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
func installCRA() (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	fmt.Println(
		"\nLooks like you don't have create-react-app. Let's install it...")
	// Install it locally instead of globally
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	err = os.MkdirAll(home+"/.canaveral", os.ModePerm)
	err = os.Chdir(home + "/.canaveral")
	installCRA := exec.Command("npm", "install", "create-react-app")
	installCRA.Stdout = os.Stdout
	installCRA.Stderr = os.Stderr
	err = installCRA.Run()
	lib.Check(err)
	return nil
}

// setCRAPath generates the path to create-react-app
// ? untested, trivial
func setCRAPath() (res string, finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	return home + "/.canaveral/node_modules/.bin/create-react-app", nil
}

// AddReactProj launches a new react project.
// The main mechanism is similar to addProj (in root folder).
// However, create-react-app plays a large role in setup.
// * tested
func AddReactProj(projName string, wsPath string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	craPath, err := setCRAPath()
	lib.Check(err)
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	err = os.MkdirAll(string(ws), os.ModePerm)
	lib.Check(err)
	if !checkCRAExists(craPath) {
		installCRA()
	}
	cmd := exec.Command(craPath, string(ws)+"/"+projName)
	// set correct pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	lib.Check(err)
	fmt.Printf("Added React project %s to workspace %s\n", projName, string(ws))
	return nil
}
