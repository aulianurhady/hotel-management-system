package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func GetConnection() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_DBNAME")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Cannot connect to DB, cause %s", err)
		return err
	}

	err = result.Ping()
	if err != nil {
		log.Fatalf("Cannot ping db, cause %s", err)
		return err
	}

	DB = result

	return nil
}
