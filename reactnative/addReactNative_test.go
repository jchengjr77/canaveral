package reactnative

import (
	"testing"
)

func TestCheckToolExists(t *testing.T) {
	actual := checkToolExists("")
	expect := false
	if expect != actual {
		t.Errorf("func checkToolExists() did not fail given empty toolname")
	}
	actual = checkToolExists("shouldnotexist")
	expect = false
	if expect != actual {
		t.Errorf("func checkToolExists() did not fail given fake toolname")
	}
	actual = checkToolExists("--passingAnOption")
	expect = false
	if expect != actual {
		t.Fatalf("func checkToolExists() did not fail given bad parameter")
	}
	actual = checkToolExists("which") // runs 'which which'
	expect = true
	if expect != actual {
		t.Errorf("func checkToolExists() did not find built-in command 'which'")
	}
}
