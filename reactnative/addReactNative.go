package reactnative

import (
	"bufio"
	"canaveral/lib"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

// installExpo checks first that npm is installed.
// If it is, then it uses npm to globally install expo
// * tested
func installExpo() error {
	if !lib.CheckToolExists("npm") {
		return errors.New("prerequisite tool 'npm' is not installed")
	}
	cmd := exec.Command("npm", "i", "-g", "expo-cli")
	// set correct pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

// confirmInstall listens for user confirmation and returns a boolean
// Takes in an input channel to increase testability
// * tested
func confirmInstall(stdin io.Reader) (res bool, finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	fmt.Printf("Allow canaveral to globally install 'expo'? ('y' or 'n')\n>")
	reader := bufio.NewReader(stdin)
	response, err := reader.ReadByte()
	lib.Check(err)
	return (response == 'y'), nil
}

// AddReactNativeProj takes a project name and path, and inits an expo project.
// First, must check if the user has installed expo or not.
// If yes, then continue. If no, then ask to install.
// Running 'expo init' with the project name will create a new project.
// * tested
func AddReactNativeProj(projName string, wsPath string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	expoInstalled := lib.CheckToolExists("expo")
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	err = os.MkdirAll(string(ws), os.ModePerm)
	lib.Check(err)
	if !expoInstalled {
		fmt.Println("Looks like you haven't installed 'expo'...")
		confirm, err := confirmInstall(os.Stdin)
		lib.Check(err)
		if !confirm {
			fmt.Printf("Install rejected. Aborting creation of '%s'\n", projName)
			return
		}
		fmt.Println("Attempting to install expo using npm")
		err = installExpo()
		lib.Check(err)
	}
	cmd := exec.Command("expo", "i", "-t", "blank", "--npm", "--name", projName, projName)
	// set correct pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Go to canaveral workspace
	err = os.Chdir(string(ws))
	lib.Check(err)
	// expo init projName
	err = cmd.Run()
	lib.Check(err)
	fmt.Printf("Added React Native project %s to workspace %s\n", projName, string(ws))
	return nil
}
