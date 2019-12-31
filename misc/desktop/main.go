package desktop

import (
	"log"
	"net/url"

	"github.com/urfave/cli"
	"github.com/zserge/lorca"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:     "desktop",
		Aliases:  []string{"desktop"},
		Usage:    "Demonstration of desktop",
		Action:   lorcaAction,
		Category: "Desktop",
	})
}

func lorcaAction(c *cli.Context) error {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()

	return nil
}
