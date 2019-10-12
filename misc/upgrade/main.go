package upgrade

import (
	"fmt"
	"github.com/urfave/cli"
	"golang-example/cmd"
	"net/http"
	"time"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "upgrade",
		Aliases: []string{"upg"},

		Usage:    "Start exe with upgrade",
		Action:   upgAgent,
		Category: "misc",
	})
}

// BuildID is compile-time variable
var BuildID = "0"

func upgAgent(c *cli.Context) error {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":5001",
		Fetcher: &fetcher.File{Path: "./my_app_next"},
		//Fetcher: &fetcher.HTTP{
		//	URL:      "http://localhost:4000/myapp.exe",
		//	Interval: 1 * time.Second,
		//},
		Debug: true, //display log of overseer actions
	})
	return nil
}

//convert your 'main()' into a 'prog(state)'
//'prog()' is run in a child process
func prog(state overseer.State) {
	fmt.Printf("app#%s (%s) listening...\n", BuildID, state.ID)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, _ := time.ParseDuration(r.URL.Query().Get("d"))
		time.Sleep(d)
		_, _ = fmt.Fprintf(w, "app#%s (%s) says hello\n", BuildID, state.ID)
	}))
	_ = http.Serve(state.Listener, nil)
	fmt.Printf("app#%s (%s) exiting...\n", BuildID, state.ID)
}
