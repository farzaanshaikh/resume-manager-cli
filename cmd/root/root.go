/*
Copyright © 2023 Farzaan Shaikh

This code is licensed under the Apache License 2.0.
For more information, please see the LICENSE file.
*/
package root

import (
	"os"

	"github.com/farzaanshaikh/resume-manager-cli/cmd/config"
	initialize "github.com/farzaanshaikh/resume-manager-cli/cmd/init"
	newPkg "github.com/farzaanshaikh/resume-manager-cli/cmd/new"
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
	rootCmd.AddCommand(newPkg.NewCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .reman)")

	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Search config in current directory with name ".reman" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(cmdutil.DefaultConfigFileName)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()

}
