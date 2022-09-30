package classify

import (
	"errors"
	"fmt"

	"github.com/cohere-ai/cohere-go"
	"github.com/google/go-github/v47/github"
	"github.com/kfcampbell/cohere-demo/internal/pkg/issues"
)

// FIXME: use issue label to improve quality of classification examples
func CounterExampleFromIssueAndLabel(issue *github.Issue, label *github.Label) cohere.Example {
	return cohere.Example{
		Text:  issues.IssueText(issue),
		Label: fmt.Sprintf("not %s", label.GetName()),
	}
}

func ExampleFromIssueAndLabel(issue *github.Issue, label *github.Label) cohere.Example {
	return cohere.Example{
		Text:  issues.IssueText(issue),
		Label: label.GetName(),
	}
}

func NewClassifyAsLabelRequest(unlabelledIssue *github.Issue, label *github.Label, repositoryIssues []*github.Issue) (opts cohere.ClassifyOptions, err error) {
	var examples []cohere.Example

	otherIssues, err := issues.IssuesWithoutLabel(repositoryIssues, label, maxOtherIssues)
	if err != nil {
		return opts, err
	}

	labelledIssues, err := issues.IssuesForLabel(repositoryIssues, label, maxLabelledIssues)
	if len(labelledIssues) == 0 {
		return opts, ErrNoIssuesForLabel
	}

	if err != nil {
		return opts, err
	}

	otherIssues = issues.ShuffleIssues(otherIssues)
	for _, issue := range otherIssues {
		examples = append(examples, CounterExampleFromIssueAndLabel(issue, label))
	}

	for _, issue := range labelledIssues {
		examples = append(examples, ExampleFromIssueAndLabel(issue, label))
	}

	opts = cohere.ClassifyOptions{
		Model:           defaultModel,
		TaskDescription: fmt.Sprintf("Classify an unlabelled issue as %s or not %[1]s", label.GetName()),
		Inputs:          []string{issues.IssueText(unlabelledIssue)},
		Examples:        examples,
	}

	return opts, err
}

const defaultModel string = `large`
const maxLabelledIssues int = 40
const maxOtherIssues int = 10

var ErrNoIssuesForLabel error = errors.New("failed to find any issues with matching label")
