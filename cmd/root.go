/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cohere-ai/cohere-go"
	"github.com/google/go-github/v47/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type ClassifyExample struct {
	Text  string `json:"text"`
	Label string `json:"label"`
}

type ClassifyRequest struct {
	TaskDescription string            `json:"taskDescription,omitempty"`
	Inputs          []string          `json:"inputs"`
	Examples        []ClassifyExample `json:"examples,omitempty"`
	OutputIndicator string            `json:"outputIndicator,omitempty"`
	Options         []string          `json:"options,omitempty"`
	Model           string            `json:"model"`
}

func DefaultClassifyRequest(model string) cohere.ClassifyOptions {
	return cohere.ClassifyOptions{
		Model:           model,
		OutputIndicator: "Classify this movie review",
		TaskDescription: "Classify these movie reviews as positive reviews, negative reviews, or neutral reviews",
		Inputs:          []string{"this movie was great", "this movie was bad"},
		Examples: []cohere.Example{
			{Text: "I would not recommend this movie to my friends", Label: "negative review"},
			{Text: "we made it only a quarter way through before we stopped", Label: "negative review"},
			{Text: "I did not want to finish the movie", Label: "negative review"},
			{Text: "worst movie of all time", Label: "negative review"},
			{Text: "hate this movie", Label: "negative review"},
			{Text: "this movie lacked any originality or depth", Label: "neutral review"},
			{Text: "this movie was okay", Label: "neutral review"},
			{Text: "this movie was neither amazing or terrible", Label: "neutral review"},
			{Text: "I would not watch this movie again but it was not a waste of time", Label: "neutral review"},
			{Text: "this movie was nothing special", Label: "neutral review"},
			{Text: "I would watch this movie again", Label: "positive review"},
			{Text: "i liked this movie", Label: "positive review"},
			{Text: "this is my favourite movie", Label: "positive review"},
			{Text: "I would watch this movie again with my friends", Label: "positive review"},
		},
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cohere-demo",
	Short: "Retrieve issues from a given GitHub repository",
	Long:  `Retrieve issues from a given GitHub repository`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("GITHUB_TOKEN") == "" {
			return fmt.Errorf("GITHUB_TOKEN must be set in environment")
		}
		ghToken := os.Getenv("GITHUB_TOKEN")

		if os.Getenv("COHERE_TOKEN") == "" {
			return fmt.Errorf("COHERE_TOKEN must be set in environment")
		}
		cohereToken := os.Getenv("COHERE_TOKEN")

		cohereClient, err := cohere.CreateClient(cohereToken)
		if err != nil {
			return err
		}

		req := DefaultClassifyRequest("large")
		resp, err := cohereClient.Classify(cohere.ClassifyOptions(req))
		if err != nil {
			return err
		}
		fmt.Println(resp)

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghToken},
		)

		authedGhClient := oauth2.NewClient(ctx, ts)
		ghClient := github.NewClient(authedGhClient)

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
			pagedIssues, resp, err := ghClient.Issues.ListByRepo(ctx, "integrations", "terraform-provider-github", opts)
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
