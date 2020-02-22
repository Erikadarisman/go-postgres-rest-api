package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Erikadarisman/go-rest-api/controllers"

	"github.com/Erikadarisman/go-rest-api/driver"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	godotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	defer db.Close()
	controller := controllers.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/listcategories", controller.GetAllCategory(db)).Methods("GET")
	http.Handle("/", router)

	fmt.Println("Connected to port 9999")
	log.Fatal(http.ListenAndServe(":9999", router))

}
