package git

import (
	"bufio"
	"canaveral/lib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// InitRepo initializes a git repo in the current directory
// * tested
func InitRepo(wsPath, project string) {
	if project != "" {
		// Get workspace path
		ws, err := ioutil.ReadFile(wsPath)
		lib.Check(err)
		err = os.Chdir(string(ws) + "/")
		lib.Check(err)
		if !lib.DirExists(project) {
			fmt.Println("No project found: ", project)
			return
		}
		err = os.Chdir(string(ws) + "/" + project)
		lib.Check(err)
	}
	initRepo := exec.Command("git", "init")
	initRepo.Stdout = os.Stdout
	initRepo.Stdin = os.Stdin
	initRepo.Stderr = os.Stderr
	err := initRepo.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Status prints current git status in a git directory
// * tested
func Status(wsPath, project string) {
	if project != "" {
		// Get workspace path
		ws, err := ioutil.ReadFile(wsPath)
		lib.Check(err)
		err = os.Chdir(string(ws) + "/")
		lib.Check(err)
		if !lib.DirExists(project) {
			fmt.Println("No project found: ", project)
			return
		}
		err = os.Chdir(string(ws) + "/" + project)
		lib.Check(err)
	}
	gitStatus := exec.Command("git", "status")
	gitStatus.Stdout = os.Stdout
	gitStatus.Stdin = os.Stdin
	gitStatus.Stderr = os.Stderr
	err := gitStatus.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Add performs a git add on the specified files
// * tested
func Add(files []string, wsPath string, project string) {
	if project != "" {
		// Get workspace path
		ws, err := ioutil.ReadFile(wsPath)
		lib.Check(err)
		err = os.Chdir(string(ws) + "/")
		lib.Check(err)
		if !lib.DirExists(project) {
			fmt.Println("No project found: ", project)
			return
		}
		err = os.Chdir(string(ws) + "/" + project)
		lib.Check(err)
	}
	gitAdd := exec.Command("git", "add")
	gitAdd.Stdout = os.Stdout
	gitAdd.Stdin = os.Stdin
	gitAdd.Stderr = os.Stderr
	gitAdd.Args = append(gitAdd.Args, files...)
	err := gitAdd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getStaged() []string {
	out, err := exec.Command("git", "diff", "--name-only", "--staged").Output()
	lib.Check(err)
	staged := strings.Split(string(out), "\n")
	var ret []string
	for i := range staged {
		curr := strings.TrimSpace(staged[i])
		if curr != "" {
			ret = append(ret, curr)
		}
	}
	return ret
}

func confirmCommit(stdin io.Reader) bool {
	fmt.Printf("Would you still like to commit? ('y' or 'n')> ")
	reader := bufio.NewReader(stdin)
	response, err := reader.ReadByte()
	lib.Check(err)
	return (response == 'y')
}

// Commit performs a git commit on added files
// * tested
// ! reminders untested
func Commit(commitMessage, wsPath, project string) {
	if project != "" {
		// Get workspace path
		ws, err := ioutil.ReadFile(wsPath)
		lib.Check(err)
		err = os.Chdir(string(ws) + "/")
		lib.Check(err)
		if !lib.DirExists(project) {
			fmt.Println("No project found: ", project)
			return
		}
		err = os.Chdir(string(ws) + "/" + project)
		lib.Check(err)
	}
	reminders := loadReminders()
	stagedFiles := getStaged()
	sawRems := false
	confirm := true
	for _, file := range stagedFiles {
		sawRems = sawRems || checkReminders(file, false, reminders)
	}
	if sawRems {
		confirm = confirmCommit(os.Stdin)
	}
	if !confirm {
		return
	}
	gitCommit := exec.Command("git", "commit")
	gitCommit.Stdout = os.Stdout
	gitCommit.Stdin = os.Stdin
	gitCommit.Stderr = os.Stderr
	if commitMessage != "" {
		gitCommit.Args = append(gitCommit.Args, "-m", "\""+commitMessage+"\"")
	}
	err := gitCommit.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Checks if searchFor is a line in file
// ? untested, low priority
func inFile(file io.Reader, searchFor string) bool {
	s := bufio.NewScanner(file)
	if err := s.Err(); err != nil {
		fmt.Println(err)
	}
	for s.Scan() {
		if s.Text() == searchFor {
			return true
		}
	}
	return false
}

// Ignore adds files to the .gitignore file in the current directory
// * tested
func Ignore(files []string, wsPath, project string) {
	if project != "" {
		// Get workspace path
		ws, err := ioutil.ReadFile(wsPath)
		lib.Check(err)
		err = os.Chdir(string(ws) + "/")
		lib.Check(err)
		if !lib.DirExists(project) {
			fmt.Println("No project found: ", project)
			return
		}
		err = os.Chdir(string(ws) + "/" + project)
		lib.Check(err)
	}
	var startsEmpty = false
	gitignore, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_WRONLY, 0644)
	ignoreReader, err := os.OpenFile(".gitignore", os.O_RDONLY, 0644)
	lib.Check(err)
	stat, err := gitignore.Stat()
	if stat.Size() == 0 {
		startsEmpty = true
	}
	defer gitignore.Close()
	for idx, file := range files {
		ignoreReader.Seek(0, io.SeekStart)
		if inFile(ignoreReader, file) {
			fmt.Println("Skipping", file, "because it is already being ignored.")
			continue
		}
		if idx == 0 && !startsEmpty {
			gitignore.Write([]byte{'\n'})
		} else if idx > 0 {
			gitignore.Write([]byte{'\n'})
		}
		_, err = gitignore.WriteString(file)
	}
	lib.Check(err)
}
