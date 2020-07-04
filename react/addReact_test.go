package react

import (
	"os"
	"os/exec"
	"os/user"
	"testing"

	"github.com/jchengjr77/canaveral/lib"
)

func TestCheckCRAExists(t *testing.T) {
	craPath, err := setCRAPath()
	lib.Check(err)
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	err = os.MkdirAll(home+"/.canaveral", os.ModePerm)
	lib.Check(err)
	err = os.Chdir(home + "/.canaveral")
	lib.Check(err)
	uninCRA := exec.Command("npm", "uninstall", "create-react-app")
	uninCRA.Stderr = os.Stderr
	uninCRA.Stdout = os.Stdout
	err = uninCRA.Run()
	lib.Check(err)
	res := checkCRAExists(craPath)
	if res != false {
		t.Errorf(
			"func checkCRAExists() true when create-react-app is uninstalled")
	}
	inCRA := exec.Command("npm", "install", "create-react-app")
	inCRA.Stderr = os.Stderr
	inCRA.Stdout = os.Stdout
	err = inCRA.Run()
	lib.Check(err)
	out, err := exec.Command("ls node_modules/.bin").Output()
	t.Logf("%s\n", string(out))
	res = checkCRAExists(craPath)
	if res != true {
		t.Errorf(
			"func checkCRAExists() false when create-react-app is installed")
	}
}

func TestAddReactProj(t *testing.T) {
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
	AddReactProj(testProjName, wsPath)
	if !lib.DirExists(newPath + "/" + testProjName) {
		t.Errorf("func AddReactProj() failed to create ws at path: %s\n",
			newPath+"/"+testProjName)
		return
	}
	os.RemoveAll(newPath)
}
