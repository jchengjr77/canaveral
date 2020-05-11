package git

import (
	"canaveral/lib"
	"os"
	"os/exec"
	"os/user"
	"testing"
)

func TestStatus(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	err = os.Chdir(home + "/canaveral")
	ret := lib.CaptureOutput(Status)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\nexit status 128\n" {
		t.Errorf("Bad state. Expected directory not to be a git repo, instead got: %s\n", ret)
	}
	err = exec.Command("git", "init").Run()
	lib.Check(err)
	ret = lib.CaptureOutput(Status)
	if ret[:16] != "On branch master" {
		t.Errorf("Expected git repo to exist, instead got: %s\n", ret)
	}
	err = exec.Command("rm", "-rf", ".git").Run()
	lib.Check(err)
	ret = lib.CaptureOutput(Status)
	if ret != "fatal: not a git repository (or any of the parent directories): .git\nexit status 128\n" {
		t.Errorf("Status failed after cleanup. Expected no repo to exist, instead got: %s", ret)
	}
}
