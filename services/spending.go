package services

import (
	"time"

	"github.com/nt2311-vn/goth_spendings/db"
)

type SpendingService struct {
	store db.SpendingsStore
}

func NewSpendingService(store db.SpendingsStore) *SpendingService {
	return &SpendingService{store}
}

func (s *SpendingService) AddItem(spending db.Spending) error {
	spending.SpentAt = time.Now()

	return s.store.Insert(spending)
}

func (s *SpendingService) ListItems() ([]db.Spending, error) {
	return s.store.GetAll()
}

func (s *SpendingService) UpdateItem(id string, newVal db.Spending) error {
	return s.store.Update(id, newVal)
}

func (s *SpendingService) DeleteItem(id string) error {
	return s.store.Delete(id)
}

type BalanceService struct {
	store db.BalanceStore
}

func NewBalanceService(store db.BalanceStore) *BalanceService {
	return &BalanceService{store: store}
}

func (b *BalanceService) GetBalance() int64 {
	return b.store.GetBalance()
}
