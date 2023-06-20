package main

import (
	"testing"
)

func Test_PrintReleaseNotes(t *testing.T) {

	ymlconfig := "/home/auser/go/src/hfta/release_notes/hfta_release_notes.txt"

	if s, e := GetReleaseNotes(ymlconfig); e != nil {
		t.Fatal(e)
	} else {
		Debugl(s)
	}

}
