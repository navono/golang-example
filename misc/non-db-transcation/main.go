package non_db_transcation

import (
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "tx",
		Aliases: []string{"tx"},

		Usage:    "Demonstration of ono-db transaction",
		Action:   txAction,
		Category: "Misc",
	})
}

func txAction(c *cli.Context) error {
	// 生成一个内存队列
	queue := NewMemoryQueue(10)

	// 生成一个 worker
	name := "test"
	worker := NewWorker(name, func(data map[string]string) {
		for key, value := range data {
			println("worker get task key: " + key + " value: " + value)
		}
	})

	queue.AddListener(worker)

	taskName := "test"
	// events 发送的任务事件
	configs := []map[string]string{
		{"task1": "SendEmail", "params1": "Hello world"},
		{"task2": "SendMQ", "params2": "Hello world"},
	}

	// 分发任务
	queue.Push(ClearTaskPrefix, taskName, nil)
	for _, conf := range configs {
		queue.Push(TaskPrefix, taskName, Task{Name: taskName, Group: taskName, Config: conf})
	}
	// queue.Push(CommitTaskPrefix, taskName, nil)
	queue.Push(ClearTaskPrefix, taskName, Task{Name: "task2", Group: taskName})

	queue.Start()
	// 停止队列
	queue.Stop()

	return nil
}
