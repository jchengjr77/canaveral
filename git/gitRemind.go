package git

import (
	"bufio"
	"canaveral/lib"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

// * tested
func loadReminders() (res map[string]interface{}, finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	if !lib.FileExists(".remind.json") {
		return make(map[string]interface{}), nil
	}
	remindContents, err := ioutil.ReadFile(".remind.json")
	lib.Check(err)
	var reminders map[string]interface{}
	err = json.Unmarshal(remindContents, &reminders)
	if err != nil && err.Error() == "unexpected end of JSON input" {
		// file is empty
		reminders = make(map[string]interface{})
	} else {
		lib.Check(err)
	}
	return reminders, nil
}

// * tested
func addReminder(file string, reminder string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	reminders, err := loadReminders()
	lib.Check(err)
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
	}
	jsonData, err := json.Marshal(reminders)
	lib.Check(err)
	fmt.Printf("Will remind you, \"%s\" when you commit %s\n", reminder, file)
	// This is not that robust of a solution because it rewrites the entire file
	// ! Research ways to improve this
	err = ioutil.WriteFile(".remind.json", jsonData, 0644)
	lib.Check(err)
	return nil
}

func confirmPrint(stdin io.Reader) (res bool, finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	fmt.Printf("You have reminders stored for this file. Would you like to display them? ('y' or 'n')> ")
	reader := bufio.NewReader(stdin)
	response, err := reader.ReadByte()
	lib.Check(err)
	return (response == 'y'), nil
}

// ? untested but low priority because just prints info + manual tests are good
func checkReminders(file string, forcePrint bool, reminders map[string]interface{}) (res bool, finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	if !lib.FileExists(".remind.json") {
		return false, nil
	}
	fetched, found := reminders[file]
	stored := reflect.ValueOf(fetched)
	print := false
	if !forcePrint {
		if found {
			print, err := confirmPrint(os.Stdin)
			lib.Check(err)
		}
	}
	if forcePrint || print {
		fmt.Printf("=====Printing reminders for %s=====\n", file)
		if !found || stored.Len() == 0 {
			fmt.Println("No reminders found for", file)
		} else {
			var reminder string
			for i := 0; i < stored.Len(); i++ {
				reminder = fmt.Sprintf("%v", stored.Index(i))
				fmt.Printf("Reminder (%d): %s\n", i+1, reminder)
			}
		}
	}
	return forcePrint || print, nil
}

// CheckReminders prints all of the reminders stored for file
// * tested
func CheckReminders(file string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	if !lib.FileExists(".remind.json") {
		return errors.New("You don't have any reminders stored")
	}
	if !lib.FileExists(file) {
		return errors.New("Cannot find file " + file)
	}
	arg, err := loadReminders()
	lib.Check(err)
	res, err := checkReminders(file, true, arg)
	lib.Check(err)
	return nil
}

// AddReminder adds a reminder to be displayed back to the user when committing
// to git
// * tested
func AddReminder(file, reminder string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	var err error
	if !lib.DirExists(".git") {
		return errors.New("Git reminders can only be added in git repos")
	}
	if !lib.FileExists(file) {
		return errors.New("Cannot find " + file)
	}
	if reminder == "" {
		return errors.New("reminder cannot be empty")
	}
	if !lib.FileExists(".remind.json") {
		if !lib.FileExists(".gitignore") {
			err = lib.CreateFile(".gitignore")
			lib.Check(err)
		}
		err = lib.CreateFile(".remind.json")
		lib.Check(err)
		err = Ignore([]string{".remind.json"}, "", "")
		lib.Check(err)
	}
	err = addReminder(file, reminder)
	lib.Check(err)
	return nil
}

// * tested
func confirmDeleteAll(file string, stdin io.Reader) (res bool, finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	fmt.Printf("Are you sure you want to delete all reminders for %s ('y' or 'n')>", file)
	reader := bufio.NewReader(stdin)
	response, err := reader.ReadByte()
	lib.Check(err)
	return (response == 'y'), nil
}

// DelReminder deletes the specified reminder
// * tested
func DelReminder(file, reminder string) (finalErr error) {
	// defer a recover function that returns the thrown error
	defer func() {
		if r := recover(); r != nil {
			finalErr = r.(error)
		}
	}()
	if !lib.FileExists(".remind.json") {
		return errors.New("No reminders found")
	}
	reminders, err := loadReminders()
	lib.Check(err)
	fetched, found := reminders[file]
	stored := reflect.ValueOf(fetched)

	if !found || stored.Len() == 0 {
		return errors.New("Couldn't find reminder: \"" + reminder + "\" for file " + file)
	}

	if reminder == "" {
		res, err := confirmDeleteAll(file, os.Stdin)
		lib.Check(err)
		if res {
			delete(reminders, file)
			jsonData, err := json.Marshal(reminders)
			// This is not that robust of a solution because it rewrites the entire file
			// ! Research ways to improve this
			err = ioutil.WriteFile(".remind.json", jsonData, 0644)
			lib.Check(err)
		}
		return nil
	}

	if val, err := strconv.Atoi(reminder); err == nil {
		if val <= 0 {
			return errors.New("Reminder number must be positive")
		}
		if val > stored.Len() {
			return errors.New(file + " has fewer than " + reminder + " reminders stored.")
		}
		reminder = fmt.Sprintf("%v", stored.Index(val-1))
	}

	var ret []string
	found = false

	for i := 0; i < stored.Len(); i++ {
		curr := fmt.Sprintf("%v", stored.Index(i))
		if curr != reminder {
			ret = append(ret, curr)
		} else {
			found = true
		}
	}
	if len(ret) == 0 {
		delete(reminders, file)
		jsonData, err := json.Marshal(reminders)
		// This is not that robust of a solution because it rewrites the entire file
		// ! Research ways to improve this
		err = ioutil.WriteFile(".remind.json", jsonData, 0644)
		lib.Check(err)
		return nil
	}
	reminders[file] = ret
	if !found {
		fmt.Printf("Couldn't find reminder \"%s\" for file %s\n", reminder, file)
		return nil
	}

	fmt.Printf("Deleting reminder \"%s\" from file %s\n", reminder, file)
	jsonData, err := json.Marshal(reminders)
	// This is not that robust of a solution because it rewrites the entire file
	// ! Research ways to improve this
	err = ioutil.WriteFile(".remind.json", jsonData, 0644)
	lib.Check(err)
	return nil
}
