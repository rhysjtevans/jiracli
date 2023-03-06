/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"os"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile   string
	jiraURL      string
	jiraUsername string
	jiraPassword string
	cfgFile      string
	// jiraClient   *jira.Client
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A cli helper tool to help you ",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}

func jiraClient() *jira.Client {
	jiraURL = viper.GetString("jira_url")
	jiraUsername = viper.GetString("jira_username")
	jiraPassword = viper.GetString("jira_token")
	if jiraURL == "" || jiraUsername == "" || jiraPassword == "" {
		log.Fatal("JIRA configuration not set: jira.url, jira.username, and jira.password are required")
	}
	tp := jira.BasicAuthTransport{
		Username: jiraUsername,
		Password: jiraPassword,
	}
	// fmt.Println(jiraURL)
	jiraClient, _ := jira.NewClient(tp.Client(), jiraURL)
	return jiraClient
}

func init() {

	cobra.OnInitialize(initConfig)
	// Find home directory.
	viper.AutomaticEnv()

	// Initialize the JIRA client with authentication credentials from the configuration file

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, _ := os.UserHomeDir()
	// cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.SetConfigName("JIRA")
	viper.SetConfigType("env")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jiracli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {

	// 	// // Search config in home directory with name ".jiracli" (without extension).
	// 	// viper.AddConfigPath(home)
	// 	// viper.SetConfigType("yaml")
	// 	// viper.SetConfigName(".jiracli")
	// }

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }

}
