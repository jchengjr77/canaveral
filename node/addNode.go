package node

import (
	"canaveral/lib"
	"os"
)

// AddNodeProj launches a new node project.
// The main mechanism is similar to addProj (in root folder).
// However, npm init plays a large role in setup.
// ! untested
func AddNodeProj(projName string, wsPath string) {
	// Navigate to canaveral workspace
	err := os.Chdir(wsPath)
	lib.Check(err)

}
