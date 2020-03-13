package react

import (
	"fmt"
	"io/ioutil"
	"os"
)

// AddReactProj launches a new react project.
// The main mechanism is similar to addProj (in root folder).
// However, create-react-app plays a large role in setup.
func AddReactProj(projName string, wsPath string) {
	ws, err := ioutil.ReadFile(wsPath)
	check(err)
	err = os.MkdirAll(string(ws)+"/"+projName, os.ModePerm)
	check(err)
	err = os.Chdir(string(ws) + "/" + projName)
	check(err)
	fmt.Printf("Added React project %s to workspace %s\n", projName, string(ws))
}
