package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	gorm.Model
	id_book	uint	`gorm:"primaryKey"`
	title	string
	author	string
	lang	string
	status	string
}

var db *gorm.DB
var err error

var (
	books = []Book{
		{title: "Atomic Habit", author: "James Clear", lang: "en", status: "in progress"},
		{title: "Learning Go", author: "Jon Bodner", lang: "en", status: "in progress"},
	}
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(&books)

}

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost  port=5432 user=postgres dbname=book_db sslmode=disable password=example")

	if err != nil {
		log.Println(fmt.Sprintf("%s", err))
	}

	defer db.Close()

	db.AutoMigrate(&Book{})

	for index := range books {
		db.Create(&books[index])
	}

	router.HandleFunc("/books", GetBooks).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8081", handler))

}
