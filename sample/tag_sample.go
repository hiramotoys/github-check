package main

import (
	"fmt"
	"github.com/hiramotoys/github-check/tag"
)

func main() {
	b := tag.IsBranchHeadTagged("master")
	fmt.Println(b)
}
