/*
Copyright © 2023 Farzaan Shaikh fvshaikh93@gmail.com

Use of this source code is governed by a GPL
license that can be found in the LICENSE file.
*/
package root

import (
	"fmt"
	"os"

	"github.com/farzaanshaikh/resume-manager-cli/cmd/config"
	initialize "github.com/farzaanshaikh/resume-manager-cli/cmd/init"
	"github.com/farzaanshaikh/resume-manager-cli/cmdutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "reman [flags]",
	Short: "Resume Manager CLI",
	Long:  `Resume Manager is a CLI tool for your ever-changing resume needs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func addSubCommandPalettes() {
	rootCmd.AddCommand(initialize.InitCmd)
	rootCmd.AddCommand(config.ConfigCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .reman)")

	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if wd, err := os.Getwd(); err != nil {
		// Search config in current directory with name ".reman" (without extension).
		viper.AddConfigPath(wd)
		viper.SetConfigType("yaml")
		viper.SetConfigName(cmdutil.ConfigDefaultName)
	} else {
		cobra.CheckErr(err)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
