/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kfcampbell/cohere-demo/internal/pkg/env"
	"github.com/kfcampbell/cohere-demo/internal/pkg/labels"
	"github.com/spf13/cobra"
)

// labelsCmd represents the labels command
var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "list all labels associated with a repository",
	Long:  "list all labels associated with a repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("labels called")

		nwo := env.RepositoryNWO()
		repositoryLabels, err := labels.NewRepositoryLabels(nwo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("found %v labels\n", len(repositoryLabels))
		for _, label := range repositoryLabels {
			fmt.Println(label.GetName())
		}
	},
}

func init() {
	rootCmd.AddCommand(labelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
