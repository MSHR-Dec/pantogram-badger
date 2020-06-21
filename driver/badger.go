package driver

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"log"
)

type Store struct {
	Conn *badger.DB
}

func NewStore() *Store {
	opts := badger.DefaultOptions("/tmp/badger")
	opts.Logger = nil
	opts.ValueLogLoadingMode = options.FileIO
	opts.TableLoadingMode = options.FileIO
	opts.NumMemtables = 2
	opts.NumLevelZeroTables = 2
	opts.NumLevelZeroTablesStall = 4
	opts.MaxTableSize = 64 << 5
	opts.ValueLogFileSize = 1<<10 - 1
	db, err := badger.Open(opts)
	if err != nil {
		log.Println(err)
	}

	return &Store{
		Conn: db,
	}
}

func (s *Store) Close()  {
	defer s.Conn.Close()
}
