package node

import (
	"canaveral/lib"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// AddNodeProj launches a new node project.
// The main mechanism is similar to addProj (in root folder).
// However, npm init plays a large role in setup.
// ! untested
func AddNodeProj(projName string, wsPath string) (finalErr error) {
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
	cmd := exec.Command("npm", "init", "-y")
	// set correct pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	lib.Check(err)
	fmt.Printf("Added Node Project %s to workspace %s\n", projName, string(ws))
	return nil
}
