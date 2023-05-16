/*
Copyright © 2023 Farzaan Shaikh

This code is licensed under the Apache License 2.0.
For more information, please see the LICENSE file.
*/
package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/farzaanshaikh/resume-manager-cli/cmdutil"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the CLI",
	Long: `Adds configuration for the CLI.
These configs are directory level and not system-wide. Changing config
file directly might break things, best use this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.Get(cmdutil.VersionKey) == nil {
			cobra.CheckErr(errors.New("valid config file not found, try running 'reman init'"))
		}

		if err := authorConfig(); err != nil {
			cobra.CheckErr(err)
		}

		viper.WriteConfig()
	},
}

func authorConfig() error {

	// Bold writer
	bWriter := color.New(color.Bold)
	bWriter.Fprint(os.Stdout, "\nAuthor Info\n")

	// Get author.name
	currName, _ := viper.Get(cmdutil.AuthorNameKey).(string)
	p := cmdutil.Prompter{
		Question:     "Name of author",
		DefaultValue: currName,
	}
	authorName := p.Input()

	if authorName == "" && currName == "" {
		return errors.New("must have a name")
	} else if authorName != "" {
		if err := validateAuthorName(authorName); err != nil {
			return err
		}
		viper.Set(cmdutil.AuthorNameKey, authorName)
	}

	return nil
}

// func to check author name length and special characters
func validateAuthorName(authorName string) error {
	if len(authorName) > cmdutil.AuthorNameLimit {
		msg := fmt.Sprintf("name has a max length of %d", cmdutil.AuthorNameLimit)
		return errors.New(msg)
	}

	if err := cmdutil.IsValidNameString(authorName); err != nil {
		return err
	}

	return nil
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
