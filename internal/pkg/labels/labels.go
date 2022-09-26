package labels

import (
	"context"

	"github.com/google/go-github/v47/github"
	"github.com/kfcampbell/cohere-demo/internal/pkg/auth"
	"github.com/kfcampbell/cohere-demo/internal/pkg/env"
)

func NewRepositoryLabels(nwo string) (repositoryLabels []*github.Label, err error) {

	ctx := context.Background()
	ghClient, err := auth.NewGitHubClient(ctx)
	if err != nil {
		return nil, err
	}
	opts := &github.ListOptions{Page: 0}

	repositoryOwner := env.RepositoryOwner()
	repositoryName := env.RepositoryName()

	for {
		labels, resp, err := ghClient.Issues.ListLabels(ctx, repositoryOwner, repositoryName, opts)
		if err != nil {
			return nil, err
		}

		repositoryLabels = append(repositoryLabels, labels...)
		if resp.NextPage == 0 {
			break
		} else {
			opts.Page = resp.NextPage
		}
	}

	return repositoryLabels, nil
}
