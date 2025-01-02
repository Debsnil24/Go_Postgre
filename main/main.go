package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Debsnil24/Go_Postgre.git/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting Server on Port 8080.....")
	log.Fatal(http.ListenAndServe(":8080", r))
}