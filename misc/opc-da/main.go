package opc_da

import (
	"fmt"

	"github.com/konimarti/opc"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "opc-da",
		Aliases: []string{"opc-da"},

		Usage:    "Demonstration of OPC DA",
		Action:   opcAction,
		Category: "OPC",
	})
}

func opcAction(c *cli.Context) error {
	progid := "Graybox.Simulator"
	nodes := []string{"localhost"}

	// create browser and collect all tags on OPC server
	browser, err := opc.CreateBrowser(progid, nodes)
	if err != nil {
		panic(err)
	}

	// extract subtree
	subtree := opc.ExtractBranchByName(browser, "textual")

	// print out all the information
	opc.PrettyPrint(subtree)

	// create opc connection with all tags from subtree
	conn, _ := opc.NewConnection(
		progid,
		nodes,
		opc.CollectTags(subtree),
	)
	defer conn.Close()

	fmt.Println(conn.Read())

	return nil
}
