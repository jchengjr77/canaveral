package lib

import (
	"errors"
	"os"
	"os/user"
	"testing"
)

const testFile = "canaveral_test_file"

func checkPanic(t *testing.T, f func()) {
	defer func() bool {
		if r := recover(); r == nil {
			t.Log("The code did not panic")
			return false
		}
		t.Log("The code panicked")
		return true
	}()
	f()
}

func TestCheck(t *testing.T) {
	var err1 error = nil
	t.Log("Following check should not panic")
	checkPanic(t, func() {
		Check(err1)
	})

	var err2 error = errors.New("Test Error")
	t.Log("Following check should panic")
	checkPanic(t, func() {
		Check(err2)
	})
}

func TestFileExists(t *testing.T) {
	tempusr, err := user.Current()
	Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_dir/"
	if FileExists(newPath+testFile) == true {
		t.Errorf("function FileExists() found a non-existent file at: %s",
			tempHome+testFile)
		t.Errorf("Check that no file exists there already.")
	}

	err = os.MkdirAll(newPath, os.ModePerm)
	Check(err)
	f, err := os.Create(newPath + testFile)
	Check(err)
	defer os.RemoveAll(newPath)
	defer f.Close()
	if FileExists(newPath+testFile) == false {
		t.Errorf("function FileExists() failed to recognize file at: %s",
			tempHome+testFile)
	}

}

func TestDirExists(t *testing.T) {
	tempusr, err := user.Current()
	Check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_dir/"
	if DirExists(newPath) == true {
		t.Errorf("function DirExists() found a non-existent dir at: %s",
			newPath)
		t.Errorf("Check that no dir exists there already.")
	}

	err = os.MkdirAll(newPath, os.ModePerm)
	Check(err)
	defer os.RemoveAll(newPath)
	if DirExists(newPath) == false {
		t.Errorf("function DirExists() failed to recognize dir at: %s",
			newPath)
	}
}

func TestCheckToolExists(t *testing.T) {
	actual := CheckToolExists("")
	expect := false
	if expect != actual {
		t.Errorf("func CheckToolExists() did not fail given empty toolname")
	}
	actual = CheckToolExists("shouldnotexist")
	expect = false
	if expect != actual {
		t.Errorf("func CheckToolExists() did not fail given fake toolname")
	}
	actual = CheckToolExists("--passingAnOption")
	expect = false
	if expect != actual {
		t.Fatalf("func CheckToolExists() did not fail given bad parameter")
	}
	actual = CheckToolExists("which") // runs 'which which'
	expect = true
	if expect != actual {
		t.Errorf("func CheckToolExists() did not find built-in command 'which'")
	}
}
