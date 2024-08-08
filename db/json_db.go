package db

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"sync"
	"time"
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

func generateId() string {
	sha256 := sha256.New()
	sha256.Write([]byte(time.Now().String()))

	return hex.EncodeToString(sha256.Sum(nil))
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

func (s *storeManager) set(store storeSchema) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	formattedJson, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_TRUNC, 0644)

	_, err = f.Write(formattedJson)
	if err != nil {
		return err
	}

	return nil
}

type SpendingsStoreJson struct {
	mu sync.RWMutex
}

func NewSpendingStoreJson() SpendingsStore {
	return &SpendingsStoreJson{}
}

func (s *SpendingsStoreJson) Insert(spending Spending) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	spending.Id = generateId()
	store.Spedings = append(store.Spedings, spending)
	store.Balance = spending.Price

	err = jsonMgr.set(store)
	if err != nil {
		return err
	}

	return nil
}

func (s *SpendingsStoreJson) GetAll() ([]Spending, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	store, err := jsonMgr.get()
	if err != nil {
		return nil, err
	}

	return store.Spedings, nil
}

func (s *SpendingsStoreJson) Update(id string, values Spending) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store, err := jsonMgr.get()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(store.Spedings, func(s Spending) bool {
		return s.Id == id
	})

	if idx == -1 {
		return fmt.Errorf("Item was not found")
	}

	store.Balance += store.Spedings[idx].Price
	store.Balance -= values.Price

	store.Spedings[idx] = values
	err = jsonMgr.set(store)
	if err != nil {
		return err
	}

	return nil
}

func (s *SpendingsStoreJson) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	panic("not implemented")
}

type BalanceStoreJson struct {
	mu sync.RWMutex
}

func NewBalanceStoreJson() BalanceStore {
	return &BalanceStoreJson{}
}

func (b *BalanceStoreJson) GetBalance() int64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	panic("not implemented")
}

func (b *BalanceStoreJson) SetBalance(_ int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	panic("not implemented")
}
