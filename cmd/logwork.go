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
	"fmt"
	"log"
	"strconv"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
)

var (
	jiraKey      string
	timeSpentStr string
	description  string
)

// logworkCmd represents the logwork command
var logworkCmd = &cobra.Command{
	Use:   "logwork",
	Short: "A brief description of your command",
	Long:  `Add time (minutes) to a Jira ticket`,
	Run: func(cmd *cobra.Command, args []string) {

		logWork(jiraKey, timeSpentStr, description)
	},
}

func init() {
	rootCmd.AddCommand(logworkCmd)
	// logworkCmd.PersistentFlags().StringVarP(&jiraKey, "k", "", "JIRA issue key")
	// logworkCmd.PersistentFlags().StringVarP("time", "t", "", "Time spent on the issue (in minutes)")
	// logworkCmd.PersistentFlags().StringVarP("description", "d", "", "Description of the work log")
	logworkCmd.PersistentFlags().StringVarP(&jiraKey, "key", "k", "", "JIRA issue key")
	logworkCmd.PersistentFlags().StringVarP(&timeSpentStr, "timespent", "t", "", "Time spent on the issue (in minutes)")
	logworkCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "Description of the work log")
	logworkCmd.MarkPersistentFlagRequired("key")
	logworkCmd.MarkPersistentFlagRequired("timespent")
	logworkCmd.MarkPersistentFlagRequired("description")

	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logworkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logworkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logWork(key string, timeSpentStr string, description string) {

	// Parse the time spent argument into an integer
	jiraClient := jiraClient()
	timeSpent, err := strconv.Atoi(timeSpentStr)
	if err != nil {
		log.Fatalf("Invalid time spent argument: %s", timeSpentStr)
	}

	// user, _, _ := jiraClient.User.GetSelf()
	// print(user.Active)
	issue, _, err := jiraClient.Issue.Get(key, nil)

	// issue, _, err := (jiraClient).Issue.Get(key, nil)
	if err != nil {
		log.Fatalf("Failed to get JIRA issue %s: %v", key, err)
	} else {
		log.Println(fmt.Sprintf("Adding %s minutes to ticket %s - %s", strconv.Itoa(timeSpent), key, issue.Fields.Summary))
	}

	// Log work on the JIRA issue
	worklog := &jira.WorklogRecord{
		TimeSpentSeconds: timeSpent * 60,
		Comment:          description,
	}
	_, _, err = jiraClient.Issue.AddWorklogRecord(issue.ID, worklog)
	if err != nil {
		log.Fatalf("Failed to log work on JIRA issue %s: %v", key, err)
	}

	log.Println("Work logged successfully!")
}
