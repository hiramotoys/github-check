package main

import (
	"fmt"
	"github.com/hiramotoys/github-check/task"
)

func main() {
	fmt.Println("hello")
	tsk := task.LoadTask("../task/test.yaml")
	fmt.Println(tsk.Repositories[0].Name)
	//tsk.Repositories[0]
}
