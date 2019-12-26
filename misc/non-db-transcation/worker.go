package non_db_transcation

import (
	"fmt"
	"strings"
)

type (
	Worker struct {
		name                string
		deferredTaskUpdates map[string][]Task
		onCommit            ConfigUpdateCallback
	}
)

type ConfigUpdateCallback func(data map[string]string)

func NewWorker(name string, fn ConfigUpdateCallback) *Worker {
	return &Worker{
		name:                name,
		deferredTaskUpdates: make(map[string][]Task),
		onCommit:            fn,
	}
}

func (w *Worker) onEvent(event *Event) {
	switch {
	case strings.Contains(event.Key, TaskPrefix):
		w.onTaskEvent(event)
	case strings.Contains(event.Key, ClearTaskPrefix):
		w.onTaskClear(event)
	case strings.Contains(event.Key, CommitTaskPrefix):
		w.onTaskCommit(event)
	}
}

func (w *Worker) onTaskClear(event *Event) {
	task, err := event.Value.(Task)
	if !err {
		// log
		return
	}
	_, found := w.deferredTaskUpdates[task.Group]
	if !found {
		return
	}
	delete(w.deferredTaskUpdates, task.Group)

	// 还可以继续停止本地已经启动的任务
	fmt.Printf("clear %v\n", task)
}

// onTaskCommit 接收任务提交, 从延迟队列中取出数据然后进行业务逻辑处理
func (w *Worker) onTaskCommit(event *Event) {
	// 获取之前本地接收的所有任务
	tasks, found := w.deferredTaskUpdates[event.Name]
	if !found {
		return
	}

	// 获取配置
	config := w.getTasksConfig(tasks)
	if w.onCommit != nil {
		w.onCommit(config)
	}
	delete(w.deferredTaskUpdates, event.Name)
}

// onTaskEvent 接收任务数据，此时需要丢到本地暂存不能进行应用
func (w *Worker) onTaskEvent(event *Event) {
	task, err := event.Value.(Task)
	if !err {
		// log
		return
	}

	// 保存任务到延迟更新map
	configs, found := w.deferredTaskUpdates[task.Group]
	if !found {
		configs = make([]Task, 0)
	}
	configs = append(configs, task)
	w.deferredTaskUpdates[task.Group] = configs
}

// getTasksConfig 获取task任务列表
func (w *Worker) getTasksConfig(tasks []Task) map[string]string {
	config := make(map[string]string)
	for _, t := range tasks {
		config = t.UpdateConfig(config)
	}
	return config
}
