package bolt_tx_deadlock

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/msgpack"
	"github.com/urfave/cli"
	"go.etcd.io/bbolt"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "stormTxDead",
		Aliases: []string{"smtd"},

		Usage:    "bbolt db deadlock with tx",
		Action:   deadlockAction,
		Category: "DB",
	})
}

type person struct {
	ID   int `storm:"id,increment"`
	Name string
	Age  uint16
}

func deadlockAction(c *cli.Context) error {
	options := []func(*storm.Options) error{
		storm.Codec(msgpack.Codec),
		storm.BoltOptions(0600, &bbolt.Options{Timeout: 1 * time.Second}),
	}

	db, err := storm.Open("./misc/bolt_tx_deadlock/db", options...)
	if err != nil {
		fmt.Println(err)
	}

	fillData(db, "students", "john", 10)
	fillData(db, "students", "lily", 12)

	for i := 0; i < 2; i++ {
		go updateData(db, i)
	}

	queryData(db)
	return nil
}

func fillData(db *storm.DB, bucketName, name string, age uint16) {
	node, err := db.Begin(true)
	if err != nil {
		fmt.Println(err)
	}
	defer node.Rollback()

	err = node.From(bucketName).Save(&person{
		Name: name,
		Age:  age,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = node.Commit()
}

func updateData(db *storm.DB, id int) {
	tx, err := db.Begin(true)
	if err != nil {
		return
	}
	defer tx.Rollback()

	if err := tx.From("students").Update(&person{
		ID:   id,
		Name: fmt.Sprintf("update %d name", id),
	}); err != nil {
		fmt.Println(err)
		return
	}

	_ = tx.Commit()
}

func queryData(db *storm.DB) {
	var p []person
	err := db.From("students").All(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
