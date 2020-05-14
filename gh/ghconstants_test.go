package github

import "testing"

func TestLabel(t *testing.T) {
	res := label
	if res != "github credentials" {
		t.Errorf("label is %s --> should be 'github credentials'", res)
	}
}

func TestUrl(t *testing.T) {
	res := url
	if res != "https://api.github.com" {
		t.Errorf("url is %s --> should be 'https://api.github.com'", res)
	}
}
