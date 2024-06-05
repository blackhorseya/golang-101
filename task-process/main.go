package main

import (
	"fmt"
)

func main() {
	task := NewTask("task")
	fmt.Println("current step:", task.Step.String())

	for task.Execute() == nil {
		fmt.Println("current step:", task.Step.String())
	}
}
