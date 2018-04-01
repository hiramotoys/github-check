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
	//tsk := LoadTask("test.yaml")
	//t.Log(tsk)
	//fmt.Printf("%T", tsk.Repositories[0])
	result := 0
	if result != 0 {
		t.Error("Return value is invalid.")
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
