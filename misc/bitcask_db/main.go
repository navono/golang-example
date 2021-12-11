package bitcask_db

import (
	"fmt"
	"github.com/urfave/cli"
	"golang-example/cmd"

	"git.mills.io/prologic/bitcask"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "bitcask",
		Aliases: []string{"bc"},

		Usage:    "Demonstration of bitcask db",
		Action:   bcAction,
		Category: "DB",
	})
}

func bcAction(c *cli.Context) error {
	db, _ := bitcask.Open("./misc/bitcask_db/db")
	defer db.Close()

	_ = db.Put([]byte("Hello"), []byte("World"))
	val, _ := db.Get([]byte("Hello"))
	fmt.Println(string(val))

	return nil
}
