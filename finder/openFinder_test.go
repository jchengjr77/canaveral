package finder

import (
	"errors"
	"os/user"
	"testing"

	"github.com/jchengjr77/canaveral/lib"
)

func TestOpenFinder(t *testing.T) {
	// Redirect testing output
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
	actual := OpenFinder("", wsPath)
	if expect.Error() != actual.Error() {
		t.Errorf("Empty project name did not yield correct error")
		t.Errorf("expected: %s\n", expect)
		t.Errorf("actual: %s\n", actual)
		return
	}
	expect = errors.New("No canaveral workspace set")
	actual = OpenFinder("FakeProject", wsPath)
	if expect.Error() != actual.Error() {
		t.Errorf("No workspace set did not yield correct error")
		t.Errorf("expected: %s\n", expect)
		t.Errorf("actual: %s\n", actual)
		return
	}
}
