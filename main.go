package main

import (
	"database/sql"
	"log"

	"github.com/Erikadarisman/go-rest-api/driver"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	godotenv.Load()
}

func main() {
	db = driver.ConnectDB()

	log.Println("Server started successfully on port 8000.!!!")
}
