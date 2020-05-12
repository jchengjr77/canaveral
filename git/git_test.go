package git

import (
	"canaveral/lib"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"testing"
)

func TestStatus(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_git_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	// Initially no git repo
	ret := lib.CaptureOutput(Status)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\nexit status 128\n" {
		t.Errorf("Bad state. Expected directory not to be a git repo, instead got: %s\n", ret)
	}

	// Initialize git repo
	err = exec.Command("git", "init").Run()
	lib.Check(err)
	ret = lib.CaptureOutput(Status)
	if ret[:16] != "On branch master" {
		t.Errorf("Expected git repo to exist, instead got: %s\n", ret)
	}
}

func TestInit(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_git_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	// Initially no git repo
	retByte, err := exec.Command("git", "status").CombinedOutput()
	ret := string(retByte)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\n" {
		t.Errorf("Bad state. Expected directory not to be a git repo, instead got: %s\n", ret)
	}

	// Initialize repo
	ret = lib.CaptureOutput(InitRepo)
	if ret[:32] != "Initialized empty Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}

	// Reinitialize repo
	ret = lib.CaptureOutput(InitRepo)
	if ret[:37] != "Reinitialized existing Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}
}

func TestAdd(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_git_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	// Initially no git repo
	retByte, err := exec.Command("git", "status").CombinedOutput()
	ret := string(retByte)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\n" {
		t.Errorf("Bad state. Expected directory not to be a git repo, instead got: %s\n", ret)
	}

	// Initialize git repo
	retByte, err = exec.Command("git", "init").CombinedOutput()
	ret = string(retByte)
	if ret[:32] != "Initialized empty Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}

	// Create file to add
	f, err := os.Create("add-test")
	lib.Check(err)
	defer func() {
		f.Close()
		os.Remove("add-test")
	}()

	// Add file
	adding := []string{"add-test"}
	Add(adding)

	// Check added
	retByte, err = exec.Command("git", "status").CombinedOutput()
	ret = string(retByte)
	re := regexp.MustCompile(`new file:   add-test`)
	output := re.FindString(ret)
	if output == "" {
		t.Errorf("Expected status to have file added. Instead, got: %s", ret)
	}

	// Remove file and git
	err = os.Remove("add-test")
	lib.Check(err)
	err = os.RemoveAll(".git")
	lib.Check(err)

	// Initialize git repo
	retByte, err = exec.Command("git", "init").CombinedOutput()
	ret = string(retByte)
	if ret[:32] != "Initialized empty Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}

	// Make multiple files
	f1, err := os.Create("add-test1")
	lib.Check(err)
	f2, err := os.Create("add-test2")
	lib.Check(err)
	f3, err := os.Create("add-test3")
	lib.Check(err)
	defer func() {
		f1.Close()
		f2.Close()
		f3.Close()
		os.Remove("add-test1")
		os.Remove("add-test2")
		os.Remove("add-test3")
	}()

	// Add files
	adding = []string{"add-test1", "add-test2", "add-test3"}
	Add(adding)

	// Check added all
	retByte, err = exec.Command("git", "status").CombinedOutput()
	ret = string(retByte)
	re = regexp.MustCompile(`new file:   add-test1\s*new file:   add-test2\s*new file:   add-test3`)
	output = re.FindString(ret)
	if output == "" {
		t.Errorf("Expected status to have file added. Instead, got: %s", ret)
	}
}

func TestCommit(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_git_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	// Initially no git repo
	retByte, err := exec.Command("git", "status").CombinedOutput()
	ret := string(retByte)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\n" {
		t.Errorf("Bad state. Expected directory not to be a git repo, instead got: %s\n", ret)
	}

	// Initialize git repo
	retByte, err = exec.Command("git", "init").CombinedOutput()
	ret = string(retByte)
	if ret[:32] != "Initialized empty Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}

	// Create file to commit
	f, err := os.Create("commit-test")
	defer func() {
		f.Close()
		os.Remove("commit-test")
	}()
	lib.Check(err)

	// Add file
	err = exec.Command("git", "add", "commit-test").Run()
	lib.Check(err)

	// Commit file with message
	ret = lib.CaptureOutput(func() {
		Commit("Testing commit.")
	})

	re := regexp.MustCompile(`Testing commit.`)
	output := re.FindString(ret)
	if output == "" {
		t.Errorf("Expected commit message to go through. Instead, got: %s", ret)
	}
}

func TestIgnore(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_git_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	// Initially no git repo
	retByte, err := exec.Command("git", "status").CombinedOutput()
	ret := string(retByte)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\n" {
		t.Errorf("Bad state. Expected directory not to be a git repo, instead got: %s\n", ret)
	}

	// Initialize git repo
	retByte, err = exec.Command("git", "init").CombinedOutput()
	ret = string(retByte)
	if ret[:32] != "Initialized empty Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}

	// Create file to ignore and gitignore file
	f, err := os.Create("IgnoreMe")
	defer func() {
		f.Close()
		os.Remove("IgnoreMe")
	}()
	lib.Check(err)
	gitignore, err := os.Create(".gitignore")
	defer func() {
		gitignore.Close()
		os.Remove(".gitignore")
	}()
	lib.Check(err)

	// Check isn't ignored
	retByte, err = exec.Command("git", "status").CombinedOutput()
	ret = string(retByte)
	re := regexp.MustCompile(`IgnoreMe`)
	output := re.FindString(ret)
	if output == "" {
		t.Errorf("Expected to find IgnoreMe in git status. Instead got: %s", ret)
	}

	// Ignore the IgnoreMe file
	ignoring := []string{"IgnoreMe"}
	Ignore(ignoring)

	// Check is ignored
	retByte, err = exec.Command("git", "status").CombinedOutput()
	ret = string(retByte)
	re = regexp.MustCompile(`IgnoreMe`)
	output = re.FindString(ret)
	if output != "" {
		t.Errorf("Expected IgnoreMe to not be in git status. Instead got: %s", ret)
	}

	// Create another file to ignore
	f2, err := os.Create("IgnoreMeToo")
	defer func() {
		f2.Close()
		os.Remove("IgnoreMeToo")
	}()
	lib.Check(err)

	// Check isn't ignored
	retByte, err = exec.Command("git", "status").CombinedOutput()
	ret = string(retByte)
	re = regexp.MustCompile(`IgnoreMeToo`)
	output = re.FindString(ret)
	if output == "" {
		t.Errorf("Expected to find IgnoreMeToo in git status. Instead got: %s", ret)
	}

	// Ignore the IgnoreMeToo file and skip the IgnoreMe file
	ignoring = append(ignoring, "IgnoreMeToo")
	ret = lib.CaptureOutput(func() {
		Ignore(ignoring)
	})
	if ret != "Skipping IgnoreMe because it is already being ignored.\n" {
		t.Errorf("Expected to skip IgnoreMe because it's already being ignored. Instead got: %s", ret)
	}

	// Check the IgnoreMeToo is being ignored now
	retByte, err = exec.Command("git", "status").CombinedOutput()
	ret = string(retByte)
	re = regexp.MustCompile(`IgnoreMeToo`)
	output = re.FindString(ret)
	if output != "" {
		t.Errorf("Expected IgnoreMeToo to not be in git status. Instead got: %s", ret)
	}
}
