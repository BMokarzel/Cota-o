package database

import (
	"database/sql"
	"fmt"
)

type Quote struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
	EURBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
	BTCBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
}

func CreateTables(db *sql.DB) error {
	createTable := `CREATE TABLE IF NOT EXISTS quote(
		coin VARCHAR(3) NOT NULL,
		bid VARCHAR(10) NOT NULL
	)`

	_, err := db.Exec(createTable)
	if err != nil {
		return err
	}

	return nil
}

func InsertQuote(db *sql.DB, q Quote) error {

	stmt, err := db.Prepare("INSERT INTO quote(coin, bid) VALUES (?, ?)")
	//	stmt, err := db.Prepare("insert into quotes(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, creatDate) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		fmt.Printf("Erro 11: %v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(q.USDBRL.Name, q.USDBRL.Bid)

	//	_, err = stmt.Exec(q.USDBRL.Code, q.USDBRL.Codein, q.USDBRL.Name, q.USDBRL.High, q.USDBRL.Low, q.USDBRL.VarBid, q.USDBRL.PctChange, q.USDBRL.Bid, q.USDBRL.Ask, q.USDBRL.Timestamp, q.USDBRL.CreateDate)
	if err != nil {
		fmt.Println("erro 12")
	}

	return nil
}
