package nats_client

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

// go run main.go nats publish -s "nats://10.30.26.99:4222" aa bb
func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "nats",
		Aliases: []string{"nats"},

		Usage:    "Demonstration of nats client",
		Category: "MQ",
		Subcommands: []cli.Command{
			{
				Name:   "publish",
				Usage:  "basic nats publish",
				Action: publishAction,
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
				},
			},
			{
				Name:   "reply",
				Usage:  "basic nats reply",
				Action: replyAction,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "s",
						Usage:    "The nats server URLs (separated by comma)",
						Required: true,
					},
					cli.StringFlag{
						Name:  "q",
						Value: "NATS-RPLY-22",
						Usage: "Message queue",
					},
					cli.StringFlag{
						Name:  "creds",
						Usage: "User Credentials File",
					},
				},
			},
			{
				Name:   "request",
				Usage:  "basic nats request",
				Action: requestAction,
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
				},
			},
		},
	})
}

func publishAction(c *cli.Context) error {
	urls := c.String("s")
	userCreds := c.String("creds")

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

func replyAction(c *cli.Context) error {
	urls := c.String("s")
	queueName := c.String("q")
	userCreds := c.String("creds")

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
	subj, reply, i := args[0], []byte(args[1]), 0

	nc.QueueSubscribe(subj, queueName, func(msg *nats.Msg) {
		i++
		printMsg(msg, i)
		msg.Respond([]byte(reply))
	})
	nc.Flush()

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatalf("Exiting")
	return nil
}

func requestAction(c *cli.Context) error {
	urls := c.String("s")
	userCreds := c.String("creds")

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
	subj, payload := args[0], []byte(args[1])

	msg, err := nc.Request(subj, payload, 2*time.Second)
	if err != nil {
		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		log.Fatalf("%v for request", err)
	}

	log.Printf("Published [%s] : '%s'", subj, payload)
	log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))

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

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}
