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
	"golang.org/x/oauth2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cohere-demo",
	Short: "Retrieve issues from a given GitHub repository",
	Long:  `Retrieve issues from a given GitHub repository`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("GITHUB_TOKEN") == "" {
			return fmt.Errorf("GITHUB_TOKEN must be set in environment")
		}
		token := os.Getenv("GITHUB_TOKEN")
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)

		authedClient := oauth2.NewClient(ctx, ts)
		client := github.NewClient(authedClient)

		// retrieve issues on hard-coded repo for now
		// TODO: find examples of associations between labels and issues as a seed corpus
		// TODO: for each unlabeled issue, call cohere to get an appropriate issue
		opts := &github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{
				Page: 0,
			},
		}
		issues := make([]*github.Issue, 0)

		// TODO: results don't seem to be matching the GitHub UI's number of issues. fix that.
		done := false
		for {
			if done {
				break
			}
			pagedIssues, resp, err := client.Issues.ListByRepo(ctx, "integrations", "terraform-provider-github", opts)
			if err != nil {
				return err
			}

			issues = append(issues, pagedIssues...)
			if resp.NextPage == 0 {
				done = true
			} else {
				opts.ListOptions.Page = resp.NextPage
			}
		}

		fmt.Printf("found %v issues\n", len(issues))
		labeledIssues := make([]github.Issue, 0)
		unlabeledIssues := make([]github.Issue, 0)
		for _, issue := range issues {
			//fmt.Println(*issue.Title)
			if len(issue.Labels) > 0 {
				labeledIssues = append(labeledIssues, *issue)
			} else {
				unlabeledIssues = append(unlabeledIssues, *issue)
			}
		}
		fmt.Printf("found %v unlabeled issues and %v labeled issues\n", len(unlabeledIssues), len(labeledIssues))
		return nil
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
