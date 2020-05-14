package reactnative

import (
	"fmt"
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
