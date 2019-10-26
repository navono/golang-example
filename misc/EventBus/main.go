package EventBus

import (
	"fmt"
	"math/rand"

	"golang-example/cmd"
	"golang-example/misc/EventBus/calculator"
	"golang-example/misc/EventBus/counter"
	"golang-example/misc/EventBus/models"
	"golang-example/misc/EventBus/printer"

	"github.com/mustafaturan/bus"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
	"github.com/urfave/cli"
)

func init() {
	initBus()

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "bus",
		Aliases: []string{"bus"},

		Usage:    "demonstration of bus event",
		Action:   busAgent,
		Category: "event driven",
	})
}

func initBus() {
	// configure id generator (it doesn't have to be monoton)
	node := uint(1)
	initialTime := uint(0)
	monoton.Configure(sequencer.NewMillisecond(), node, initialTime)

	// configure bus package
	if err := bus.Configure(bus.Config{Next: monoton.Next}); err != nil {
		panic("Whoops, couldn't configure the bus package!")
	}

	// register topics
	bus.RegisterTopics("order.created", "order.canceled")

	// load printer package
	printer.Load()

	// no need to load counter and calculator packages since we are running the
	// FetchEventCount, and TotalAmount function from the counter package, it
	// will auto execute the init function on load
}

func busAgent(c *cli.Context) error {
	txID := monoton.Next()
	for i := 0; i < 3; i++ {
		bus.Emit(
			"order.created",
			&models.Order{Name: fmt.Sprintf("Product #%d", i), Amount: randomAmount()},
			txID,
		)
	}

	evt, err := bus.Emit(
		"order.canceled", // topic
		&models.Order{Name: "Product #N", Amount: randomAmount()}, // data
		"", // when blank bus package auto assigns an ID using the provided gen
	)
	if err != nil {
		return err
	}

	// 等待被处理
	o := evt.Data.(*models.Order)
	o.SyncGroup.Wait()

	// printer consumer processed all events at that moment since it is synchronous
	fmt.Println("You should see 4 events printed above!^^^")

	// give some time to process events for async consumers
	// time.Sleep(time.Millisecond * 25)

	printEventCounts()
	printOrderTotalAmount()

	return nil
}

func printEventCounts() {
	// Let's print event counts for each topic
	for _, topic := range bus.ListTopics() {
		fmt.Printf(
			"Total event count for %s: %d\n",
			topic.Name,
			counter.FetchEventCount(topic.Name),
		)
	}
}

func printOrderTotalAmount() {
	fmt.Printf("Order total amount %d\n", calculator.TotalAmount())
}

func randomAmount() int {
	max := 100
	min := 10
	return rand.Intn(max-min) + min
}
