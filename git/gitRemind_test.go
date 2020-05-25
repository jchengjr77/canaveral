package git

import (
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"reflect"
	"testing"

	"github.com/jchengjr77/canaveral/lib"
)

func TestAddDeleteReminders(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_add_del_remind_ws"
	if !lib.DirExists(workingPath) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("test folder already exists")
		return
	}

	os.Chdir(workingPath)

	actual, err := loadReminders()
	lib.Check(err)
	intended := make(map[string]interface{})

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed when .remind.json is not present")
		return
	}

	lib.CreateFile(".remind.json")
	defer os.Remove(".remind.json")

	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed when .remind.json is empty")
		return
	}

	addReminder("test", "This is a test message")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed after calling addReminder")
		return
	}

	out := lib.CaptureOutput(func() {
		err = addReminder("test", "This is a test message")
		lib.Check(err)
	})
	if out != "\"This is a test message\" is already a stored reminder for this file\n" {
		t.Errorf("failed to recognize adding already added reminder")
		return
	}

	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to not add already added reminder")
		return
	}

	addReminder("test", "Another test")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message", "Another test"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to add second reminder")
		return
	}

	addReminder("test2", "Different file")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message", "Another test"}
	intended["test2"] = []interface{}{"Different file"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to add test to second file")
		return
	}

	DelReminder("test2", "Different file")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message", "Another test"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to delete entire file when deleting last reminder")
		return
	}

	content := []byte("y")
	tmpfile, err := ioutil.TempFile("", "example")
	lib.Check(err)

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		lib.Check(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		lib.Check(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	DelReminder("test", "")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to delete entire file with empty second arg")
		return
	}

	os.Stdin = oldStdin

	addReminder("test", "Message 1")
	addReminder("test", "Message 2")
	addReminder("test", "Message 3")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"Message 1", "Message 2", "Message 3"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to on add three test")
		return
	}

	out = lib.CaptureOutput(func() {
		DelReminder("test", "Message 4")
	})
	if out != "Couldn't find reminder \"Message 4\" for file test\n" {
		t.Errorf("failed to recognize missing reminder on delete")
		return
	}
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to not delete on bad reminder")
		return
	}

	DelReminder("test", "1")
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"Message 2", "Message 3"}
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to delete reminder 1")
		return
	}

	err = DelReminder("test", "-1")
	if err.Error() != "Reminder number must be positive" {
		t.Errorf("failed to recognize negative number")
		return
	}
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"Message 2", "Message 3"}
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("deleted with negative number")
		return
	}

	err = DelReminder("test", "3")
	if err.Error() != "test has fewer than 3 reminders stored." {
		t.Errorf("failed to recognize too big number")
		return
	}
	actual, err = loadReminders()
	lib.Check(err)
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"Message 2", "Message 3"}
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("deleted with too big number")
		return
	}

}

func TestWrappers(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_wrappers_remind_ws"
	if !lib.DirExists(workingPath) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("test folder already exists")
		return
	}

	os.Chdir(workingPath)

	err = AddReminder("test", "test message")
	if err.Error() != "Git reminders can only be added in git repos" {
		t.Errorf("failed to recognize no git repo")
		return
	}

	err = exec.Command("git", "init").Run()
	lib.Check(err)

	err = AddReminder("test", "test message")
	if err.Error() != "Cannot find test" {
		t.Errorf("failed to recognize bad filename")
		return
	}
	lib.CreateFile("test")
	err = AddReminder("test", "")
	if err.Error() != "reminder cannot be empty" {
		t.Errorf("failed to recognize empty reminder")
		return
	}

	AddReminder("test", "test message")
	if !lib.FileExists(".remind.json") {
		t.Errorf("failed to create .remind.json")
		return
	}
	if !lib.FileExists(".gitignore") {
		t.Errorf("failed to create .gitignore")
		return
	}
	f, _ := os.Open(".gitignore")
	defer f.Close()
	if !inFile(f, ".remind.json") {
		t.Errorf("failed to add remind to gitignore")
		return
	}

	err = CheckReminders("bad-test")
	if err.Error() != "Cannot find file bad-test" {
		t.Errorf("Check failed to recognize bad filename")
		return
	}
	os.Remove(".remind.json")
	err = CheckReminders("bad-test")
	if err.Error() != "You don't have any reminders stored" {
		t.Errorf("Check failed to recognize missing .remind.json")
		return
	}
}
