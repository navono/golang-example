package watermill

import (
	stdSQL "database/sql"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/urfave/cli"
)

var (
	logger   = watermill.NewStdLogger(false, false)
	pubTopic = "events"
	sqlTable = "events"
)

func txAction(c *cli.Context) error {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

	pub := createPublisher()
	defer func() {
		if err := pub.Close(); err != nil {
			panic(err)
		}
	}()

	// sub := createSubscriber(db)
	//
	// router.AddHandler()

	return nil
}

func createPublisher() message.Publisher {
	return gochannel.NewGoChannel(
		gochannel.Config{},
		logger,
	)
}

func createSubscriber(db *stdSQL.DB) message.Subscriber {
	sub, err := sql.NewSubscriber(
		db,
		sql.SubscriberConfig{
			ConsumerGroup:    "",
			SchemaAdapter:    sql.DefaultMySQLSchema{},
			OffsetsAdapter:   sql.DefaultPostgreSQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	return sub
}

func createDB() *stdSQL.DB {

	// stdSQL.Open()
	return nil
}
