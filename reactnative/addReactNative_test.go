package reactnative

import (
	"bytes"
	"canaveral/lib"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"testing"
)

func TestInstallExpo(t *testing.T) {
	// make sure expo is installed
	hadNpmInitially := lib.CheckToolExists("npm")
	var actual, expect error
	actual = installExpo()
	if !hadNpmInitially {
		expect = errors.New("prerequisite tool 'npm' is not installed")
		if actual.Error() != expect.Error() {
			t.Errorf(
				"func installExpo() mismatched error (case of missing npm)")
		}
	} else {
		expect = nil
		if actual != expect {
			t.Errorf(
				"func installExpo() threw error when it wasn't supposed to")
		} else if !lib.CheckToolExists("expo") {
			t.Errorf(
				"func installExpo() did not properly install expo")
		}
	}
}

func TestConfirmInstall(t *testing.T) {
	origOut := lib.RedirOut()
	defer func() {
		lib.ResetOut(origOut) // reset it
	}()
	var stdin bytes.Buffer // testable io
	stdin.WriteByte(byte('y'))
	res, err := confirmInstall(&stdin)
	lib.Check(err)
	if !res {
		t.Errorf("func confirmInstall() did not return true when fed 'y'")
	}
	stdin.WriteByte(byte('n'))
	res, err = confirmInstall(&stdin)
	lib.Check(err)
	if res {
		t.Errorf("func confirmInstall() did not return false when fed 'n'")
	}
	stdin.Write([]byte("foo"))
	res, err = confirmInstall(&stdin)
	lib.Check(err)
	if res {
		t.Errorf("func confirmInstall() did not return false when fed 'foo'")
	}
}

func TestAddReactNativeProj(t *testing.T) {
	// make sure expo is installed
	hadExpoInitially := lib.CheckToolExists("expo")
	if !hadExpoInitially {
		expoIn := exec.Command("npm", "install", "-g", "expo-cli")
		expoOut := exec.Command("npm", "uninstall", "-g", "expo-cli")
		err := expoIn.Run()
		lib.Check(err)
		defer func() {
			err := expoOut.Run()
			lib.Check(err)
		}()
	}
	// Create a new react native project
	testProjName := "testproj"
	tempusr, err := user.Current()
	lib.Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_ws"
	wsPath := tempHome + "/tmpcnavrlws"
	f, err := os.Create(wsPath)
	lib.Check(err)
	defer func() {
		f.Close()
		os.Remove(wsPath)
	}()
	f.WriteString(newPath)
	err = os.Chdir("../")
	lib.Check(err)
	dir, err := os.Getwd()
	lib.Check(err)
	t.Logf("\nCurrent Dir: %s\n", dir)
	AddReactNativeProj(testProjName, wsPath)
	if !lib.DirExists(newPath + "/" + testProjName) {
		t.Errorf("func AddReactNativeProj() failed to create ws at path: %s\n",
			newPath+"/"+testProjName)
		return
	}
	os.RemoveAll(newPath)
}
