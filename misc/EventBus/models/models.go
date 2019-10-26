package models

import (
	"sync"
)

// Order struct for sample event
type Order struct {
	Name   string
	Amount int
	// Ack       chan interface{}
	SyncGroup sync.WaitGroup
}
