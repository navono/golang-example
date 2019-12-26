package work

import (
	"errors"
	"fmt"
	"time"

	"golang-example/misc/non-db-transcation2/job"
)

type (
	worker struct {
		job.TaskContext
		name string
		id   int
	}
)

func NewWorker(name string, id int) *worker {
	return &worker{
		id:   id,
		name: name,
	}
}

func (w *worker) Execute() *job.Result {
	fmt.Printf("start work: %d\n", w.id)

	// 模拟错误
	if w.name == "short" {
		fmt.Printf("cancel job by: %d\n", w.id)
		w.Cancel()
		return &job.Result{
			Error: errors.New("bad request"),
			Data:  nil,
		}
	}
	if w.name == "long" {
		fmt.Printf("do hard work %d\n", w.id)
		duration := time.Duration(5 * w.id)
		time.Sleep(time.Second * duration)
	} else {
		fmt.Printf("do hard work %d\n", w.id)
		duration := time.Duration(2 * w.id)
		time.Sleep(time.Second * duration)
	}

	fmt.Printf("%d work done\n", w.id)
	return &job.Result{
		Error: nil,
		Data:  "completed",
	}
}

func (w *worker) Clear() {
	fmt.Printf("clear work: %d\n", w.id)
}
