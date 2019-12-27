package pointer_copy

import (
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "pointer",
		Aliases: []string{"pointer"},

		Usage:    "Demonstration of pointer and copy",
		Action:   pointerAction,
		Category: "Misc",
	})
}

func pointerAction(c *cli.Context) error {
	return nil
}

type (
	S struct {
		a, b, c int64
		d, e, f string
		g, h, i float64
	}
)

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}
