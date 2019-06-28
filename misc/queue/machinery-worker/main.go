package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func main() {
	var conf = config.Config{
		Broker:        "redis://127.0.0.1:6379",
		ResultBackend: "redis://127.0.0.1:6379",
	}

	server, err := machinery.NewServer(&conf)
	if err != nil {
		_ = fmt.Errorf("can not create server, %v", err.Error())
		return
	}

	err = server.RegisterTask("Say", Say)
	if err != nil {
		_ = fmt.Errorf("could not register task!, %v", err.Error())
		return
	}

	worker := server.NewWorker("worker-1", 10)
	err = worker.Launch()
	if err != nil {
		_ = fmt.Errorf("could not launch worker!, %v", err.Error())
		return
	}
}

// Say "Hello World"
func Say(name string) (string, error) {
	return "Hello " + name + "!", nil
}
