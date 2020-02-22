package main

import (
	"net/http"
	"os"

	"github.com/Erikadarisman/go-rest-api/controllers"
	"github.com/jackc/pgx/pgxpool"

	"github.com/Erikadarisman/go-rest-api/driver"
	"github.com/joho/godotenv"
	log "gopkg.in/inconshreveable/log15.v2"
)

func init() {
	godotenv.Load()
}

var db *pgxpool.Pool

func main() {
	db := driver.ConnectDB()
	controller := controllers.Controller{}

	http.HandleFunc("/listcategories", controller.GetAllCategory(db))

	log.Info("Starting URL shortener on localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Crit("Unable to start web server", "error", err)
		os.Exit(1)
	}

}
