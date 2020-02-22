package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Erikadarisman/go-rest-api/driver"
	"github.com/Erikadarisman/go-rest-api/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	godotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/getcategories", getAllCategory).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 9999")
	log.Fatal(http.ListenAndServe(":9999", router))

	// log.Println("Server started successfully on port 8000.!!!")
}

func getAllCategory(w http.ResponseWriter, r *http.Request) {

	var category models.Category
	var arrCat []models.Category
	var response models.Response

	rows, err := db.Query("Select id,name from public.category")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCat = append(arrCat, category)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrCat

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
