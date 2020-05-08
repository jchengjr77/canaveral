package vscodesupport

import (
	"canaveral/lib"
	"errors"
	"os/user"
	"testing"
)

func TestOpenCode(t *testing.T) {
	origOut := lib.RedirOut()
	defer func() {
		lib.ResetOut(origOut)
	}()
	// Set home directory path of current user
	tempusr, err := user.Current()
	lib.Check(err)
	tempHome := tempusr.HomeDir
	wsPath := tempHome + "/tmpcnavrlws"
	// NOT writing a workspace path

	// no project name
	expect := errors.New("No project name specified")
	actual := OpenCode("", wsPath)
	if expect.Error() != actual.Error() {
		t.Errorf("Empty project name did not yield correct error")
		t.Errorf("actual: %s\n", actual)
		return
	}
	expect = errors.New("No canaveral workspace set")
	actual = OpenCode("FakeProject", wsPath)
	if expect.Error() != actual.Error() {
		t.Errorf("No workspace set did not yield correct error")
		return
	}
}
