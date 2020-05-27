package server

import (
	"fmt"
	"github.com/MSHR-Dec/pantogram-badger/store"
	"net/http"
)

type Server struct {}

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	delayReader := store.NewDelayReader()

	pref := delayReader.Get(key)
	fmt.Fprintf(w,"%s", pref)
}
