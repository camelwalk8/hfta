package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func GenerateMD5(pathToFile string) (string, int64, error) {
	if b, err := Exists(pathToFile); !b {
		if !os.IsNotExist(err) {
			return "", 0, fmt.Errorf("unable to calculate md5: %w", err)
		}
	}

	f, err := os.Open(pathToFile)
	if err != nil {
		return "", 0, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	md5 := md5.New()
	var bytesWritten int64
	if bytesWritten, err = io.Copy(md5, r); err != nil {
		return "", 0, err
	}

	return fmt.Sprintf("%x", md5.Sum(nil)), bytesWritten, nil
}
