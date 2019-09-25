package bolt_db

import (
	bh "github.com/timshannon/bolthold"

	"reflect"
	"strings"
)

type anonStorer struct {
	rType   reflect.Type
	indexes map[string]bh.Index
}

// Type returns the name of the type as determined from the reflect package
func (t *anonStorer) Type() string {
	return t.rType.Name()
}

// Indexes returns the Indexes determined by the reflect package on this type
func (t *anonStorer) Indexes() map[string]bh.Index {
	return t.indexes
}

func newStorer(dataType interface{}, encode bh.EncodeFunc) bh.Storer {
	str, ok := dataType.(bh.Storer)

	if ok {
		return str
	}

	tp := reflect.TypeOf(dataType)

	for tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	storer := &anonStorer{
		rType:   tp,
		indexes: make(map[string]bh.Index),
	}

	if storer.rType.Name() == "" {
		panic("Invalid Type for Storer.  Type is unnamed")
	}

	if storer.rType.Kind() != reflect.Struct {
		panic("Invalid Type for Storer.  BoltHold only works with structs")
	}

	for i := 0; i < storer.rType.NumField(); i++ {
		if strings.Contains(string(storer.rType.Field(i).Tag), bh.BoltholdIndexTag) {
			indexName := storer.rType.Field(i).Tag.Get(bh.BoltholdIndexTag)

			if indexName != "" {
				indexName = storer.rType.Field(i).Name
			}

			storer.indexes[indexName] = func(name string, value interface{}) ([]byte, error) {
				tp := reflect.ValueOf(value)
				for tp.Kind() == reflect.Ptr {
					tp = tp.Elem()
				}

				return encode(tp.FieldByName(name).Interface())
			}
		}
	}

	return storer
}
