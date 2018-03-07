package task

import (
	"context"
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

//Github utility.
func GetRepository(owner string, repoName string) *github.Repository {
	client := GetGithubClient()
	repo, _, err := client.Github.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return nil
	}
	return repo
}

//Task definition
type Task struct {
	isDryRun *bool
}

func (t *Task) Load() int {
	return 0
}

func (t *Task) Run() int {
	return 0
}

func (t *Task) runTagChecker(owner string, repoName string) {
	GetRepository(owner, repoName)
}

func (t *Task) PushResult() int {
	return 0
}

//TaskRunner definition
type TaskRunner struct {
}

func (taskRunner *TaskRunner) execute() int {
	return 0
}
