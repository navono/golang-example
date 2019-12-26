package non_db_transcation2

import (
	"github.com/urfave/cli"

	"golang-example/cmd"
	"golang-example/misc/non-db-transcation2/job"
	"golang-example/misc/non-db-transcation2/work"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "ndt",
		Aliases: []string{"ndt"},

		Usage:    "Demonstration of pipeline",
		Action:   ndtAction,
		Category: "Misc",
	})
}

func ndtAction(c *cli.Context) error {
	j := job.NewJob("test")

	t2 := work.NewWorker("short", 2)
	t1 := work.NewWorker("long", 1)
	t3 := work.NewWorker("a", 3)

	j.AddTask(t1)
	j.AddTask(t2)
	j.AddTask(t3)
	j.Run()

	return nil
}
