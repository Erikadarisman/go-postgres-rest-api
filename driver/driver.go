package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("APP_DB_URL"))
	if err != nil {
		log.Fatalln("Error while connecting to DB")
		os.Exit(1)
	}

	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping() //if ping does not retunr anything that means connection is established successfully
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}

// func connectDB() *sql.DB {
// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
// 		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USERDB"), os.Getenv("PASSWORDDB"), os.Getenv("DBNAME"))
// 	// host, port, user, password, dbname

// 	conn, err := sql.Open("postgres", psqlInfo)

// }
