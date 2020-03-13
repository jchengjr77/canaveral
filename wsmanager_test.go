package main

import "testing"

func TestFormatProjects(t *testing.T) {
	rawString := `total 0
	drwxr-xr-x   6 jjcheng77  staff  192 Mar 13 15:33 .
	drwxr-xr-x@ 15 jjcheng77  staff  480 Mar 10 08:27 ..
	drwxr-xr-x@  4 jjcheng77  staff  128 Mar  8 10:43 moments
	drwxr-xr-x@ 13 jjcheng77  staff  416 Mar  8 14:40 mysite
	drwxr-xr-x   5 jjcheng77  staff  160 Dec 16 12:55 random
	drwxr-xr-x@  6 jjcheng77  staff  192 Dec 21 16:23 react-native`

	want := `	drwxr-xr-x@  4 jjcheng77  staff  128 Mar  8 10:43 moments
	drwxr-xr-x@ 13 jjcheng77  staff  416 Mar  8 14:40 mysite
	drwxr-xr-x   5 jjcheng77  staff  160 Dec 16 12:55 random
	drwxr-xr-x@  6 jjcheng77  staff  192 Dec 21 16:23 react-native`

	res := formatProjects(rawString)

	if res != want {
		t.Logf("Input: '%s'\n\n", rawString)
		t.Logf("Output: '%s'\n\n", res)
		t.Logf("Ideal: '%s'\n\n", want)
		t.Error("func formatProjects() did not remove lines correctly.")
	}
}
