package git

import (
	"canaveral/lib"
	"encoding/json"
	"errors"
	"io/ioutil"
	"fmt"
	"reflect"
)

func addReminder(file, reminder string) {
	remindContents, err := ioutil.ReadFile(".remind.json")
	lib.Check(err)
	var reminders map[string]interface{}
	err = json.Unmarshal(remindContents, &reminders)
	if _, found := reminders[file]; !found {
		reminders[file] = []string{reminder}
	} else {
		stored := reflect.ValueOf(reminders[file])
		var new []string = make([]string, stored.Len()+1)
		for i := 0; i < stored.Len(); i++ {
			adding := fmt.Sprintf("%v", stored.Index(i))
			if adding == reminder {
				fmt.Printf("\"%s\" is already a stored reminder for this file\n", reminder)
				return
			}
			new[i] = fmt.Sprintf("%v", stored.Index(i))
		}
		new[stored.Len()] = reminder
		reminders[file] = new
		// reminders[file] = append(reminders[file], reminder)
	}
	jsonData, err := json.Marshal(reminders)
	// This is not that robust of a solution because it rewrites the entire file
	// ! Research ways to improve this
	err = ioutil.WriteFile(".remind.json", jsonData, 0644)
}

// AddReminder adds a reminder to be displayed back to the user when committing 
// to git
func AddReminder(file, reminder string) error {
	var err error
	if !lib.DirExists(".git") {
		return errors.New("Git reminders can only be added in git repos")
	}
	// ! Change this back to !
	if lib.FileExists(file) {
		return errors.New("Cannot find " + file)
	}
	if !lib.FileExists(".remind.json") {
		if !lib.FileExists(".gitignore") {
			err = lib.CreateFile(".gitignore")
			lib.Check(err)
		}
		err = lib.CreateFile(".remind.json")
		lib.Check(err)
		Ignore([]string{".remind.json"})
	}
	addReminder(file, reminder)
	return nil
}