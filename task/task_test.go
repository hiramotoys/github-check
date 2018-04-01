package task

import (
	"context"
	"testing"
)

func TestGithubClient(t *testing.T) {
	client := GetGithubClient()
	repo, _, _ := client.Github.Repositories.Get(context.Background(), "hiramotoys", "cookbook-sample")
	if *repo.GitURL != "git://github.com/hiramotoys/cookbook-sample.git" {
		t.Error("Failed GitURL is not equal.")
		t.Log("GitURL: ", *repo.GitURL)
	}
}

func TestGetRepository(t *testing.T) {
	repo := GetRepository("hiramotoys", "cookbook-sample")
	if *repo.GitURL != "git://github.com/hiramotoys/cookbook-sample.git" {
		t.Error("Failed GitURL is not equal.")
		t.Log("GitURL: ", *repo.GitURL)
	}
}

func TestLoadTask(t *testing.T) {
	tsk := LoadTask("test.yaml")
	repo1 := tsk.Repositories[0]
	if repo1.Name != "cookbook-sample" {
		t.Error("Failed can't load repository name correctly.")
	}
	if repo1.Branch != "master" {
		t.Error("Failed can't load repository branch correctly.")
	}
	if repo1.Owner != "hiramotoys" {
		t.Error("Failed can't load repository owner correctly.")
	}
}

func TestTaskRun(t *testing.T) {
	tsk := Task{}
	result := tsk.Run()
	if result != 0 {
		t.Error("Return value is invalid.")
	}
}

func TestBranchHeadIsTagged(t *testing.T) {
	b, e := BranchHeadIsTagged("hiramotoys", "cookbook-sample", "master")
	if e != nil {
		t.Errorf("%s\n", e)
	}
	if !b {
		t.Error("Failed: the branch is tagged.")
	}
}
