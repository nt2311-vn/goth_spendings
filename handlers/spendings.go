package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nt2311-vn/goth_spendings/db"
	"github.com/nt2311-vn/goth_spendings/services"
)

type SpendingsHandler struct {
	serivce services.SpendingService
}

func NewSpendingHandler(service services.SpendingService) *SpendingsHandler {
	return &SpendingsHandler{service}
}

func (s *SpendingsHandler) HandleAddSpendingItem(w http.ResponseWriter, r *http.Request) {
	var spending db.Spending

	err := json.NewDecoder(r.Body).Decode(&spending)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.serivce.AddItem(spending)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *SpendingsHandler) HandleRemoveSpendingItem(w http.ResponseWriter, r *http.Request) {
	id, exist := r.URL.Query()["id"]

	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.serivce.DeleteItem(id[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *SpendingsHandler) HandleUpdateSpendingItem(w http.ResponseWriter, r *http.Request) {
	id, exist := r.URL.Query()["id"]

	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newSpending db.Spending

	err := json.NewDecoder(r.Body).Decode(&newSpending)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.serivce.UpdateItem(id[0], newSpending)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
