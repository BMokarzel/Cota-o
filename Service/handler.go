package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Service) QuoteUSDHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	req, err := QuoteUSD()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(req.USDBRL.Bid)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) QuoteEURHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	req, err := QuoteEUR()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(req.EURBRL.Bid)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) QuoteBTCHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	req, err := QuoteBTC()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(req.BTCBRL.Bid)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
