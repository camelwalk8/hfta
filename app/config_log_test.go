package main

import (
	"testing"
)

func Test_SetupLogging(t *testing.T) {

	if b, _ := SetupLogging("/home/auser/go/src/hfta/hfta-email-json/logs"); !b {
		t.Errorf("filled setuping log")
	}

}

func Test_SetupLoggingNoPathTest(t *testing.T) {

	if b, _ := SetupLogging(""); b {
		t.Errorf("filled test pass")
	}

}

func Test_InfoLogging(t *testing.T) {

	if b, _ := SetupLogging("/home/auser/go/src/hfta/hfta-email-json/logs"); !b {
		t.Errorf("filled setuping log")
	}

	IL.Println("Information Logging Test")

}

func Test_ErrorLogging(t *testing.T) {

	if b, _ := SetupLogging("/home/auser/go/src/hfta/hfta-email-json/logs"); !b {
		t.Errorf("filled setuping log")
	}

	EL.Println("Error Logging Test")

}

func Test_FatalLogging(t *testing.T) {

	if b, _ := SetupLogging("/home/auser/go/src/hfta/hfta-email-json/logs"); !b {
		t.Errorf("filled setuping log")
	}

	FL.Println("Fatal Logging Test")

}
