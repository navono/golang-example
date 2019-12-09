package dynamic_struct

import (
	"fmt"

	"github.com/ompluscator/dynamic-struct"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "dynamic-struct",
		Aliases: []string{"ds"},

		Usage:    "Demonstration of dynamic struct",
		Action:   dsAction,
		Category: "Misc",
	})
}

func dsAction(c *cli.Context) error {
	pv := dynamicstruct.NewStruct().
		AddField("ID", "", `json:"id"`).
		AddField("Description", "", `json:"desc"`).
		Build().
		New()

	fmt.Println(pv)

	return nil
}
