package repository

import (
	"context"
	"fmt"
	"github.com/hiramotoys/github-check/clt"
)

type Repository struct {
	Name     *string
	UserName *string
	Branches []Branch
}

type Branch struct {
	Name *string
}

func (branch *Branch) IsHeadTagged() {
	client := clt.GetClt()
	orgs, _, err := client.Github.Organizations.List(context.Background(), "willnorris", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for i, organization := range orgs {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
	repo, _, err := client.Github.Repositories.Get(context.Background(), "hiramotoys", "nao-imageprocessing-module")
	fmt.Println(repo.CreatedAt)
}
