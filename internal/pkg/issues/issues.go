package issues

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/go-github/v47/github"
	"github.com/kfcampbell/cohere-demo/internal/pkg/auth"
	"github.com/kfcampbell/cohere-demo/internal/pkg/env"
)

func ShuffleIssues(src []*github.Issue) []*github.Issue {
	final := make([]*github.Issue, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))
	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func IssueText(issue *github.Issue) (text string) {
	// abbreviatedBody := issue.GetBody()[0:maxIssueBodyChars]
	// text = fmt.Sprintf("%s:\n\n%s", issue.GetTitle(), abbreviatedBody)
	// text = fmt.Sprintf("%s:\n\n%s", issue.GetTitle(), abbreviatedBody)
	return issue.GetTitle()
}

func IssuesWithoutLabel(issues []*github.Issue, targetLabel *github.Label, limit int) (withoutLabelIssues []*github.Issue, err error) {
	for i, issue := range issues {
		matched := false
		for _, label := range issue.Labels {
			if targetLabel.Name == label.Name {
				matched = true
			}
		}
		if !matched {
			withoutLabelIssues = append(withoutLabelIssues, issue)
		}
		if i > limit {
			break
		}

	}

	return withoutLabelIssues, nil
}

func IssuesForLabel(issues []*github.Issue, targetLabel *github.Label, limit int) (labelIssues []*github.Issue, err error) {
	for i, issue := range issues {
		matched := false
		for _, label := range issue.Labels {
			if targetLabel.Name == label.Name {
				matched = true
			}
		}
		if matched {
			labelIssues = append(labelIssues, issue)
		}
		if i > limit {
			break
		}
	}

	return labelIssues, nil
}

func UnlabelledIssues(issues []*github.Issue) (unlabelledIssues []*github.Issue) {
	for _, issue := range issues {
		if len(issue.Labels) != 0 {
			unlabelledIssues = append(unlabelledIssues, issue)
		}
	}
	return unlabelledIssues
}

func NewRepositoryIssues(nwo string) (issues []*github.Issue, err error) {

	ctx := context.Background()
	ghClient, err := auth.NewGitHubClient(ctx)
	if err != nil {
		return nil, err
	}

	repositoryOwner := env.RepositoryOwner()
	repositoryName := env.RepositoryName()

	opts := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{
			Page: 0,
		},
	}

	for {
		issueBatch, resp, err := ghClient.Issues.ListByRepo(ctx, repositoryOwner, repositoryName, opts)
		if err != nil {
			return nil, err
		}

		issues = append(issues, issueBatch...)
		if resp.NextPage == 0 {
			break
		} else {
			opts.ListOptions.Page = resp.NextPage
		}
	}

	return issues, nil
}

const maxIssueBodyChars = 1024
