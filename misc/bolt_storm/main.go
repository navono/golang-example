package bolt_storm

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"golang-example/cmd"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/msgpack"
	"github.com/urfave/cli"
	"go.etcd.io/bbolt"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "storm",
		Aliases: []string{"sm"},

		Usage:    "Start exe with bolt db by storm",
		Action:   smAction,
		Category: "DB",
	})
}

type Nested struct {
	ID   int `storm:"id,increment"`
	Name string
	Embed
	L1 Nest
	L2 Level2
	//Pointer *Nest
	Pointer interface{}
}

type Embed struct {
	Color string
}

type Nest struct {
	Name string
}

type Level2 struct {
	Name string
	L3   Nest
}

func smAction(c *cli.Context) error {
	//fmt.Println(GetGID())

	options := []func(*storm.Options) error{
		storm.Codec(msgpack.Codec),
		storm.BoltOptions(0600, &bbolt.Options{Timeout: 1 * time.Second}),
	}

	db, err := storm.Open("./misc/bolt_storm/db", options...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 填充
	fillData(db, "bucket1", "Name1")
	fillData(db, "bucket2", "Name2")

	// 更新
	for i := 1; i < 3; i++ {
		go updateData(db, i)
	}

	// 查询
	queryData(db)

	// 删除
	deleteData(db)

	var nl []Nested
	_ = db.From("bucket1").All(&nl)

	return nil
}

func fillData(db *storm.DB, bucketName, name string) {
	node, err := db.Begin(true)
	if err != nil {
		fmt.Println(err)
	}
	defer node.Rollback()

	err = node.From(bucketName).Save(&Nested{
		Name: name,
		Embed: Embed{
			Color: "red",
		},
		L1: Nest{
			Name: "Xiv",
		},
		L2: Level2{
			Name: "Xiv2",
			L3:   Nest{Name: "Xiv3"},
		},
		Pointer: &Nest{Name: "Xiv"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = node.Commit()
}

func queryData(db *storm.DB) {
	var n []Nested
	err := db.From("bucket1").Find("Color", "red", &n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func updateData(db *storm.DB, id int) {
	tx, err := db.Begin(true)
	if err != nil {
		return
	}

	//defer tx.Rollback()
	defer func() {
		if err := tx.Rollback(); err != nil {
			return
		}
	}()

	err = tx.From("bucket1").Update(&Nested{
		ID:   id,
		Name: fmt.Sprintf("update %d name", id),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = tx.Commit()
	return
}

func deleteData(db *storm.DB) {
	////Read tx
	//readTx, err := db.Bolt.Begin(false)
	//if err != nil {
	//	return
	//}
	//defer readTx.Rollback()
	//
	//bs := readTx.Bucket([]byte("bucket1"))
	//fmt.Println(bs)
	//_ = readTx.Commit()

	//err := db.Bolt.View(func(tx *bbolt.Tx) error {
	//	projectBkt := tx.Bucket([]byte("bucket1"))
	//	if projectBkt == nil {
	//		// 无数据
	//		return errors.New("not found")
	//	}
	//
	//	return nil
	//})
	//fmt.Println(err)

	if err := db.From("bucket1").DeleteStruct(&Nested{
		ID: 1,
	}); err != nil {
		fmt.Println(err)
	}

	queryData(db)
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
