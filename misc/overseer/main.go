package overseer

import (
	"fmt"
	"os"
	"time"

	pm "github.com/ShinyTrinkets/overseer"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "process",
		Aliases: []string{"process"},

		Usage:    "Demonstration of process manager",
		Action:   overseerAction,
		Category: "Misc",
	})
}

func overseerAction(c *cli.Context) error {
	ovr := pm.NewOverseer()

	cmdOptions := pm.Options{
		Buffered:  false,
		Streaming: true,
	}

	id1 := "ping1"
	pingCmd := ovr.Add(id1, "ping", []string{"localhost"}, cmdOptions)

	statusFeed := make(chan *pm.ProcessJSON)
	ovr.Watch(statusFeed)

	// Capture status updates from the command
	go func() {
		for state := range statusFeed {
			fmt.Printf("STATE: %v\n", state)
		}
	}()

	// Capture STDOUT and STDERR lines streaming from Cmd
	// If you don't capture them, they will be written into
	// the overseer log to Info or Error.
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case line := <-pingCmd.Stdout:
				fmt.Println(line)
			case line := <-pingCmd.Stderr:
				fmt.Fprintln(os.Stderr, line)
			case <-ticker.C:
				if !ovr.IsRunning() {
					fmt.Println("Closing Stdout and Stderr loop")
				}
			}
		}
	}()

	// Run and wait for all commands to finish
	ovr.SuperviseAll()

	// Even after the command is finished, you can still access detailed info
	time.Sleep(100 * time.Millisecond)
	fmt.Println(ovr.Status(id1))
	return nil
}
