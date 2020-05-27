package server

import (
	"context"
	"github.com/MSHR-Dec/pantogram-badger/badgerpb"
	"github.com/MSHR-Dec/pantogram-badger/store"
)

type BadgerServer struct {}

func (s *BadgerServer) GetPrefecture(ctx context.Context, msg *badgerpb.PrefRequest) (*badgerpb.PrefResponse, error) {
	delayReader := store.NewDelayReader()

	pref := delayReader.Get(msg.Company)

	return &badgerpb.PrefResponse{
		Prefecture: pref,
	}, nil
}
