package gops

import (
	"github.com/google/gops/agent"
	"github.com/urfave/cli"
	"golang-example/cmd"

	"fmt"
	"log"
	"time"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gops",
		Aliases: []string{"ps"},

		Usage:    "Start exe with gops",
		Action:   goAgent,
		Category: "perf",
	})
}

func goAgent(c *cli.Context) error {
	fmt.Println("Running app with gops agent.")
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Hour)
	return nil
}
