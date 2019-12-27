package trace_pkg

import (
	"os"
	"runtime/trace"

	"github.com/urfave/cli"

	"golang-example/cmd"
)

// go run main.go trace 2> trace.out
// go tool trace trace.out

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "trace",
		Aliases: []string{"trace"},

		Usage:    "Demonstration of trace",
		Action:   traceAction,
		Category: "pprof",
	})
}

func traceAction(c *cli.Context) error {
	trace.Start(os.Stderr)
	defer trace.Stop()
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	<-ch

	return nil
}
