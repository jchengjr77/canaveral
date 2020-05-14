package git

import (
	"canaveral/lib"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"reflect"
	"testing"
)

func TestAddDeleteReminders(t *testing.T) {
	usr, err := user.Current()
	lib.Check(err)
	home := usr.HomeDir
	workingPath := home + "/canaveral_remind_ws"
	if !lib.DirExists(workingPath) {
		os.MkdirAll(workingPath, os.ModePerm)
		defer os.RemoveAll(workingPath)
	} else {
		t.Errorf("test folder already exists")
		return
	}

	os.Chdir(workingPath)

	actual := loadReminders()
	intended := make(map[string]interface{})

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed when .remind.json is not present")
		return
	}

	lib.CreateFile(".remind.json")
	defer os.Remove(".remind.json")

	actual = loadReminders()
	intended = make(map[string]interface{})
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed when .remind.json is empty")
		return
	}

	addReminder("test", "This is a test message")
	actual = loadReminders()
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed after calling addReminder")
		return
	}

	out := lib.CaptureOutput(func() {
		addReminder("test", "This is a test message")
	})
	if out != "\"This is a test message\" is already a stored reminder for this file\n" {
		t.Errorf("failed to recognize adding already added reminder")
		return
	}

	actual = loadReminders()
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to not add already added reminder")
		return
	}

	addReminder("test", "Another test")
	actual = loadReminders()
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message", "Another test"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to add second reminder")
		return
	}

	addReminder("test2", "Different file")
	actual = loadReminders()
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"This is a test message", "Another test"}
	intended["test2"] = []interface{}{"Different file"}

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to add test to second file")
		return
	}

	DelReminder("test2", "Different file")
	actual = loadReminders()
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
	actual = loadReminders()
	intended = make(map[string]interface{})
	fmt.Println(actual)
	fmt.Println(intended)

	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("failed to delete entire file with empty second arg")
		return
	}

	os.Stdin = oldStdin

	addReminder("test", "Message 1")
	addReminder("test", "Message 2")
	addReminder("test", "Message 3")
	actual = loadReminders()
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
	actual = loadReminders()
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
	actual = loadReminders()
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
	actual = loadReminders()
	intended = make(map[string]interface{})
	intended["test"] = []interface{}{"Message 2", "Message 3"}
	if !reflect.DeepEqual(actual, intended) {
		t.Errorf("deleted with too big number")
		return
	}

}

func TestWrappers(t *testing.T) {

}
