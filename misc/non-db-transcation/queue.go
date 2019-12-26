package non_db_transcation

import (
	"sync"
)

type (
	MemoryQueue struct {
		done      chan struct{}
		queue     chan Event
		listeners []EventListener
		wg        sync.WaitGroup
	}
)

func NewMemoryQueue(capacity uint) *MemoryQueue {
	return &MemoryQueue{
		done:  make(chan struct{}),
		queue: make(chan Event),
		wg:    sync.WaitGroup{},
	}
}

func (mq *MemoryQueue) Push(eventType, name string, value interface{}) {
	mq.wg.Add(1)
	mq.queue <- Event{Key: eventType, Name: name, Value: value}
}

func (mq *MemoryQueue) AddListener(listener EventListener) bool {
	for _, eventListener := range mq.listeners {
		if eventListener == listener {
			return false
		}
	}

	mq.listeners = append(mq.listeners, listener)
	return true
}

func (mq *MemoryQueue) Notify(event *Event) {
	defer mq.wg.Done()

	for _, listener := range mq.listeners {
		listener.onEvent(event)
	}
}

func (mq *MemoryQueue) poll() {
	for {
		select {
		case <-mq.done:
			break
		case event := <-mq.queue:
			mq.Notify(&event)
		}
	}
}

func (mq *MemoryQueue) Start() {
	go mq.poll()
}

func (mq *MemoryQueue) Stop() {
	mq.wg.Wait()
	close(mq.done)
}
