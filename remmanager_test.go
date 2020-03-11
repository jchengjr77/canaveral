package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/user"
	"sync"
	"testing"
)

// captureOutput takes in a function and reads all print statements.
// Code snippet taken from:
// https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestConfirmDelete(t *testing.T) {
	var stdin bytes.Buffer // testable io
	stdin.WriteByte(byte('y'))
	res := confirmDelete("testProj", &stdin)
	if !res {
		t.Errorf("func confirmDelete() did not return true when fed 'y'")
	}
	stdin.WriteByte(byte('n'))
	res = confirmDelete("testProj", &stdin)
	if res {
		t.Errorf("func confirmDelete() did not return false when fed 'n'")
	}
	stdin.Write([]byte("foo"))
	res = confirmDelete("testProj", &stdin)
	if res {
		t.Errorf("func confirmDelete() did not return false when fed 'foo'")
	}
}

func TestTryRemProj(t *testing.T) {
	tempusr, err := user.Current()
	check(err)
	tempHome := tempusr.HomeDir
	newPath := tempHome + "/canaveral_test_ws/"
	err = os.MkdirAll(newPath, os.ModePerm)
	check(err)
	f, err := os.Create(newPath + "testProj")
	defer os.RemoveAll(newPath)
	defer f.Close()
	wsF, err := os.Create(tempHome + "/tempWSPath")
	defer os.Remove(tempHome + "/tempWSPath")
	defer wsF.Close()
	wsF.Write([]byte(newPath))
	res := tryRemProj("testProjFoo", tempHome+"/tempWSPath") // should be false
	if res {
		t.Logf("Path: %s\n", newPath)
		t.Errorf("func tryRemProj() returned true. Should be false.")
	}
	// omitted success case, as it is partially tested by confirmDelete
}

func TestRemProjectHandler(t *testing.T) {
	res := captureOutput(func() {
		remProjectHandler("")
	})
	if res !=
		"Cannot remove an unspecified project. Please provide the project name.\n" {
		t.Logf("remProjectHandler('') output: %s\n", res)
		t.Error("func remProjectHandler() failed in case of ''")
	}
	// omitted case 2 and 3, as they are tested in the aggregate
}
