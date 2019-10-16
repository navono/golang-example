package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "gjson",
		Aliases: []string{"gj"},

		Usage:    "Demonstration of gjson operations",
		Action:   gjsonAction,
		Category: "data",
	})
}

const data = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func gjsonAction(c *cli.Context) error {
	r := gjson.Get(data, "name.list")

	fmt.Print(r.Exists())
	fmt.Print(r)

	return nil
}
