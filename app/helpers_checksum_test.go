package main

import (
	"fmt"
	"testing"
)

func Test_GenerateMD5(t *testing.T) {
	//f := "/home/auser/go/src/hfta/app/helpers_checksum.go"
	f := "/home/auser/test-10mb.bin"
	if md5, bl, err := GenerateMD5(f); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Printf("md5 for %v is %v and length in bytes is %d", f, md5, bl)
	}
}
