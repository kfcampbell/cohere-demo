/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kfcampbell/cohere-demo/internal/pkg/env"
	"github.com/kfcampbell/cohere-demo/internal/pkg/issues"
	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issues called")

		nwo := env.RepositoryNWO()
		repositoryIssues, err := issues.NewRepositoryIssues(nwo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("found %v issues\n", len(repositoryIssues))
		for _, issue := range repositoryIssues {
			fmt.Println(issue.GetTitle())
		}

		unlabelledIssues := issues.UnlabelledIssues(repositoryIssues)
		fmt.Printf("found %v unlabelled issues\n", len(unlabelledIssues))
		for _, issue := range unlabelledIssues {
			fmt.Println(issue.GetTitle())
		}

	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
