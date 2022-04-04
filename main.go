package main

import (
	"fmt"
	"hotel-management-system/app/database"
	"hotel-management-system/app/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.GetConnection()
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error when loading env file")
	}

	routes.Routes()

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	fmt.Printf("Server start on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic("Error listening")
	}
}
