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
	err = os.Chdir(home + "/canaveral")

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

	// Remove git repo
	err = exec.Command("rm", "-rf", ".git").Run()
	lib.Check(err)
	ret = lib.CaptureOutput(Status)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\nexit status 128\n" {
		t.Errorf("Status failed after cleanup. Expected no repo to exist, instead got: %s", ret)
	}
}

func TestInit(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	err = os.Chdir(home + "/canaveral")

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

	// Reset state, remove repo
	err = exec.Command("rm", "-rf", ".git").Run()
	lib.Check(err)
}

func TestAdd(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	err = os.Chdir(home + "/canaveral")
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
	err = exec.Command("touch", "add-test").Run()
	lib.Check(err)

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
	err = exec.Command("rm", "add-test").Run()
	lib.Check(err)
	err = exec.Command("rm", "-rf", ".git").Run()
	lib.Check(err)

	// Initialize git repo
	retByte, err = exec.Command("git", "init").CombinedOutput()
	ret = string(retByte)
	if ret[:32] != "Initialized empty Git repository" {
		t.Errorf("Initialized failed. Expected an empty repo to be initialized, instead got: %s", ret)
	}

	// Make multiple files
	err = exec.Command("touch", "add-test1").Run()
	lib.Check(err)
	err = exec.Command("touch", "add-test2").Run()
	lib.Check(err)
	err = exec.Command("touch", "add-test3").Run()
	lib.Check(err)

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

	// Remove files and git
	err = exec.Command("rm", "add-test1", "add-test2", "add-test3").Run()
	lib.Check(err)
	err = exec.Command("rm", "-rf", ".git").Run()
	lib.Check(err)
}
