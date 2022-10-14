/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kfcampbell/cohere-demo/internal/pkg/auth"
	"github.com/kfcampbell/cohere-demo/internal/pkg/classify"
	"github.com/kfcampbell/cohere-demo/internal/pkg/env"
	"github.com/kfcampbell/cohere-demo/internal/pkg/issues"
	"github.com/kfcampbell/cohere-demo/internal/pkg/labels"
	"github.com/spf13/cobra"
)

// classifyCmd represents the classify command
var classifyCmd = &cobra.Command{
	Use:   "classify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("classify called")

		nwo := env.RepositoryNWO()

		repositoryIssues, err := issues.NewRepositoryIssues(nwo)
		if err != nil {
			panic(err)
		}

		unlabelledIssues := issues.UnlabelledIssues(repositoryIssues)

		repositoryLabels, err := labels.NewRepositoryLabels(nwo)
		if err != nil {
			panic(err)
		}

		cohereClient, err := auth.NewCohereClient()
		if err != nil {
			panic(err)
		}

		for _, unlabelledIssue := range unlabelledIssues {
			fmt.Println("-------------")
			fmt.Printf("Title: %s\n", unlabelledIssue.GetTitle())
			var predictedLabels []string
		nextLabel:
			for _, label := range repositoryLabels {

				req, err := classify.NewClassifyAsLabelRequest(
					unlabelledIssue,
					label,
					repositoryIssues)

				if err != nil {

					// FIXME: excludes labels as classify targets
					if err == classify.ErrInsufficientIssuesForLabel {
						// fmt.Printf("WARN: unable to classify by label '%s'\n", *label.Name)
						continue nextLabel
					}

					panic(err)
				}

				resp, err := cohereClient.Classify(req)
				if err != nil {
					panic(err)
				}

				for i := range resp.Classifications {
					predictedLabels = append(predictedLabels, resp.Classifications[i].Prediction)
				}

			}

			if len(predictedLabels) == 0 {
				fmt.Println("No Predictions!")
			} else {
				for i := range predictedLabels {
					fmt.Printf("Prediction: %s\n", predictedLabels[i])
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(classifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// classifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// classifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
