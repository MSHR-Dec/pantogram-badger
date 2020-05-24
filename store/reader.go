package store

import (
	"github.com/MSHR-Dec/pantogram-badger/driver"
	"github.com/dgraph-io/badger/v2"
)

type DelayReader struct {
	store *driver.Store
}

func NewDelayReader() *DelayReader {
	return &DelayReader{
		store: driver.NewStore(),
	}
}

func (dr *DelayReader) Get(key string) string {
	value := ""

	dr.store.Conn.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				if key == string(k) {
					value = string(v)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		dr.store.Conn.Close()
		return nil
	})

	return value
}
