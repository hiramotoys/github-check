package main

import (
	"fmt"
	"github.com/hiramotoys/github-check/task"
)

func main() {
	fmt.Println("hello")
	tsk := task.LoadTask("../task/test.yaml")
	fmt.Println(tsk.Repositories[0].Name)
	b, e := task.BranchHeadIsTagged("hiramotoys", "cookbook-sample", "master")
	if e == nil {
		fmt.Errorf("%s\n", e)
	}
	fmt.Println(b)
	//tsk.Repositories[0]
}
