package main

import (
	"fmt"
	"github.com/hiramotoys/github-check/clt"
	"github.com/hiramotoys/github-check/tag"
)

func main() {
	fmt.Println("aaaaaa")
	b := tag.IsBranchHeadTagged("master")
	fmt.Println(b)
}
