package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nt2311-vn/goth_spendings/components"
)

func HandleGreet(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	ageStr := r.URL.Query().Get("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return fmt.Errorf("Invalid age request: %s", ageStr)
	}
	return Render(w, r, components.Greet(name, age))
}
