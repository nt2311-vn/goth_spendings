package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
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

	h := &handlers.Handlers{
		SpendingService: spendingService,
		BalanceService:  balanceService,
	}

	spendingHandler := handlers.NewSpendingHandler(*spendingService)

	r := chi.NewRouter()
	r.Get("/", handlers.Make(h.HomePage))
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/greet", handlers.Make(handlers.HandleGreet))

	r.Post("/spending/add", spendingHandler.HandleAddSpendingItem)

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)

	http.ListenAndServe(listenAddr, r)
}
