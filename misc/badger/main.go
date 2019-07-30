// https://colobu.com/2017/10/11/badger-a-performant-k-v-store/
package badger

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"log"
)

func init() {
	opts := badger.DefaultOptions("E:\\data\\badger")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// set
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	if err != nil {
		panic(err)
	}

	// get
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			fmt.Printf("The answer is: %s\n", val)
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	// iterate
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(val []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, val)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Prefix scans
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		prefix := []byte("ans")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(val []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, val)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	// iterate keys
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			fmt.Printf("key=%s\n", k)
		}
		return nil
	})

	// delete
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte("answer"))
	})
	if err != nil {
		panic(err)
	}
}
