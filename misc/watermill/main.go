package watermill

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/urfave/cli"
	"golang-example/cmd"

	"context"
	"log"
	"time"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "watermill",
		Aliases: []string{"wm"},

		Usage:    "Demonstration of watermill",
		Action:   wwAction,
		Category: "event driven",
	})
}

func wwAction(c *cli.Context) error {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	defer func() {
		if err := pubSub.Close(); err != nil {
			panic(err)
		}
	}()

	messages, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)
	publishMessages(pubSub)
	return nil
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
