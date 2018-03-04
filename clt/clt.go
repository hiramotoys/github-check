package clt

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
)

type clt struct {
	Github *github.Client
}

var cltInstance *clt = newClt()

func newClt() *clt {
	c := clt{}
	c.Github = github.NewClient(nil)
	orgs, _, err := c.Github.Organizations.List(context.Background(), "willnorris", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}
	for i, organization := range orgs {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
	return &c
}

func GetClt() *clt {
	return cltInstance
}
