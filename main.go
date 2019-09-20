package main

import (
	"github.com/urfave/cli"
	"golang-example/cmd"

	"log"
	"os"

	_ "golang-example/lang-related/misc/sort"
	_ "golang-example/misc/distributedSystem/01-cluster-join"
	_ "golang-example/misc/distributedSystem/02-channel-event"
	_ "golang-example/misc/gops"
	_ "golang-example/misc/goreleaser"
	_ "golang-example/misc/gorm"
	_ "golang-example/misc/json"
	_ "golang-example/misc/xml"
)

func main() {
	app := cli.NewApp()
	app.Name = "golangExample"
	app.Usage = "golang 示例代码集"
	app.UsageText = `golangExample.exe [global options] command [command options] [arguments...]`
	app.Description = "提供了一些演示 golang 相关方面特性或库的使用方法"
	app.Version = "0.5.0"
	app.Author = "Ping"
	app.Email = "navono007@gmail.com"
	app.Commands = cmd.Cmds
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
