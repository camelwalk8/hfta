package main

import (
	"bufio"
	"os"
	"strings"
)

func GetReleaseNotes(pathToReleaseNotesYmlFile string) (string, error) {
	if b, err := Exists(pathToReleaseNotesYmlFile); !b {
		return "", err
	}

	f, err := os.OpenFile(pathToReleaseNotesYmlFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var r strings.Builder

	sc := bufio.NewScanner(f)
	s := false
	e := false
	for sc.Scan() {
		if s && e {
			break
		}
		t := sc.Text()
		if t == "---START---" {
			r.WriteString("-------\n")
			s = true
			continue
		}

		if t == "---END---" {
			r.WriteString("-------\n\n")
			e = true
			continue
		}

		r.WriteString(t)
		r.WriteString("\n")
	}

	return r.String(), nil

}
