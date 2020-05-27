// protoc -I badgerpb/ badgerpb/badger.proto --go_out=plugins=grpc:.

package main

import (
	"github.com/MSHR-Dec/pantogram-badger/badgerpb"
	"github.com/MSHR-Dec/pantogram-badger/server"
	"github.com/MSHR-Dec/pantogram-badger/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	badgerInit := store.NewBadgerInit()
	badgerInit.Execute()

	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	badgerpb.RegisterBadgerServer(s, &server.BadgerServer{})
	s.Serve(listenPort)
}
