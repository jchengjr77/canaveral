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

// checkToolExists uses the 'which' command to find a specific tool.
// It then parses the output of the command, and checks if
// 'which' found the toolname in the path or not.
// * tested
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

// installExpo checks first that npm is installed.
// If it is, then it uses npm to globally install expo
// ? untested, trivial
func installExpo() error {
	if !checkToolExists("npm") {
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
func confirmInstall(stdin io.Reader) bool {
	fmt.Printf("Allow canaveral to globally install 'expo'? ('y' or 'n')\n>")
	reader := bufio.NewReader(stdin)
	response, err := reader.ReadByte()
	lib.Check(err)
	return (response == 'y')
}

// AddReactNativeProj takes a project name and path, and inits an expo project.
// First, must check if the user has installed expo or not.
// If yes, then continue. If no, then ask to install.
// Running 'expo init' with the project name will create a new project.
// * tested
func AddReactNativeProj(projName string, wsPath string) {
	expoInstalled := checkToolExists("expo")
	ws, err := ioutil.ReadFile(wsPath)
	lib.Check(err)
	err = os.MkdirAll(string(ws), os.ModePerm)
	lib.Check(err)
	if !expoInstalled {
		fmt.Println("Looks like you haven't installed 'expo'...")
		confirm := confirmInstall(os.Stdin)
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
}
