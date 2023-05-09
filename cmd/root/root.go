/*
Copyright © 2023 Farzaan Shaikh fvshaikh93@gmail.com

Use of this source code is governed by a GPL
license that can be found in the LICENSE file.
*/
package root

import (
	"fmt"
	"os"

	initialize "github.com/farzaanshaikh/resume-manager-cli/cmd/init"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.resume-manager-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		wd, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".resume-manager-cli" (without extension).
		viper.AddConfigPath(wd)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".resume-manager-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
