package store

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
)

const perm = 0755

func TestIsEmptyDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		createDir   string
		createFiles []string
		args        args
		wantOut     bool
		wantError   error
	}{
		{
			name:        "GeneralNotExists",
			createDir:   "",
			createFiles: []string{},
			args: args{
				path: "path/doesn't/exist",
			},
			wantOut:   false,
			wantError: fs.ErrNotExist,
		},
		{
			name:        "IsEmpty",
			createDir:   "check",
			createFiles: []string{},
			args: args{
				path: "check",
			},
			wantOut:   true,
			wantError: nil,
		},
		{
			name:        "IsEmptyMulti",
			createDir:   "check/here",
			createFiles: []string{},
			args: args{
				path: "check/here",
			},
			wantOut:   true,
			wantError: nil,
		},
		{
			name:        "NotEmptyHasFolder",
			createDir:   "check/here",
			createFiles: []string{},
			args: args{
				path: "check",
			},
			wantOut:   false,
			wantError: nil,
		},
		{
			name:        "NotEmptyHasFile",
			createDir:   "check",
			createFiles: []string{"check/test.txt", "check/check.yaml"},
			args: args{
				path: "check",
			},
			wantOut:   false,
			wantError: nil,
		},
		{
			name:        "NotEmptyHasHidden",
			createDir:   "check",
			createFiles: []string{"check/.hiddenFile"},
			args: args{
				path: "check",
			},
			wantOut:   false,
			wantError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := os.MkdirAll(tc.createDir, perm); err != nil && !os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			for _, fileName := range tc.createFiles {
				os.Create(fileName)
			}

			gotOut, gotError := IsEmptyDir(tc.args.path)

			if err := os.RemoveAll(strings.Split(tc.createDir, "/")[0]); err != nil && !os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if gotOut != tc.wantOut {
				t.Errorf("returned output '%t', expected '%t'", gotOut, tc.wantOut)
			}

			if !errors.Is(gotError, tc.wantError) {
				t.Errorf("returned error '%v', expected '%v'", gotError, tc.wantError)
			}
		})
	}

}

func TestDirExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name      string
		createDir string // DO NOT include sub-dirs
		args      args
		wantError error
	}{
		{
			name:      "nonExistentPathMulti",
			createDir: "",
			args: args{
				path: "/randompath/doesn't/exist/",
			},
			wantError: fs.ErrNotExist,
		},
		{
			name:      "existentPath",
			createDir: "check",
			args: args{
				path: "check",
			},
			wantError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := os.Mkdir(tc.createDir, perm); err != nil && !os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			gotError := DirExists(tc.args.path)

			if err := os.RemoveAll(tc.createDir); err != nil && !os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if !errors.Is(gotError, tc.wantError) {
				t.Errorf("returned error '%v', expected '%v'", gotError, tc.wantError)
			}
		})
	}

}
