package job

import (
	"fmt"
	"sync"
)

type (
	Job struct {
		name          string
		cancel        chan interface{}
		cancelDone    chan interface{}
		canceled      bool
		tasks         []Task
		completedTask []Task
		wg            sync.WaitGroup
	}
)

func NewJob(name string) *Job {
	return &Job{
		name:       name,
		cancel:     make(chan interface{}),
		cancelDone: make(chan interface{}),
	}
}

func (j *Job) AddTask(t Task) {
	ctx := &taskContextVal{
		name:   j.name,
		index:  0,
		cancel: j.cancel,
	}

	t.setCtx(ctx)
	j.tasks = append(j.tasks, t)
}

func (j *Job) Run() *Result {
	fmt.Println("job running...")
	defer close(j.cancel)
	defer close(j.cancelDone)

	go j.handleCancel()

	for _, task := range j.tasks {
		go j.runTask(task)
	}

	j.wg.Wait()
	if len(j.completedTask) == len(j.tasks) {
		return &Result{
			Error: nil,
			Data:  "completed",
		}
	}

	return nil
}

func (j *Job) runTask(t Task) {
	if j.canceled {
		fmt.Println("job canceled")
		return
	}

	j.wg.Add(1)
	resultChan := make(chan *Result, 1)

	go func() {
		defer j.wg.Done()
		result := t.Execute()
		resultChan <- result
	}()

	select {
	case result := <-resultChan:
		if result.Error != nil {
			<-j.cancelDone

			// return &Result{
			// 	Error: result.Error,
			// 	Data:  nil,
			// }
		} else {
			j.completedTask = append(j.completedTask, t)
		}
	}
}

func (j *Job) handleCancel() {
	for {
		select {
		case <-j.cancel:
			fmt.Println("receive job cancel event")
			j.canceled = true
			go j.cancelTask(j.cancelDone)
		}
	}
}

func (j *Job) cancelTask(done chan<- interface{}) {
	fmt.Println("start job task cancel")
	for _, task := range j.completedTask {
		task.Clear()
	}

	done <- true
}
