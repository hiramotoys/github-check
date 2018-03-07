package task

import (
	"context"
	"fmt"
	"testing"
)

func TestGithubClient(t *testing.T) {
	client := GetGithubClient()
	repo, _, _ := client.Github.Repositories.Get(context.Background(), "hiramotoys", "cookbook-sample")
	fmt.Println(repo.CreatedAt)
	if *repo.GitURL != "git://github.com/hiramotoys/cookbook-sample.git" {
		t.Error("Failed GitURL is not equql.")
		t.Log("GitURL: ", *repo.GitURL)
	}
}

func TestTaskLoad(t *testing.T) {
	tsk := Task{}
	result := tsk.Load()
	if result != 0 {
		t.Error("Return value is invalid.")
	}
}
