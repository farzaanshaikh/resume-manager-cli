/*
Copyright © 2023 Farzaan Shaikh fvshaikh93@gmail.com

Use of this source code is governed by a GPL
license that can be found in the LICENSE file.
*/
package initialize

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/farzaanshaikh/resume-manager-cli/store"
	"github.com/spf13/cobra"
)

var path string

// InitCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init [flags]",
	Short: "Initialize a resume store",
	Long: `Used to initialize a new resume store in a directory.

A store contains the following directories:
Your-Folder/
├── src/ 			# Source directory, open editor here
│   ├── first_resume.tex	# Latex files are in src/ folder
│   ├── custom/			# All support files (.sty, .cls) go here
│   ├── outputs/		# Render outputs such as log and axiliary
│   └── templates/		# Templates are resume you store for reuse
├── preview/			# Save multiple previews before you finalize
│   └── first_resume_p1.pdf	
└── Resume/			# Finalized files go here
    └── first_resume_p1.pdf`,

	Run: func(cmd *cobra.Command, args []string) {
		if err := initDir(cmd, args); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func initDir(cmd *cobra.Command, args []string) error {
	if path == "" {
		if newPath, err := os.Getwd(); err != nil {
			cobra.CheckErr(err)
		} else {
			path = newPath
		}
	}

	// Verify path exists
	if err := store.DirExists(path); err != nil {
		cobra.CheckErr(err)
	}

	// Warning if directory is not empty
	if isEmpty, err := store.IsEmptyDir(path); !isEmpty {
		var choice string
		fmt.Fprint(os.Stdout, "Warning: Directory not empty, do you wish to continue? (y or n) ")
		fmt.Fscanf(os.Stdin, "%s", &choice)
		if choice == "n" || choice == "N" {
			fmt.Fprintln(os.Stderr, "Aborted.")
			os.Exit(1)
		}
	} else if err != nil {
		cobra.CheckErr(err)
	}

	if err := createDirs(path); err != nil {
		cobra.CheckErr(err)
	}

	return nil
}

// Func to create the derectories specified in the init command
func createDirs(path string) error {
	// List of dir to create
	toCreate := []string{
		filepath.Join(path, store.SrcSubPath(store.Custom)),
		filepath.Join(path, store.SrcSubPath(store.Templates)),
		filepath.Join(path, store.SrcSubPath(store.Output)),
		filepath.Join(path, store.Resume),
		filepath.Join(path, store.Preview),
	}

	// check command by creating src dir
	if err := os.Mkdir(filepath.Join(path, store.Src), store.Perm); err != nil && !os.IsExist(err) {
		return errors.New("cannot create dir, check permissions")
	}

	for _, dir := range toCreate {
		err := os.MkdirAll(dir, store.Perm)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	return nil
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	InitCmd.Flags().StringVarP(&path, "dir", "d", "", "Specify directory to initialize the resume store")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
