package python

import (
	"canaveral/lib"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"testing"
)

func TestAddPythonProj(t *testing.T) {
	testProjName := "testproj"
	tempusr, err := user.Current()
	lib.Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_python_test_ws"
	wsPath := tempHome + "/tmpcnavrlws_python"
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
	AddPythonProj(testProjName, wsPath)
	if !lib.DirExists(newPath + "/" + testProjName) {
		t.Errorf("func AddPythonProj() failed to create ws at path: %s\n",
			newPath+"/"+testProjName)
		return
	}
	if !lib.FileExists(newPath + "/" + testProjName + "/" + testProjName + ".py") {
		t.Errorf("failed to create main python file\n")
		return
	}
	if !lib.FileExists(newPath + "/" + testProjName + "/" + "install_packages.sh") {
		t.Errorf("failed to create install_packages.sh file\n")
		return
	}
	if !lib.FileExists(newPath + "/" + testProjName + "/" + "README.md") {
		t.Errorf("failed to create README file\n")
		return
	}
	re, _ := regexp.Compile(testProjName)
	out, err := exec.Command("conda", "info", "--envs").CombinedOutput()
	if err != nil {
		t.Errorf("conda info failed with %s\n", err.Error())
		return
	}
	find := re.MatchString(string(out))
	if !find {
		t.Errorf("conda failed to create environment\n")
		return
	}
	err = exec.Command("conda", "env", "remove", "-n", testProjName).Run()
	if err != nil {
		t.Errorf("conda remove failed with %s\n", err.Error())
		return
	}
	os.RemoveAll(newPath)
}
