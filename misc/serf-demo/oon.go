package serf_demo

import (
	"sync"
)

type oneAndOnlyNumber struct {
	num        int
	generation int
	numMutex   sync.RWMutex
}

const MembersToNotify = 2

func initTheNumber(val int) *oneAndOnlyNumber {
	return &oneAndOnlyNumber{
		num: val,
	}
}

func (oon *oneAndOnlyNumber) setValue(newVal int) {
	oon.numMutex.Lock()
	defer oon.numMutex.Unlock()

	oon.num = newVal
	oon.generation = oon.generation + 1
}

func (oon *oneAndOnlyNumber) getValue() (int, int) {
	oon.numMutex.RLock()
	defer oon.numMutex.RUnlock()

	return oon.num, oon.generation
}

func (oon *oneAndOnlyNumber) notifyValue(curVal int, curGeneration int) bool {
	if curGeneration > oon.generation {
		oon.numMutex.Lock()
		defer oon.numMutex.Unlock()

		oon.generation = curGeneration
		oon.num = curVal
		return true
	}
	return false
}
