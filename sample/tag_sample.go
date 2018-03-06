package main

import (
	"github.com/hiramotoys/github-check/repository"
)

func main() {
	b := new(repository.Branch)
	b.IsHeadTagged()
}
