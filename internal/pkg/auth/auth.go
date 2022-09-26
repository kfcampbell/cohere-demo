package auth

import (
	"context"
	"os"

	"github.com/cohere-ai/cohere-go"
	"github.com/google/go-github/v47/github"
	"github.com/kfcampbell/cohere-demo/internal/pkg/env"
	"golang.org/x/oauth2"
)

func NewGitHubClient(ctx context.Context) (gh *github.Client, err error) {
	if err = env.Valid(); err != nil {
		return nil, err
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	authedGhClient := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(authedGhClient)

	return ghClient, nil
}

func NewCohereClient() (co *cohere.Client, err error) {
	if err = env.Valid(); err != nil {
		return nil, err
	}

	cohereToken := os.Getenv("COHERE_TOKEN")
	cohereClient, err := cohere.CreateClient(cohereToken)
	if err != nil {
		return co, err
	}
	return cohereClient, nil
}
