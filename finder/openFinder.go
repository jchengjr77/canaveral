/*Package finder contains:
	- functionality for opening canaveral project in finder
// ? NOTE: this feature is targeted at macOS users
*/
package finder

import (
	"canaveral/lib"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// OpenFinder will take a project name and config path,
// and attempt to open the specified project in a new finder window.
// Returns error if it fails. Otherwise returns nil.
// * tested
func OpenFinder(projName string, configPath string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	if projName == "" {
		// Check for blank project name
		fmt.Println("Please provide a project name.")
		fmt.Println("(For more info, 'canaveral --help')")
		return errors.New("No project name specified")
	} else if !lib.FileExists(configPath) {
		// Check that workspace is set
		fmt.Println("No canaveral workspace set. Please specify a workspace.")
		fmt.Println(
			"Canaveral needs to know where to look for your projects.")
		fmt.Println("(For help, type 'canaveral --help')")
		return errors.New("No canaveral workspace set")
	}
	fmt.Printf("Attempting to open Project '%s'\n", projName)
	ws, err := ioutil.ReadFile(configPath)
	lib.Check(err)
	if !lib.DirExists(string(ws) + "/" + projName) {
		// Check that project exists
		return errors.New("Project '" + projName + "' not found")
	}
	err = os.Chdir(string(ws) + "/" + projName)
	lib.Check(err)
	// Open project in file explorer (OS dependent)
	openFinder := openCmd
	// set correct pipes
	openFinder.Stdout = os.Stdout
	openFinder.Stderr = os.Stderr
	err = openFinder.Run()
	lib.Check(err)
	return nil
}
