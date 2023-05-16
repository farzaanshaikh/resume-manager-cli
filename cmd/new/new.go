/*
Copyright © 2023 Farzaan Shaikh

This code is licensed under the Apache License 2.0.
For more information, please see the LICENSE file.
*/
package newPkg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/farzaanshaikh/resume-manager-cli/cmdutil"
	"github.com/farzaanshaikh/resume-manager-cli/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new resume",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// check for vaild dir
		if viper.Get(cmdutil.VersionKey) == nil {
			cobra.CheckErr(errors.New("valid config file not found, try running 'reman init'"))
		}

		if err := newResume(); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func newResume() error {
	// Get name of resume
	p := cmdutil.Prompter{Question: "Name this resume"}
	resName := p.Input()

	if resName == "" {
		return errors.New("must have name")
	}
	if err := cmdutil.IsValidNameString(resName); err != nil {
		return err
	}

	// Promt for template
	p = cmdutil.Prompter{Question: "Would you like to use a template?"}
	useTemplate := p.Confirm()
	var choice string

	// If using template prompt to select and apply
	if useTemplate {
		dir := store.SrcSubPath(store.Templates)
		pattern := filepath.Join(dir, "*"+store.TexExt)
		// Check if templates exist
		templates, err := filepath.Glob(pattern)
		if err != nil {
			return err
		}
		if len(templates) == 0 {
			return errors.New("no templates found, please copy a latex file into src/templates to use it")
		}
		tempNames := []string{}
		for _, file := range templates {
			filename := filepath.Base(file)
			tempNames = append(tempNames, filename[:len(filename)-4])
		}

		p = cmdutil.Prompter{
			Question: "Select the template you wish to use",
			Options:  tempNames,
		}
		choice = p.Select()
	}

	return createResFile(resName, choice)
}

func createResFile(name string, template string) error {

	// Create the resume file
	var df *os.File
	df, err := os.Create(filepath.Join(store.Src, name+store.TexExt))
	if err != nil {
		return err
	}
	defer df.Close()
	fmt.Fprintf(os.Stdout, "Resume '%v' created\n", df.Name())

	// Apply a template if selected
	if template != "" {
		// open the source file
		sf, err := os.Open(filepath.Join(store.SrcSubPath(store.Templates), template+store.TexExt))
		if err != nil {
			return err
		}
		defer sf.Close()

		var size int64
		// copy the contents of the source file to the destination file
		if size, err = io.Copy(df, sf); err != nil {
			return err
		}

		fmt.Printf("Copied %v bytes\n", size)
	}

	return nil
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
