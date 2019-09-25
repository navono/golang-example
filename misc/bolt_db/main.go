package bolt_db

import (
	"fmt"
	bh "github.com/timshannon/bolthold"
	"github.com/urfave/cli"
	"go.etcd.io/bbolt"
	"golang-example/cmd"
	"strings"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "bolthold",
		Aliases: []string{"bh"},

		Usage:    "Start exe with goAgent",
		Action:   bhAction,
		Category: "DB",
	})
}

var testBucket = []byte("test-bucket")

type person struct {
	Name string
	Age  uint32
}

func bhAction(c *cli.Context) error {
	store, err := bh.Open("./misc/bolt_db/db", 0666, nil)
	if err != nil {
		panic(err)
	}

	defer store.Close()

	item := person{
		Name: "ping",
		Age:  90,
	}

	err = store.Upsert("key", &item)
	err = store.Upsert("key2", &person{
		Name: "super",
		Age:  100,
	})
	if err != nil {
		panic(err)
	}

	tx, err := store.Bolt().Begin(true)
	if err != nil {
		panic(err)
	}

	err = store.TxUpsert(tx, "key3", &person{
		Name: "xxxx",
		Age:  20,
	})

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	var p person
	err = store.Get("key3", &p)
	if err != nil {
		panic(err)
	}

	err = dump(store.Bolt())
	if err != nil {
		panic(err)
	}
	//err = store.Bolt().View(func(tx *bbolt.Tx) error {
	//	var k person
	//	bktKey := newStorer(k, bh.DefaultEncode)
	//	bkt := tx.Bucket([]byte(bktKey.Type()))
	//
	//	c := bkt.Cursor()
	//	for k, v := c.First(); k != nil; k, v = c.Next() {
	//		var pv person
	//		err = bh.DefaultDecode(v, &pv)
	//		fmt.Println(pv)
	//	}
	//
	//	return nil
	//})

	var result []person
	err = store.Find(&result, bh.Where("Age").Eq(uint32(20)))
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func dump(db *bbolt.DB) error {
	return db.View(func(tx *bbolt.Tx) error {
		var pk person
		bktKey := newStorer(pk, bh.DefaultEncode)
		// 枚举 person 类型的 bucket 下的所有 kv
		c := tx.Bucket([]byte(bktKey.Type())).Cursor()

		// 枚举所有 bucket 下的所有 kv
		//c := tx.Cursor()
		dumpCursor(tx, c, 1)
		return nil
	})
}

func dumpCursor(tx *bbolt.Tx, c *bbolt.Cursor, indent int) {
	for k, v := c.First(); k != nil; k, v = c.Next() {
		if v == nil {
			fmt.Printf(strings.Repeat("  ", indent-1)+"[%s]\n", k)
			newBucket := c.Bucket().Bucket(k)
			if newBucket == nil {
				// from the top-level cursor and not a cursor from a bucket
				newBucket = tx.Bucket(k)
			}
			newCursor := newBucket.Cursor()
			dumpCursor(tx, newCursor, indent+1)
		} else {
			fmt.Printf(strings.Repeat("  ", indent-1)+"%s\n", k)

			var pv person
			err := bh.DefaultDecode(v, &pv)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf(strings.Repeat("  ", indent-1)+"  %v\n", pv)
			}
		}
	}
}
