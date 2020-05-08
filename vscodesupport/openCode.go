/*Package vscodesupport contains:
- functionality for opening canaveral projects in vscode
*/
package vscodesupport

import (
	"fmt"
)

// OpenCode will take in a project name, and open it in vscode.
// If such a project doesn't exist, it will return an error.
func OpenCode(projName string) error {
	fmt.Printf("We see your request for %s\n", projName)
	return nil
}
