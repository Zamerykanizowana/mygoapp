package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struck {
	gorm.Model
	id_book	unit	`gorm:"primaryKey"`
	title	string
	author	string
	lang	string
	status	string
}

var db *gorn.DB
var err error

var (
	books = []Book{
		{title: "Atomic Habit", author: "James Clear", lang: "en", status: "in progress"},
		{title: "Learning Go", author: "Jon Bodner", lang: "en", status: "in progress"},
	}
)

func main() {
	router := mux.Router()

	db, err = gorm.Open("postgres", "host=localhost  port=5432 user=postgres dbname=book_db sslmode=disable password=example")

	if err != nil {
		log.Println("Problem with connecting database")
	}

	defer db.Close()

	db.AutoMigrate(&Book{})

	for index := books {
		db.Create(&books[index])
	}

	router.HandleFunc("/books", GetBooks).Method("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8081", handel))

}
