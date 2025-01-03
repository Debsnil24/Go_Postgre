package middleware

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Debsnil24/Go_Postgre.git/models"
)

func insertStock(s models.Stock) int {
	db := CreateConnection()
	defer db.Close()
	sqlSt := `INSERT INTO stocks(name,price,company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int

	err := db.QueryRow(sqlSt, s.Name, s.Price, s.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the Query. %v\n", err)
	}

	fmt.Printf("Inserted a Single Record %v\n", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	var stock models.Stock
	sqlSt := `SELECT * FROM stocks WHERE stockid = $1`

	row := db.QueryRow(sqlSt, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock , nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row %v\n", err)
	}
	return stock, err
}

func getAllStocks() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stocks []models.Stock

	sqlSt := `SELECT * from stocks`

	rows, err := db.Query(sqlSt)
	if err != nil {
		log.Fatalf("Unable to execute the query %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row %v\n", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}

func updateStock(id int64, s models.Stock) int {
	db := CreateConnection()
	defer db.Close()

	sqlSt := `UPDATE stocks SET name = $2, price = $3, company = $4 WHERE stockid = $1`

	res, err := db.Exec(sqlSt, id, s.Name, s.Price, s.Company)
	if err != nil {
		log.Fatalf("Unable to execute the query %v\n", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking affected row %v\n", err)
	}
	fmt.Printf("%v rows/records affected \n", rowsAffected)
	return int(rowsAffected)
}

func deleteStock(id int64) int {
	db := CreateConnection()
	defer db.Close()

	sqlSt := `DELETE FROM stocks WHERE stockid = $1`
	
	res, err := db.Exec(sqlSt, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking affected row %v", err)
	}
	fmt.Printf("%v rows/records affected\n", rowsAffected)
	return int(rowsAffected)
}
