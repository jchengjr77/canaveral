package main

import "testing"

// test configuration directory location
func TestConfDir(t *testing.T) {
	res := confDir
	if res != "/.canaveral/config/" {
		t.Errorf("confDir is '%s' --> should be '/.canaveral/config/'", confDir)
	}
}

func TestWsFName(t *testing.T) {
	res := wsFName
	if res != "cnavrlws" {
		t.Errorf("wsFName is '%s' --> should be 'cnavrlws'", wsFName)
	}
}

func TestTestFile(t *testing.T) {
	res := testFile
	if res != "canaveral_test_file" {
		t.Errorf("testFile is '%s' --> should be 'canaveral_test_file'",
			wsFName)
	}
}
