package main

import (
	"github.com/google/gops/agent"
	"github.com/urfave/cli"
	"golang-example/cmd"

	"fmt"
	"log"
	"os"

	_ "golang-example/app/go-kit"
	_ "golang-example/lang-related/misc/sort"
	_ "golang-example/misc/bitcask_db"
	_ "golang-example/misc/bolt_db"
	_ "golang-example/misc/bolt_storm"
	_ "golang-example/misc/bolt_tx_deadlock"
	_ "golang-example/misc/deadlock"
	_ "golang-example/misc/distributedSystem/01-cluster-join"
	_ "golang-example/misc/distributedSystem/02-channel-event"
	_ "golang-example/misc/gabs"
	_ "golang-example/misc/gops"
	_ "golang-example/misc/goreleaser"
	_ "golang-example/misc/gorm"
	_ "golang-example/misc/upgrade"
	_ "golang-example/misc/watermill"
	_ "golang-example/misc/xml"
	_ "golang-example/pattern"
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

	fmt.Println("Running app with gops agent.")
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
