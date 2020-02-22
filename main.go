package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

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

	router := mux.NewRouter()
	router.HandleFunc("/listcategories", controller.GetAllCategory(db)).Methods("GET")

	log.Info("Starting Note APP on localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Crit("Unable to start web server", "error", err)
		os.Exit(1)
	}

}
