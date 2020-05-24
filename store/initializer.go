package store

import (
	"encoding/json"
	"github.com/MSHR-Dec/pantogram-badger/driver"
	"io/ioutil"
	"log"
	"os"
)

type BadgerInit struct {
	store *driver.Store
}

func NewBadgerInit() *BadgerInit {
	return &BadgerInit{
		store: driver.NewStore(),
	}
}

type prefecture struct {
	Com  string `json:"com"`
	Pref string `json:"pref"`
}

func (bi *BadgerInit) Execute() {
	pwd, _ := os.Getwd()
	bytes, err := ioutil.ReadFile(pwd+"/store/prefecture.json")
	if err != nil {
		log.Println(err)
	}

	var prefectures []prefecture
	if err := json.Unmarshal(bytes, &prefectures); err != nil {
		log.Println(err)
	}

	txn := bi.store.Conn.NewTransaction(true)
	for _, v := range prefectures {
		txn.Set([]byte(v.Com), []byte(v.Pref))
	}

	txn.Commit()
	bi.store.Conn.Close()
}
