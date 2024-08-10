package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/nt2311-vn/goth_spendings/components"
	"github.com/nt2311-vn/goth_spendings/db"
	"github.com/nt2311-vn/goth_spendings/handlers"
	"github.com/nt2311-vn/goth_spendings/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	balanceStore := db.NewBalanceStoreJson()
	spendingStore := db.NewSpendingStoreJson()
	balanceService := services.NewBalanceService(balanceStore)
	spendingService := services.NewSpendingService(spendingStore)

	spendingHandler := handlers.NewSpendingHandler(*spendingService)

	r := chi.NewRouter()

	// Define the main page route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		spendings, err := spendingService.ListItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		components.Index(balanceService.GetBalance(), spendings).Render(r.Context(), w)
	})

	// Serve static files
	fmt.Println("Static files path: ", filepath.Join("static"))
	staticDir := filepath.Join("static")
	r.Handle("/static/**", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	r.Post("/api/spending", spendingHandler.HandleAddSpendingItem)
	r.Put("/api/spending", spendingHandler.HandleUpdateSpendingItem)
	r.Delete("/api/spending", spendingHandler.HandleRemoveSpendingItem)

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)

	http.ListenAndServe(listenAddr, r)
}
