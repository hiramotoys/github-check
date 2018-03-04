package clt

import "github.com/google/go-github/github"

type clt struct {
	Github
}

var cltInstance *clt = newClt()

func newClt() *clt {
	c := clt{}
	c.Github = github.NewClient(nil)
	orgs, _, err := github_client.Organizations.List(context.Background(), "willnorris", nil)
	return &c
}

func GetClt() *clt {
	return cltInstance
}
