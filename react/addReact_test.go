package react

import (
	"canaveral/lib"
	"os"
	"os/exec"
	"os/user"
	"testing"
)

func TestCheckCRAExists(t *testing.T) {
	uninCRA := exec.Command("npm", "uninstall", "-g", "create-react-app")
	uninCRA.Stderr = os.Stderr
	uninCRA.Stdout = os.Stdout
	err := uninCRA.Run()
	check(err)
	res := checkCRAExists()
	if res != false {
		t.Errorf(
			"func checkCRAExists() true when create-react-app is uninstalled")
	}
	inCRA := exec.Command("npm", "install", "-g", "create-react-app")
	inCRA.Stderr = os.Stderr
	inCRA.Stdout = os.Stdout
	err = inCRA.Run()
	check(err)
	res = checkCRAExists()
	if res != true {
		t.Errorf(
			"func checkCRAExists() false when create-react-app is installed")
	}
}

func TestAddReactProj(t *testing.T) {
	origOut := lib.RedirOut()
	defer func() {
		lib.ResetOut(origOut)
	}()
	testProjName := "testproj"
	tempusr, err := user.Current()
	check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_ws"
	wsPath := tempHome + "/tmpcnavrlws"
	f, err := os.Create(wsPath)
	check(err)
	defer os.Remove(wsPath)
	defer f.Close()
	f.WriteString(newPath)
	AddReactProj(testProjName, wsPath)
	if !dirExists(newPath + "/" + testProjName) {
		t.Errorf("func AddReactProj() failed to create ws at path: %s\n",
			newPath+"/"+testProjName)
		return
	}
	os.RemoveAll(newPath)
}
