package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/nt2311-vn/goth_spendings/services"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

type Handlers struct {
	SpendingService *services.SpendingService
	BalanceService  *services.BalanceService
}

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handle error", "error", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
