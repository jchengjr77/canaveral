package lib

import (
	"fmt"
	"os"
	"testing"
)

func sample() {
	fmt.Println("hello1")
	fmt.Println("hello2")
	fmt.Println("hello3")
}
func TestCaptureOutput(t *testing.T) {
	out := CaptureOutput(sample)
	want := "hello1\nhello2\nhello3\n"
	if out != want {
		t.Errorf(
			"func CaptureOutput() returned '%s' should be '%s'", out, want)
	}
}

func TestRedirOut(t *testing.T) {
	origOut := RedirOut()
	defer func() {
		os.Stdout = origOut
	}()
	if os.Stdout == origOut {
		t.Error("func RedirOut() failed to generate new stdout file_d.")
	}
}

func TestResetOut(t *testing.T) {
	_, writer, _ := os.Pipe()
	origOut := os.Stdout
	os.Stdout = writer // redirect output away
	if os.Stdout == origOut {
		t.Error("stdout redirection messed up. There's something wrong in the test.")
	}
	ResetOut(origOut)
	if os.Stdout != origOut {
		t.Error("func ResetOut() failed to reset stdout to original.")
	}
}
