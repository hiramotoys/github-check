package clt

import (
	"github.com/google/go-github/github"
)

type clt struct {
	Github *github.Client
}

var cltInstance *clt = newClt()

func newClt() *clt {
	c := clt{}
	c.Github = github.NewClient(nil)
	return &c
}

func GetClt() *clt {
	return cltInstance
}
