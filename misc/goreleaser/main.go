package goreleaser

import (
	"fmt"
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "releaser",
		Aliases: []string{"rl"},

		Usage:    "Start exe with goAgent",
		Action:   releaseAction,
		Category: "misc",
	})
}

func releaseAction(c *cli.Context) error {
	fmt.Println("Go releaser example")
	return nil
}
