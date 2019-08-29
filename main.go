package main

import (
	"github.com/urfave/cli"
	"golang-example/cmd"

	"fmt"
	"log"
	"os"

	_ "golang-example/misc/distributedSystem/01-cluster-join"
	_ "golang-example/misc/distributedSystem/02-channel-event"
)

func main() {
	app := cli.NewApp()
	app.Name = "golang example"
	app.Description = "this is a set of demo of golang"
	app.Version = "0.5.0"
	app.Author = "Ping"
	app.Email = "navono007@gmail.com"
	app.Commands = cmd.Cmds
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	fmt.Println("main exit")
}
