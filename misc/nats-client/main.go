package nats_client

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

// go run main.go nats -s "nats://10.30.26.99:4222" aa bb
func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "nats",
		Aliases: []string{"nats"},

		Usage:    "Demonstration of nats client",
		Action:   natsClientAction,
		Category: "MQ",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "s",
				Usage:    "The nats server URLs (separated by comma)",
				Required: true,
			},
			cli.StringFlag{
				Name:  "creds",
				Usage: "User Credentials File",
			},
			// cli.BoolFlag{
			// 	Name:        "h",
			// 	Usage:       "Show help message",
			// 	Destination: nil,
			// },
		},
	})
}

func natsClientAction(c *cli.Context) error {
	// var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	// var userCreds = flag.String("creds", "", "User Credentials File")
	// var showHelp = flag.Bool("h", false, "Show help message")

	urls := c.String("s")
	userCreds := c.String("creds")
	showHelp := c.Bool("h")

	// log.SetFlags(0)
	// flag.Usage = usage
	// flag.Parse()

	if showHelp {
		showUsageAndExit(0)
	}

	if c.NArg() != 2 {
		showUsageAndExit(1)
	}

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	// Use UserCredentials
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}

	// Connect to NATS
	nc, err := nats.Connect(urls, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	args := c.Args()
	subj, msg := args[0], []byte(args[1])

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}

	return nil
}

func usage() {
	log.Printf("Usage: nats-pub [-s server] [-creds file] <subject> <msg>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}
