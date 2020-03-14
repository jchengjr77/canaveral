package react

import (
	"os"
	"os/exec"
	"os/user"
	"testing"
)

func TestCheckCRAExists(t *testing.T) {
	craPath := setCRAPath()
	usr, err := user.Current()
	check(err)
	home := usr.HomeDir
	err = os.Chdir(home + "/canaveral")
	uninCRA := exec.Command("npm", "uninstall", "create-react-app")
	uninCRA.Stderr = os.Stderr
	uninCRA.Stdout = os.Stdout
	err = uninCRA.Run()
	check(err)
	res := checkCRAExists(craPath)
	if res != false {
		t.Errorf(
			"func checkCRAExists() true when create-react-app is uninstalled")
	}
	inCRA := exec.Command("npm", "install", "create-react-app")
	inCRA.Stderr = os.Stderr
	inCRA.Stdout = os.Stdout
	err = inCRA.Run()
	check(err)
	res = checkCRAExists(craPath)
	if res != true {
		t.Errorf(
			"func checkCRAExists() false when create-react-app is installed")
	}
}

func TestAddReactProj(t *testing.T) {
	testProjName := "testproj"
	tempusr, err := user.Current()
	check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_ws"
	wsPath := tempHome + "/tmpcnavrlws"
	f, err := os.Create(wsPath)
	check(err)
	defer func() {
		f.Close()
		os.Remove(wsPath)
	}()
	f.WriteString(newPath)
	err = os.Chdir("../")
	check(err)
	dir, err := os.Getwd()
	check(err)
	t.Logf("\nCurrent Dir: %s\n", dir)
	AddReactProj(testProjName, wsPath)
	if !dirExists(newPath + "/" + testProjName) {
		t.Errorf("func AddReactProj() failed to create ws at path: %s\n",
			newPath+"/"+testProjName)
		return
	}
	os.RemoveAll(newPath)
}
