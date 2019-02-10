package main

import (
	"gorilla/handlers"
	"net/http"

	"github.com/emmanuelbenson/gpi-validate-v2/routes"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := routes.InitRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// log.Fatal(http.ListenAndServe(":9005", handlers.CORS(allowedOrigins, allowedMethods)(router)))

	http.ListenAndServe(":3003", handlers.CORS(allowedOrigins, allowedMethods)(r))
}
