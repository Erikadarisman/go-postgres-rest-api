package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Erikadarisman/go-rest-api/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Controller this
type Controller struct{}

//GetAllCategory - method to handle get categories
func (c Controller) GetAllCategory(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category models.Category
		var arrCat []models.Category
		var response models.Response

		rows, err := db.Query(context.Background(), "Select id,name from public.category")
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

		response.Status = 0
		response.Message = "Success"
		response.Data = arrCat

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
