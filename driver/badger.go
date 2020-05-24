package driver

import (
	"github.com/dgraph-io/badger/v2"
	"log"
)

type Store struct {
	Conn *badger.DB
}

func NewStore() *Store {
	opts := badger.DefaultOptions("/tmp/badger")
	opts.Logger = nil
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
