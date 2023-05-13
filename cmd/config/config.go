/*
Copyright © 2023 Farzaan Shaikh fvshaikh93@gmail.com

Use of this source code is governed by a GPL
license that can be found in the LICENSE file.
*/
package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

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

	// Readers and writers of different colors
	reader := bufio.NewReader(os.Stdin)
	bWriter := color.New(color.Bold)
	gWriter := color.New(color.Faint)

	bWriter.Fprint(os.Stdout, "\nAuthor info\n")

	// Get author.name
	fmt.Fprint(os.Stdout, "Set name of author")
	currName := viper.Get(cmdutil.AuthorNameKey)
	if currName != nil {
		gWriter.Fprint(os.Stdout, "(", currName, ")")
	}
	fmt.Fprint(os.Stdout, ": ")
	authorName, _ := reader.ReadString('\n')
	authorName = strings.TrimRight(authorName, "\n")
	authorName = strings.TrimLeft(authorName, " ")

	if authorName == "" && currName == nil {
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

	if !isValidString(authorName) {
		return errors.New("use of special characters or spaces")
	}

	return nil
}

func isValidString(s string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return regex.MatchString(s)
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
