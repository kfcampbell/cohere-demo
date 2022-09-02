/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v47/github"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cohere-demo",
	Short: "Retrieve issues from a given GitHub repository",
	Long:  `Retrieve issues from a given GitHub repository`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: consider authenticating for rate limit purposes
		// initialize a GitHub client
		client := github.NewClient(nil)
		ctx := context.Background()

		// retrieve issues on hard-coded repo for now
		// TODO: find examples of associations between labels and issues as a seed corpus
		// TODO: find unlabeled issues
		// TODO: for each unlabeled issue, call cohere to get an appropriate issue
		opts := &github.IssueListByRepoOptions{}
		issues, _, err := client.Issues.ListByRepo(ctx, "integrations", "terraform-provider-github", opts)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(issues)
		titles := make([]string, len(issues))

		for _, issue := range issues {
			titles = append(titles, *issue.Title)
			fmt.Println(*issue.Title)
		}
		//fmt.Println(titles)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cohere-demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
