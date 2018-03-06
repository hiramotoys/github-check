package task

import (
	"github.com/google/go-github/github"
)

//Github Client
type githubClient struct {
	Github *github.Client
}

var githubClientInstance *githubClient = newGithubClient()

func newGithubClient() *githubClient {
	client := githubClient{}
	client.Github = github.NewClient(nil)
	return &client
}

func GetGithubClient() *githubClient {
	return githubClientInstance
}
