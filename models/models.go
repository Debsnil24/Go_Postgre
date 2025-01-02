package models

type Stock struct{
	StockID int `json:"stockid"`
	Name string `json:"name"`
	Price int `json:"price"`
	Company string `json:"company"`
}