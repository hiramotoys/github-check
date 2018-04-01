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
	tsk, err := LoadTask("test.yaml")
	if err != nil {
		t.Error("Failed can't load a yaml file.")
	}
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
	repo2 := tsk.Repositories[1]
	if repo2.Name != "cookbook-sample" {
		t.Error("Failed can't load repository name correctly.")
	}
	if repo2.Branch != "develop" {
		t.Error("Failed can't load repository branch correctly.")
	}
	if repo2.Owner != "hiramotoys" {
		t.Error("Failed can't load repository owner correctly.")
	}
}

func TestLoadTaskNoExistsYaml(t *testing.T) {
	_, err := LoadTask("test_no_exists.yaml")
	if err == nil {
		t.Error("Failed error is nil")
	}
}

func TestTaskRun(t *testing.T) {
	tsk := Task{}
	result := tsk.Run()
	if result != 0 {
		t.Error("Return value is invalid.")
	}
}

func TestBranchHeadIsTaggedOk(t *testing.T) {
	b, e := BranchHeadIsTagged("hiramotoys", "cookbook-sample", "master")
	if e != nil {
		t.Errorf("%s\n", e)
	}
	if !b {
		t.Error("Failed: the branch is tagged.")
	}
}

func TestBranchHeadIsTaggedNg(t *testing.T) {
	b, e := BranchHeadIsTagged("hiramotoys", "cookbook-sample", "develop")
	if e != nil {
		t.Errorf("%s\n", e)
	}
	if b {
		t.Error("Failed: the branch is not tagged.")
	}
}

func TestRepositoryTagCheckOk(t *testing.T) {
	r := Repository{}
	r.Owner = "hiramotoys"
	r.Name = "cookbook-sample"
	r.Branch = "master"
	r.Check = "tag"
	r.tagCheck()
	if r.ResultStatus != ResultStatusCodeOk {
		t.Error("Failed result status is not correct.")
	}
	r.Branch = "develop"
	r.tagCheck()
	if r.ResultStatus != ResultStatusCodeNg {
		t.Error("Failed result status is not correct.")
	}
}
