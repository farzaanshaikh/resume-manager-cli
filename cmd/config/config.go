/*
Copyright © 2023 Farzaan Shaikh fvshaikh93@gmail.com

Use of this source code is governed by a GPL
license that can be found in the LICENSE file.
*/
package config

import (
	"errors"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the CLI",
	Long: `Adds configuration for the CLI.
These configs are directory level and not system-wide. Changing config
file directly might break things, best use this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(errors.New("not implemented yet"))
	},
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
