package task

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

type Repository struct {
	Owner  string `yaml: "owner"`
	Name   string `yaml: "name"`
	Branch string `yaml: "branch"`
}

//Task definition
type Task struct {
	Repositories []Repository `yaml: "repository"`
}

func LoadTask(filename string) *Task {
	fmt.Println(filename)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("buf: %+v\n", string(buf))
	var t Task
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("t: %+v", t)
	fmt.Println("return")
	return &t
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
