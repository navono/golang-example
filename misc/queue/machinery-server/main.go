package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

func main() {

	var cnf = config.Config{
		Broker:        "redis://127.0.0.1:6379",
		ResultBackend: "redis://127.0.0.1:6379",
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		_ = fmt.Errorf("can not create server, %v", err.Error())
		return
	}

	sayTask := tasks.Signature{
		Name: "Say",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "god",
			},
		},
	}

	asyncResult, err := server.SendTask(&sayTask)
	if err != nil {
		fmt.Println("send task failed")
		return
	}

	taskState := asyncResult.GetState()
	fmt.Printf("Current state of %v task is:\n", taskState.TaskUUID)
	fmt.Println(taskState.State)

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	for _, result := range results {
		fmt.Println(result.Interface())
	}
}
