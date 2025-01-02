package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Debsnil24/Go_Postgre.git/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type response struct{
	ID int `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Postgres")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to Decode the Request Body. %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID: insertID,
		Message: "Stock Created Successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get Stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatalf("Unable to get stocks %v", err)
	}
	
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	var stock models.Stock

	err = json.NewEncoder(w).Encode(&stock)
	if err != nil {
		log.Fatalf("Unable to Decode request body. %v", err)
	}
	updateRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stocks updated successfully.\n %d rows/record affected", updateRows)
	res := response {
		ID: int(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stocks Deleted successfully.\n %d rows/records affected", deletedRows)
	res := response {
		ID: int(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}