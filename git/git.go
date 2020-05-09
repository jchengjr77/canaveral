package git

import (
	"canaveral/lib"
	"fmt"
	"os"
	"os/exec"
)

// InitRepo initializes a git repo in the current directory
// ? untested
func InitRepo(projName string) {
	if projName == "" {
		fmt.Println("Please provide a repo name.")
		fmt.Println("(For more info, 'canaveral --help')")
	}
	initRepo := exec.Command("git", "init", projName)
	initRepo.Stdout = os.Stdout
	initRepo.Stdin = os.Stdin
	err := initRepo.Run()
	lib.Check(err)
}

func GitStatus() {
	gitStatus := exec.Command("git", "status")
	gitStatus.Stdout = os.Stdout
	gitStatus.Stdin = os.Stdin
	err := gitStatus.Run()
	lib.Check(err)
}

func GitAdd() {
	gitAdd := exec.Command("git", "add", ".")
	gitAdd.Stdout = os.Stdout
	gitAdd.Stdin = os.Stdin
	err := gitAdd.Run()
	lib.Check(err)
	// fmt.Println("Here")
}

func GitCommit() {
	gitCommit := exec.Command("git", "commit")
	gitCommit.Stdout = os.Stdout
	gitCommit.Stdin = os.Stdin
	err := gitCommit.Run()
	lib.Check(err)
}

func GitRm() {
	gitRm := exec.Command("git", "rm")
	gitRm.Stdout = os.Stdout
	gitRm.Stdin = os.Stdin
	err := gitRm.Run()
	lib.Check(err)
}
