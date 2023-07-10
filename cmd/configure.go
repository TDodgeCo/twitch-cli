// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/twitchdev/twitch-cli/internal/util"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var clientID string
var clientSecret string
var customCfgFile string

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures your Twitch CLI with your Client ID and Secret",
	Run:   configureCmdRun,
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVarP(&clientID, "client-id", "i", "", "Client ID to use.")
	configureCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "Client Secret to use.")
	configureCmd.Flags().StringVarP(&customCfgFile, "config", "c", "", "Config path to use.")
}

func configureCmdRun(cmd *cobra.Command, args []string) {
	var err error
	if clientID == "" {
		clientIDPrompt := promptui.Prompt{
			Label: "Client IDsss",
			Validate: func(s string) error {
				if len(s) == 30 || len(s) == 31 {
					return nil
				}
				return errors.New("Invalid length for Client ID")
			},
		}

		clientID, err = clientIDPrompt.Run()
	}

	if clientSecret == "" {
		clientSecretPrompt := promptui.Prompt{
			Label: "Client Secretsss",
			Validate: func(s string) error {
				if len(s) == 30 || len(s) == 31 {
					return nil
				}
				return errors.New("Invalid length for Client Secret")
			},
		}

		clientSecret, err = clientSecretPrompt.Run()
	}

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

	if clientID == "" && clientSecret == "" {
		fmt.Println("Must specify either the Client ID or Secret")
		return
	}
	fmt.Println(cfgFile)
	configPath := path(cfgFile, err)

	viper.Set("clientId", clientID)
	viper.Set("clientSecret", clientSecret)
	viper.Set("config", configPath)
	if err := viper.WriteConfigAs(configPath); err != nil {
		log.Fatalf("Failed to write configuration: %v", err.Error())
	}

	fmt.Println("Updated configuration.")
	return
}

func path(file string, err error) string {
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
