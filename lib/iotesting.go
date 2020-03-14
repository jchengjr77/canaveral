//Package lib contains testing functions for canaveral
// includes:
// - capture output (from stdout)
// - redirect stdout
// - reset stdout
package lib

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

// CaptureOutput takes in a function and reads all print statements.
// Code snippet taken from:
// https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
func CaptureOutput(f func()) string {
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

// RedirOut redirects standard out away, and returns original *file.
// This is mainly so whilst testing, the console doesn't get flooded.
func RedirOut() *os.File {
	_, writer, _ := os.Pipe()
	realStdout := os.Stdout
	os.Stdout = writer // redirect output away
	return realStdout
}

// ResetOut resets os.stdout to the original stdout
func ResetOut(stdout *os.File) {
	tempstdOut := os.Stdout
	os.Stdout = stdout
	tempstdOut.Close()
}
