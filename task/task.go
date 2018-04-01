package task

import (
	"context"
	"errors"
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

func BranchHeadIsTagged(owner string, repoName string, branchName string) (bool, error) {
	client := GetGithubClient()
	branch, _, err := client.Github.Repositories.GetBranch(context.Background(), owner, repoName, branchName)
	if err != nil {
		return false, errors.New("Get branch error.")
	}
	tags, _, err := client.Github.Repositories.ListTags(context.Background(), owner, repoName, nil)
	if err != nil {
		return false, errors.New("Get tags error.")
	}
	for _, tag := range tags {
		fmt.Println(*branch.Commit.SHA)
		fmt.Println(*tag.Commit.SHA)
		if *branch.Commit.SHA == *tag.Commit.SHA {
			return true, nil
		}
	}
	return false, nil
}

type Repository struct {
	Owner         string `yaml: "owner"`
	Name          string `yaml: "name"`
	Branch        string `yaml: "branch"`
	Check         string `yaml: "check"`
	ResultStatus  int
	ResultMessage string
}

const RepositoryCheckModeTag = "tag"

const ResultStatusCodeOk = 0
const ResultStatusCodeNg = 1
const ResultStatusCodeErr = 2

//Task definition
type Task struct {
	Repositories []Repository `yaml: "repository"`
}

func LoadTask(filename string) (*Task, error) {
	fmt.Println(filename)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fmt.Printf("buf: %+v\n", string(buf))
	var t Task
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *Task) Run() int {
	return 0
}

func (t *Task) runRepositoryChecker() {
	for _, repository := range t.Repositories {
		if repository.Check == RepositoryCheckModeTag {
			repository.tagCheck()
		}
	}
}

func (r *Repository) tagCheck() {
	ok, err := BranchHeadIsTagged(r.Owner, r.Name, r.Branch)
	if err != nil {
		r.ResultStatus = ResultStatusCodeErr
		r.ResultMessage = "Repository error."
		return
	}
	if ok {
		r.ResultStatus = ResultStatusCodeOk
		r.ResultMessage = "Branch head is tagged."
	} else {
		r.ResultStatus = ResultStatusCodeNg
		r.ResultMessage = "Branch head is not tagged."
	}
}

func (t *Task) PushResult() int {
	return 0
}

//TaskRunner definition
type TaskRunner struct {
}

func (taskRunner *TaskRunner) execute(filename string) int {
	tsk, err := LoadTask(filename)
	if err != nil {
		panic(err)
	}
	tsk.runRepositoryChecker()
	return 0

}
