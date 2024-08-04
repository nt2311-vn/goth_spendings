package handlers

import (
	"net/http"

	"github.com/nt2311-vn/goth_spendings/components"
)

func HomePage(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, components.Index())
}
