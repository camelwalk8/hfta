package main

import (
	"fmt"
	"os"
)

func Exists(pathToCheck string) (bool, error) {
	if _, err := os.Stat(pathToCheck); err != nil {
		if !os.IsNotExist(err) {
			return false, fmt.Errorf("error checking %q existence: %w", pathToCheck, err)
		}

		return false, fmt.Errorf("%q does not exist", pathToCheck)
	} else {
		return true, nil
	}
}
