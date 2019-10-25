package watermill

import (
	"context"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/urfave/cli"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message/subscriber"
)

func httpAction(c *cli.Context) error {
	pub, sub := createPubSub()
	defer func() {
		_ = pub.Close()
		_ = sub.Close()
	}()

	msg, err := sub.Subscribe(context.Background(), "/test")
	if err != nil {
		return err
	}

	go sub.StartHTTPServer()

	receivedMessages := make(chan message.Messages)
	go func() {
		received, _ := subscriber.BulkRead(msg, 100, time.Second*1)
		receivedMessages <- received
	}()

	// go func() {
	// }()
	for {
		select {
		case v, _ := <-msg:
			// 	received, _ := subscriber.BulkRead(msg, 100, time.Second*1)

			fmt.Println(v)
			v.Ack()

		case msg := <-receivedMessages:
			fmt.Println(msg)
		}
	}

	return nil
}

func createPubSub() (*http.Publisher, *http.Subscriber) {
	logger := watermill.NewStdLogger(true, true)

	sub, err := http.NewSubscriber(":10888", http.SubscriberConfig{}, logger)
	if err != nil {
		return nil, nil
	}

	publisherConf := http.PublisherConfig{
		MarshalMessageFunc: http.DefaultMarshalMessageFunc,
	}

	pub, err := http.NewPublisher(publisherConf, logger)

	return pub, sub
}
