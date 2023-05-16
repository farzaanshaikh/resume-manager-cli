/*
Copyright © 2023 Farzaan Shaikh

This code is licensed under the Apache License 2.0.
For more information, please see the LICENSE file.
*/
package store

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	Src     = "src"
	Preview = "preview"
	Resume  = "Resume"
	Perm    = 0755   // Default permission for the dirs
	TexExt  = ".tex" // Extension for latex files (.tex)
)

// Subdirectory within "src/".
//
// Use with SrcSubPath()
const (
	Custom    = "custom"
	Templates = "templates"
	Output    = "outputs"
)

// Uses "path/filepath" for compatibility
func SrcSubPath(name string) string {

	if name != Custom && name != Templates && name != Output {
		cobra.CheckErr(errors.New("incorrect filename passed"))
	}

	return filepath.Join(Src, name)
}

func IsEmptyDir(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Check if the path is a directory
	fi, err := f.Stat()
	if err != nil {
		return false, err
	}

	if !fi.IsDir() {
		return false, nil
	}

	// Checks for at least one file
	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func DirExists(path string) error {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return err
	}

	return nil
}
