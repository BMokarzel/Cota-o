package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	//"time"

	service "github.com/BMokarzel/Cota-o/Service"
	"github.com/BMokarzel/Cota-o/database"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/quote")
	if err != nil {
		log.Panic("error to open db")
	}

	defer db.Close()

	database.CreateTables(db)

	s := service.NewService(db)

	r := chi.NewRouter()

	r.Get("/USD", s.QuoteUSDHandler)

	r.Get("/EUR", s.QuoteEURHandler)

	r.Get("/BTC", s.QuoteBTCHandler)

	http.ListenAndServe(":8080", r)

}
