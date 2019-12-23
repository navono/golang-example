package pipeline

import (
	"fmt"
	"time"

	"github.com/myntra/pipeline"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "pipe",
		Aliases: []string{"process"},

		Usage:    "Demonstration of pipeline",
		Action:   pipeAction,
		Category: "Misc",
	})
}

func pipeAction(c *cli.Context) error {

	// create a new pipeline
	workpipe := pipeline.NewProgress("myProgressworkpipe", 1000, time.Second*3)
	// func NewStage(name string, concurrent bool, disableStrictMode bool) *Stage
	// To execute steps concurrently, set concurrent=true.
	stage := pipeline.NewStage("mypworkstage", false, false)

	// a unit of work
	step1 := &work{id: 1}
	// another unit of work
	step2 := &work{id: 2}

	// add the steps to the stage. Since concurrent is set false above. The steps will be
	// executed one after the other.
	stage.AddStep(step1)
	stage.AddStep(step2)

	// add the stage to the pipe.
	workpipe.AddStage(stage)

	stage2 := pipeline.NewStage("testStage", true, false)

	sstep1 := &work{id: 3}
	sstep2 := &work{id: 4}

	stage2.AddStep(sstep1)
	stage2.AddStep(sstep2)

	workpipe.AddStage(stage2)

	go readPipeline(workpipe)
	sstep2.Cancel()

	result := workpipe.Run()
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println("timeTaken:", workpipe.GetDuration())
	return nil
}
