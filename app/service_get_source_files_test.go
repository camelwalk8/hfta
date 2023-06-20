package main

import (
	"fmt"
	"testing"
)

func Test_GetFilesFromSrcDir(t *testing.T) {
	s := SourceDir{
		SrcDir:       "/home/auser/tests/processed-data-logs/dam",
		FilterRegExp: []string{"^[0-9a-f]{12}4[0-9a-f]{3}[89ab][0-9a-f]{15}.md5", "^[a-z]{2,3}_[0-9a-f]{12}4[0-9a-f]{3}[89ab][0-9a-f]{15}_[0-9]{8}-[0-9]{4}.md5"},
	}

	fs, err := s.GetFiles()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(fs)
		fmt.Printf("We have %d files in %s \n", len(fs), s.SrcDir)
	}
}
