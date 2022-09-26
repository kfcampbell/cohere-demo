package env

import (
	"fmt"
	"os"
	"strings"
)

func Valid() error {

	if os.Getenv("GITHUB_REPOSITORY") == "" {
		return fmt.Errorf("GITHUB_REPOSITORY must be set in environment")
	}

	if os.Getenv("GITHUB_TOKEN") == "" {
		return fmt.Errorf("GITHUB_TOKEN must be set in environment")
	}

	if os.Getenv("COHERE_TOKEN") == "" {
		return fmt.Errorf("COHERE_TOKEN must be set in environment")
	}

	return nil
}

func RepositoryNWO() string {
	return os.Getenv("GITHUB_REPOSITORY")
}

func RepositoryOwner() string {
	nwo := os.Getenv("GITHUB_REPOSITORY")
	return strings.Split(nwo, "/")[0]
}

func RepositoryName() string {
	nwo := os.Getenv("GITHUB_REPOSITORY")
	return strings.Split(nwo, "/")[1]
}
