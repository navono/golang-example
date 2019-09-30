package bolt_db

import (
	"fmt"
	"strings"

	"golang-example/cmd"

	"github.com/satori/go.uuid"
	bh "github.com/timshannon/bolthold"
	"github.com/urfave/cli"
	"github.com/vmihailenco/msgpack/v4"
	"go.etcd.io/bbolt"
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

type User struct {
	ID           string `boltholdKey:"ID"`
	Name         string
	Age          uint16
	LanguageID   []string
	CreditCardID []string
	//Nested       []*test
	Nested *test
}

type test struct {
	Id, Name string
}

type Language struct {
	ID     string `boltholdKey:"ID"`
	Name   string
	UserID []string
}

type CreditCard struct {
	ID     string `boltholdKey:"ID"`
	Number string
	UserID string
}

type Nested struct {
	Key int
	Embed
	L1      Nest
	L2      Level2
	Pointer *Nest
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

func msgPackEncoder(value interface{}) ([]byte, error) {
	return msgpack.Marshal(&value)
}

func msgPackDecoder(data []byte, value interface{}) error {
	return msgpack.Unmarshal(data, &value)
}

func bhAction(c *cli.Context) error {
	options := bh.Options{
		Encoder: msgPackEncoder,
		Decoder: msgPackDecoder,
	}
	store, err := bh.Open("./misc/bolt_db/db", 0666, &options)
	if err != nil {
		panic(err)
	}

	defer store.Close()

	// fill data
	fillDB(store)
	fillNested(store)

	// Transaction
	transactUpdate(store)

	// View
	err = dump(store.Bolt())
	if err != nil {
		panic(err)
	}

	// Delete
	deleteItem(store, "key3")

	// Find
	findItem(store)

	return nil
}

func deleteItem(store *bh.Store, key string) {
	//var pk User
	//bktKey := newStorer(pk, msgPackEncoder)

	if err := store.Delete(key, &User{}); err != nil {
		fmt.Println(err)
	}
}

func fillDB(store *bh.Store) {
	item := User{
		ID:   uuid.NewV4().String(),
		Name: "ping",
		Age:  18,
		//Nested: []*test{t},
		Nested: &test{
			Id:   "111",
			Name: "test name",
		},
	}

	err := store.Upsert("key", &item)
	err = store.Upsert("key2", &User{
		ID:   uuid.NewV4().String(),
		Name: "super",
		Age:  30,
	})
	err = store.Upsert("key3", &User{
		ID:   uuid.NewV4().String(),
		Name: "ping",
		Age:  30,
	})
	if err != nil {
		panic(err)
	}
}

func fillNested(store *bh.Store) {
	var nestedData = []Nested{
		Nested{
			Key: 0,
			Embed: Embed{
				Color: "red",
			},
			L1: Nest{
				Name: "Joe",
			},
			L2: Level2{
				Name: "Joe",
				L3: Nest{
					Name: "Joe",
				},
			},
			Pointer: &Nest{
				Name: "Joe",
			},
		},
		Nested{
			Key: 1,
			Embed: Embed{
				Color: "red",
			},
			L1: Nest{
				Name: "Jill",
			},
			L2: Level2{
				Name: "Jill",
				L3: Nest{
					Name: "Jill",
				},
			},
			Pointer: &Nest{
				Name: "Jill",
			},
		},
	}

	for i := range nestedData {
		err := store.Upsert(nestedData[i].Key, nestedData[i])
		if err != nil {
			fmt.Printf("Error inserting nested test data for nested find test: %s\n", err)
		}
	}

	var result []*Nested
	err := store.Find(&result, bh.Where("L1.Name").Eq("Joe"))
	if err != nil {
		fmt.Printf("Error finding data from bolthold: %s\n", err)
	}
}

func transactUpdate(store *bh.Store) {
	tx, err := store.Bolt().Begin(true)
	if err != nil {
		panic(err)
	}

	l1 := Language{
		ID:     uuid.NewV4().String(),
		Name:   "en-us",
		UserID: nil,
	}

	var userIDList []string
	err = store.TxUpdateMatching(tx, &User{}, bh.Where("Name").Eq("ping"), func(record interface{}) error {
		// record will always be a pointer
		user, ok := record.(*User)
		if !ok {
			return fmt.Errorf("Record isn't the correct type!  Wanted Person, got %T", record)
		}

		user.LanguageID = append(user.LanguageID, l1.ID)
		userIDList = append(userIDList, user.ID)
		return nil
	})
	err = store.TxUpdateMatching(tx, &User{}, bh.Where("Name").Eq("super"), func(record interface{}) error {
		// record will always be a pointer
		user, ok := record.(*User)
		if !ok {
			return fmt.Errorf("Record isn't the correct type!  Wanted Person, got %T", record)
		}

		user.LanguageID = append(user.LanguageID, l1.ID)
		userIDList = append(userIDList, user.ID)
		return nil
	})

	if err != nil {
		err = tx.Rollback()
		panic(err)
	}

	l1.UserID = userIDList
	err = store.TxUpsert(tx, "lang", &l1)
	if err != nil {
		err = tx.Rollback()
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

func findItem(store *bh.Store) {
	var userList []User
	err := store.Find(&userList, bh.Where("Name").Eq("ping"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userList)

	var langList []Language
	err = store.Find(&langList, bh.Where("Name").Eq("en-us"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(langList)

	query := bh.Where("Name").MatchFunc(func(ra *bh.RecordAccess) (bool, error) {
		u := ra.Record().(*User)
		return strings.Compare(u.Name, "ping") == 0, nil
	})
	result, err := store.FindAggregate(&User{}, query, "Name")
	for i := range result {
		var name string
		u := &User{}

		result[i].Group(&name)
		result[i].Min("Age", u)
		fmt.Println(u)
	}
}

func dump(db *bbolt.DB) error {
	return db.View(func(tx *bbolt.Tx) error {
		var pk User
		//bktKey := newStorer(pk, bh.DefaultEncode)
		bktKey := newStorer(pk, msgPackEncoder)
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

			var pv User
			//err := bh.DefaultDecode(v, &pv)
			err := msgPackDecoder(v, &pv)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf(strings.Repeat("  ", indent-1)+"  %v\n", pv)
			}
		}
	}
}
