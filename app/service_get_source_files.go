package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type SourceDir struct {
	SrcDir       string
	FilterRegExp []string
}

func (a *SourceDir) GetFiles() ([]string, error) {
	if _, err := Exists(a.SrcDir); err != nil {
		return nil, fmt.Errorf("error walking source %q existence: %w", a.SrcDir, err)
	}

	fis := []os.FileInfo{}

	walk := func(path string, de fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if de.IsDir() {
			return nil
		}

		for _, e := range a.FilterRegExp {

			re := regexp.MustCompile(e)

			n := filepath.Base(path)

			if !re.MatchString(n) {
				continue
			}

			fi, err := de.Info()

			if err != nil {
				continue
			}

			if fi.IsDir() {
				continue
			}

			fis = append(fis, fi)
			break

		}
		return nil
	}

	filepath.WalkDir(a.SrcDir, walk)

	sort.Slice(fis, func(i, j int) bool {
		return fis[i].ModTime().Before(fis[j].ModTime())
	})

	files := func(f *[]os.FileInfo) []string {
		files := []string{}
		for _, fi := range *f {
			files = append(files, fmt.Sprintf("%s/%s", a.SrcDir, fi.Name()))
		}
		return files
	}(&fis)

	return files, nil

}
