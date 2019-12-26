package scipipe

import (
	. "github.com/scipipe/scipipe"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "scipipe",
		Aliases: []string{"scipipe"},

		Usage:    "Demonstration of scipipe",
		Action:   scipipeAction,
		Category: "Misc",
	})
}

func scipipeAction(c *cli.Context) error {
	wf := NewWorkflow("FuncHookWf", 4)

	foo := NewFooer(wf, "foo")
	f2b := NewFoo2Barer(wf, "f2b")

	foo.OutFoo().To(f2b.InFoo())

	wf.Run()

	return nil
}
