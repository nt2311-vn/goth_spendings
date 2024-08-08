package handlers

import (
	"fmt"
	"net/http"

	"github.com/nt2311-vn/goth_spendings/components"
	"github.com/nt2311-vn/goth_spendings/db"
	"github.com/nt2311-vn/goth_spendings/services"
)

func HomePage(w http.ResponseWriter, r *http.Request) error {
	balanceStore := db.NewBalanceStoreJson()
	spendingStore := db.NewSpendingStoreJson()
	balanceService := services.NewBalanceService(balanceStore)
	spendingService := services.NewSpendingService(spendingStore)

	spendings, err := spendingService.ListItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("Cannot list item on spendings")
	}

	w.Header().Set("Content-Type", "text/html")

	return Render(

		w,
		r,
		components.Index(balanceService.GetBalance(), spendings),
	)
}
