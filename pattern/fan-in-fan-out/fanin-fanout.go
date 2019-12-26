package fan_in_fan_out

import (
	"fmt"
	"sync"
	"time"

	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "fan",
		Aliases: []string{"fan"},

		Usage:    "Demonstration of fan-in and fan-out",
		Action:   fanAction,
		Category: "Pattern",
	})
}

func fanAction(c *cli.Context) error {
	done := make(chan bool)
	defer close(done)

	start := time.Now()
	items := prepareItems(done)
	workers := make([]<-chan int, 4)
	for i := 0; i < 4; i++ {
		workers[i] = packItems(done, items, i)
	}

	numPackages := 0
	for range merge(done, workers...) {
		numPackages++
	}
	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)

	return nil
}

type (
	item struct {
		ID            int
		Name          string
		PackingEffort time.Duration
	}
)

func prepareItems(done <-chan bool) <-chan item {
	items := make(chan item)
	itemsToShip := []item{
		{0, "Shirt", 1 * time.Second},
		{1, "Legos", 1 * time.Second},
		{2, "TV", 5 * time.Second},
		{3, "Bananas", 2 * time.Second},
		{4, "Hat", 1 * time.Second},
		{5, "Phone", 2 * time.Second},
		{6, "Plates", 3 * time.Second},
		{7, "Computer", 5 * time.Second},
		{8, "Pint Glass", 3 * time.Second},
		{9, "Watch", 2 * time.Second},
	}

	go func() {
		for _, item := range itemsToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()

	return items
}

func packItems(done <-chan bool, items <-chan item, workerID int) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
				fmt.Printf("Worker #%d: Shipping package no. %d, took %ds to pack\n", workerID, item.ID, item.PackingEffort/time.Second)
			}
		}

		close(packages)
	}()
	return packages
}

func merge(done <-chan bool, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingPackages := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case outgoingPackages <- i:
			}
		}
	}

	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(outgoingPackages)
	}()

	return outgoingPackages
}
