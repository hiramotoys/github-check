package tag

import (
	"context"
	"fmt"
	"github.com/hiramotoys/github-check/clt"
)

func IsBranchHeadTagged(branch_name string) bool {
	client := clt.GetClt()
	orgs, _, err := client.Github.Organizations.List(context.Background(), "willnorris", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false
	}
	for i, organization := range orgs {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
	return true
}
