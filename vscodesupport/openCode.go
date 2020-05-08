/*Package vscodesupport contains:
- functionality for opening canaveral projects in vscode
*/
package vscodesupport

import (
	"canaveral/lib"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// OpenCode will take in a project name, and open it in vscode.
// If such a project doesn't exist, it will return an error.
func OpenCode(projName string, configPath string) error {
	if projName == "" {
		fmt.Println("Please provide a project name.")
		fmt.Println("(For more info, 'canaveral --help')")
		return nil
	} else if !lib.FileExists(configPath) {
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs a workspace to add projects to.")
		fmt.Println("(For help, type 'canaveral --help')")
		return nil
	}
	fmt.Printf("Attempting to open Project '%s'\n", projName)
	ws, err := ioutil.ReadFile(configPath)
	lib.Check(err)
	if !lib.DirExists(string(ws) + "/" + projName) {
		return errors.New("Project '" + projName + "' not found")
	}
	err = os.Chdir(string(ws) + "/" + projName)
	lib.Check(err)
	openVSCode := exec.Command("code", ".")
	// set correct pipes
	openVSCode.Stdout = os.Stdout
	openVSCode.Stderr = os.Stderr
	err = openVSCode.Run()
	lib.Check(err)
	return nil
}
