package deadlock

import (
	"golang-example/cmd"
	"net/http"
	"sync"
	"time"

	"github.com/sasha-s/go-deadlock"
	"github.com/urfave/cli"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "deadlock",
		Aliases: []string{"dl"},

		Usage:    "Detect dead lock with go-deadlock",
		Action:   deadlockAction,
		Category: "Misc",
	})
}

func deadlockAction(c *cli.Context) error {
	// 检测超过 100 ms 的锁等待
	deadlock.Opts.DeadlockTimeout = time.Millisecond * 100

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			do()
		}()
	}
	wg.Wait()
	return nil
}

var mu deadlock.Mutex
var url = "http://baidu.com:90"

func do() {
	mu.Lock()
	defer mu.Unlock()

	u := url
	http.Get(u) // 非预期的在持有锁期间做 IO 操作，导致锁等待时间变长
}
