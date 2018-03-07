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

//Task definition
type Task struct {
}

func (t Task) LoadTasks() int {
	return 0
}

func (t Task) RunTasks() int {
	return 0
}

func (t Task) PushResult() int {
	return 0
}

//TaskRunner definition
type TaskRunner struct {
}

func (taskRunner TaskRunner) execute() int {
	return 0
}
