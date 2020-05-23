package main

import (
	"canaveral/lib"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// formatProjects takes the raw ls output and removes extra lines.
// The command ls returns extra lines above (with . and .., etc).
// formatProject removes the top three lines from the output.
// * tested
func formatProjects(rawString string) string {
	lines := strings.Split(rawString, "\n")
	newLines := lines[3:]
	return strings.Join(newLines, "\n")
}

// showWorkspaceHandler checks the confDir for the workspace file.
// If such a file exists, it reads its contents and navigates to that path.
// Otherwise, it notifies the user that there is no canaveral workspace set.
// ? untested, low priority
func showWorkspaceHandler() (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	if !lib.FileExists(usrHome + confDir + wsFName) {
		fmt.Printf("Can't find workspace file in %s\n", usrHome+confDir+wsFName)
		fmt.Println("Please specify a canaveral workspace.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	ws, err := ioutil.ReadFile(usrHome + confDir + wsFName)
	lib.Check(err)
	fmt.Printf("\nYour canaveral path: %s\n", ws)
	fmt.Printf("\nCurrent canaveral projects:\n")
	err = os.Chdir(string(ws))
	lib.Check(err)
	cmd := exec.Command("ls", "-la")
	lib.Check(err)
	if runtime.GOOS == "windows" { // windows support
		cmd = exec.Command("tasklist")
	}
	out, err := cmd.Output()
	lib.Check(err)
	fmt.Println(formatProjects(string(out)))
	return nil
}

// setWorkspaceHandler takes in a new path and writes to the confDir.
// If the workspace file already exists, it overwrites it with the new path.
// Otherwise, it creates the workspace file and writes the path in.
// ? untested, low priority
func setWorkspaceHandler(newWorkspace string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	err := os.MkdirAll(usrHome+confDir, os.ModePerm)
	lib.Check(err)
	f, err := os.Create(usrHome + confDir + wsFName)
	// If file exists, truncates
	lib.Check(err)
	defer f.Close() // Close the file at the return of this function
	f.WriteString(newWorkspace)
	fmt.Printf("Set canaveral workspace to: %s\n", newWorkspace)
	return nil
}
