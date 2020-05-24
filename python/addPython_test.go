package python

import (
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"testing"

	"github.com/jchengjr77/canaveral/lib"
)

func TestCreateCondaEnv(t *testing.T) {
	testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_python_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	re, _ := regexp.Compile(testProjName)
	out, err := exec.Command("conda", "info", "--envs").CombinedOutput()
	if err != nil {
		t.Errorf("conda info failed with %s\n", err.Error())
		return
	}
	find := re.MatchString(string(out))
	if find {
		t.Errorf("bad state, conda environment already exists\n")
		return
	}

	err = createCondaEnv(testProjName)
	if err != nil {
		t.Errorf("createCondaEnv failed with error: %s\n", err.Error())
		return
	}
	re, _ = regexp.Compile(testProjName)
	out, err = exec.Command("conda", "info", "--envs").CombinedOutput()
	if err != nil {
		t.Errorf("conda info failed with %s\n", err.Error())
		return
	}
	find = re.MatchString(string(out))
	if !find {
		t.Errorf("conda failed to create environment\n")
		return
	}

	err = exec.Command("conda", "env", "remove", "-n", testProjName).Run()
	if err != nil {
		t.Errorf("conda remove failed with %s\n", err.Error())
		return
	}
}

func TestActivateAndSetupConda(t *testing.T) {
	testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_python_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)

	re, _ := regexp.Compile(testProjName)
	out, err := exec.Command("conda", "info", "--envs").CombinedOutput()
	if err != nil {
		t.Errorf("conda info failed with %s\n", err.Error())
		return
	}
	find := re.MatchString(string(out))
	if find {
		t.Errorf("bad state, conda environment already exists\n")
		return
	}

	err = createCondaEnv(testProjName)
	if err != nil {
		t.Errorf("createCondaEnv failed with error: %s\n", err.Error())
		return
	}

	err = activateAndSetupConda(testProjName)
	if err != nil {
		t.Errorf("activateAndSetupConda failed with error: %s\n", err.Error())
		return
	}

	byteBase, err := exec.Command("conda", "info", "--base").Output()
	base := string(byteBase[:len(byteBase)-1])
	if !lib.FileExists(base + "/envs/" + testProjName + "/bin/python") {
		t.Errorf("activateAndSetupConda failed to conda install python")
		return
	}
	if !lib.FileExists(base + "/envs/" + testProjName + "/bin/pip") {
		t.Errorf("activateAndSetupConda failed to conda install python")
		return
	}

	err = exec.Command("conda", "env", "remove", "-n", testProjName).Run()
	if err != nil {
		t.Errorf("conda remove failed with %s\n", err.Error())
		return
	}

}

func TestCreateInstallSh(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_python_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)
	if lib.FileExists("install_packages.sh") {
		t.Errorf("Bad State, install already exists\n")
		return
	}
	err = createInstallSh()
	if err != nil {
		t.Errorf("Create install failed with %s\n", err.Error())
		return
	}
	if !lib.FileExists("install_packages.sh") {
		t.Errorf("Create install packes failed to create file\n")
		return
	}
}

func TestCreateReadme(t *testing.T) {
	testProjName := "testproj"
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_python_ws"
	if _, err = os.Stat(workingPath); os.IsNotExist(err) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("canaveral folder already exists\n")
		return
	}
	err = os.Chdir(workingPath)
	lib.Check(err)
	if lib.FileExists("README.md") {
		t.Errorf("Bad State, readme already exists\n")
		return
	}
	err = createREADME(testProjName, false)
	if err != nil {
		t.Errorf("Create readme failed with %s\n", err.Error())
		return
	}
	if !lib.FileExists("README.md") {
		t.Errorf("Create readme failed to create a readme\n")
		return
	}
	err = os.Remove("README.md")
	if err != nil {
		t.Errorf("Remove readme (no conda) failed with %s\n", err.Error())
		return
	}
	err = createREADME(testProjName, false)
	if err != nil {
		t.Errorf("Create readme failed with %s\n", err.Error())
		return
	}
	if !lib.FileExists("README.md") {
		t.Errorf("Create readme (conda) failed to create a readme\n")
		return
	}
}

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
