package trace_handler

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/urfave/cli"

	"golang-example/cmd"
)

// run program

// 使用 wrk 进行压力测试（wrk 需要在 WSL 中使用）：
// wrk -c 100 -t 10 -d 60s http://localhost:8181/hello

// 收集数据：
// curl localhost:8181/debug/pprof/trace?seconds=5 > trace.out

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "trace-h",
		Aliases: []string{"trace-h"},

		Usage:    "Demonstration of trace with handler",
		Action:   traceAction,
		Category: "pprof",
	})
}

func traceAction(c *cli.Context) error {
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	return http.ListenAndServe("localhost:8181", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
