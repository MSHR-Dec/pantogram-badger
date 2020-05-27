// protoc -I pb/ pb/pantogram.proto --go_out=plugins=grpc:.

package main

import (
	"github.com/MSHR-Dec/pantogram-badger/server"
	"github.com/MSHR-Dec/pantogram-badger/store"
	"log"
	"net/http"
)

func main() {
	badgerInit := store.NewBadgerInit()
	badgerInit.Execute()

	s := new(server.Server)

	http.HandleFunc("/", s.Search)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
