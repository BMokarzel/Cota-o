package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/BMokarzel/quote/database"
	_ "github.com/go-sql-driver/mysql"
)

func QuoteUSD() (database.Quote, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/last/USD", nil)
	if err != nil {
		return database.Quote{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return database.Quote{}, err
	}

	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return database.Quote{}, err
	}

	fmt.Println(string(r))

	var newQuote database.Quote

	err = json.Unmarshal(r, &newQuote)
	if err != nil {
		return database.Quote{}, err
	}

	file, err := os.Create("Cotacao.txt")
	if err != nil {
		return database.Quote{}, err
	}

	_, err = file.WriteString(fmt.Sprintf("Cotação USD: %s", newQuote.USDBRL.Bid))
	if err != nil {
		return database.Quote{}, err
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/quote")
	if err != nil {
		return database.Quote{}, err
	}

	defer db.Close()

	err = database.InsertQuote(db, newQuote)
	if err != nil {
		return database.Quote{}, err
	}

	return newQuote, nil

}

func QuoteEUR() (database.Quote, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/last/EUR", nil)
	if err != nil {
		return database.Quote{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return database.Quote{}, err
	}

	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return database.Quote{}, err
	}

	fmt.Println(string(r))

	var newQuote database.Quote

	err = json.Unmarshal(r, &newQuote)
	if err != nil {
		return database.Quote{}, err
	}

	file, err := os.Create("Cotacao.txt")
	if err != nil {
		return database.Quote{}, err
	}

	_, err = file.WriteString(fmt.Sprintf("Cotação EUR: %s", newQuote.EURBRL.Bid))
	if err != nil {
		return database.Quote{}, err
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/quote")
	if err != nil {
		return database.Quote{}, err
	}

	defer db.Close()

	err = database.InsertQuote(db, newQuote)
	if err != nil {
		return database.Quote{}, err
	}

	return newQuote, nil
}

func QuoteBTC() (database.Quote, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/last/BTC", nil)
	if err != nil {
		return database.Quote{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return database.Quote{}, err
	}

	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return database.Quote{}, err
	}

	fmt.Println(string(r))

	var newQuote database.Quote

	err = json.Unmarshal(r, &newQuote)
	if err != nil {
		return database.Quote{}, err
	}

	file, err := os.Create("Cotacao.txt")
	if err != nil {
		return database.Quote{}, err
	}

	_, err = file.WriteString(fmt.Sprintf("Cotação BTC: %s", newQuote.USDBRL.Bid))
	if err != nil {
		return database.Quote{}, err
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/quote")
	if err != nil {
		return database.Quote{}, err
	}

	defer db.Close()

	err = database.InsertQuote(db, newQuote)
	if err != nil {
		return database.Quote{}, err
	}

	return newQuote, nil

}
