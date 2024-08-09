package handlers

import (
	"fmt"
	"net/http"

	"github.com/nt2311-vn/goth_spendings/components"
)

func (h *Handlers) HomePage(w http.ResponseWriter, r *http.Request) error {
	spendings, err := h.SpendingService.ListItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("Cannot list item on spendings")
	}

	w.Header().Set("Content-Type", "text/html")

	return Render(
		w,
		r,
		components.Index(h.BalanceService.GetBalance(), spendings),
	)
}
