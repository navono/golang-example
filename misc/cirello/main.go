package cirello

import (
	"github.com/urfave/cli"

	"golang-example/cmd"
	"golang-example/misc/cirello/process"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "cirello",
		Aliases: []string{"cirello"},

		Usage:    "Demonstration of cirello supervisor",
		Action:   cirelloAction,
		Category: "Process",
	})
}

func cirelloAction(c *cli.Context) error {
	args := []string{"localhost", "-t"}
	proc1 := process.NewProcess("ping", args...)
	process.ProcessRegister(proc1)

	// args := []string{"-d", "-f"}
	// proc1 := process.NewProcess("dhcp", args...)
	// process.ProcessRegister(proc1)

	// supervisor := oversight.New(
	// 	// oversight.WithRestartStrategy(oversight.OneForOne()),
	// 	oversight.Processes(func(ctx context.Context) error {
	// 		select {
	// 		case <-ctx.Done():
	// 			return nil
	// 		case <-time.After(time.Second):
	// 			log.Println(1)
	// 		}
	// 		return nil
	// 	}),
	// )
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// if err := supervisor.Start(ctx); err != nil {
	// 	log.Fatal(err)
	// }
	// var supervisor supervisor.Supervisor
	//
	// svc := Simpleservice(1)
	// supervisor.Add(&svc)
	//
	// // Simply, if not special context is needed:
	// // supervisor.Serve()
	// // Or, using context.Context to propagate behavior:
	// s := make(chan os.Signal, 1)
	// signal.Notify(s, os.Interrupt)
	// ctx, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	<-s
	// 	fmt.Println("halting supervisor...")
	// 	cancel()
	// }()
	//
	// supervisor.Serve(ctx)
	return nil
}
