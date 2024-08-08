package db

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)

const dbPath = "./db.json"

type storeManager struct {
	mu sync.RWMutex
}

type storeSchema struct {
	Spedings []Spending `json:"spending"`
	Balance  int64      `json:"balance"`
}

var jsonMgr = &storeManager{}

func init() {
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		_, err = os.Create(dbPath)
		if err != nil {
			panic(err)
		}
	}
}

func (s *storeManager) get() (storeSchema, error) {
	s.mu.RLock()

	defer s.mu.RUnlock()
	var store storeSchema

	f, err := os.Open(dbPath)
	if err != nil {
		return storeSchema{}, err
	}

	err = json.NewDecoder(f).Decode(&store)

	if errors.Is(err, io.EOF) {
		return storeSchema{}, nil
	}

	if err != nil {
		return storeSchema{}, err
	}
	return store, nil
}
