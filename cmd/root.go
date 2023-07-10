// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/twitchdev/twitch-cli/internal/util"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "twitch",
	Short: "A simple CLI tool for the New Twitch API and Webhook products.",
}
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Set the path to your .twitch-cli.env file",
	Run:   envCommandRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("line 26")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root.go line 33")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(envCmd)
	fmt.Println("root.go line 35")
	cfgFile, err := util.GetConfigPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("root.go line 41")
	envCmd.Flags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %s)", cfgFile))

}

func envCommandRun(cmd *cobra.Command, args []string) {
	var err error
	if customCfgFile == "" {
		fmt.Println("file is empty")
		customCfgFilePrompt := promptui.Prompt{
			Label: "Custom .twitch-cli.env path",
			Validate: func(s string) error {
				if len(s) > 4 {
					return nil
				}
				return errors.New("invalid length for custom file. must be .env")
			},
		}

		cfgFile, err = customCfgFilePrompt.Run()
	}
	fmt.Println(cfgFile)
	configPath := userPath(cfgFile, err)
	viper.Set("config", configPath)
	if err := viper.WriteConfigAs(configPath); err != nil {
		log.Fatalf("Failed to write configuration: %v", err.Error())
	}

	fmt.Println("Updated configuration.")
	return
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("root.go line 48")
	if cfgFile != "" {
		fmt.Println("cfgFile empty on root")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Println("root.go line 52")
	} else {
		// Find home directory.
		home, err := util.GetApplicationDir()
		fmt.Println("root.go line 56")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("root.go line 61")
		// Search config in home directory with name ".twitch-cli" (without extension).
		viper.AddConfigPath(home)
		fmt.Println("root.go line 64")
		viper.SetConfigName(".twitch-cli")
		viper.SetConfigType("env")
		fmt.Println("root.go line 67")
	}
	fmt.Println("root.go line 69")
	viper.SetEnvPrefix("twitch")
	viper.AutomaticEnv() // read in environment variables that match
	fmt.Println("root.go line 72")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return
	}
}
func userPath(file string, err error) string {
	fmt.Println("path function")
	if file == "" {
		configPath, err := util.GetConfigPath()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("returning configPath")
		return configPath
	}
	fmt.Println("path function should be returning file")
	return file
}
