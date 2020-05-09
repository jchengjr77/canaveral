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

// Status prints current git status in a git directory
func Status() {
	gitStatus := exec.Command("git", "status")
	gitStatus.Stdout = os.Stdout
	gitStatus.Stdin = os.Stdin
	gitStatus.Stderr = os.Stderr
	err := gitStatus.Run()
	lib.Check(err)
}

func Add(files []string) {
	gitAdd := exec.Command("git", "add")
	gitAdd.Stdout = os.Stdout
	gitAdd.Stdin = os.Stdin
	gitAdd.Stderr = os.Stderr
	gitAdd.Args = append(gitAdd.Args, files...)
	err := gitAdd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	// lib.Check(err)
}

func Commit() {
	gitCommit := exec.Command("git", "commit")
	gitCommit.Stdout = os.Stdout
	gitCommit.Stdin = os.Stdin
	gitCommit.Stderr = os.Stderr
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
