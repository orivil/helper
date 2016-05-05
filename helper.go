// Copyright 2016 orivil Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package helper provides some useful methods.
package helper
import (
	"os"
	"path/filepath"
	"io/ioutil"
)

// IsExist check if the file or dir exist
func IsExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func GetFirstSubDirs(dir string) (files []string, err error) {
	fsInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, fi := range fsInfo {
		if fi.IsDir() {
			files = append(files, filepath.Join(dir, fi.Name()))
		}
	}
	return files, nil
}

func GetAllSubDirs(dir string) (dirs []string, e error) {
	var walk filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	}

	err := filepath.Walk(dir, walk)
	if err != nil {
		e = err
		return
	}
	return
}