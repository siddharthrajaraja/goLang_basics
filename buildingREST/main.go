package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model) -- Class for OOPS in Go lang == ES6

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Init books var as a slice Book Struct
var books []Book

// Author Struct
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Get all Books :
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books { // item iterator
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func delBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)

			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(
		books,
		Book{
			ID:    "1",
			Isbn:  "4556",
			Title: "Book one",
			Author: &Author{
				FirstName: "Siddharth", LastName: "Raja"},
		},
	)

	books = append(
		books,
		Book{
			ID:    "2",
			Isbn:  "4557",
			Title: "Book two",
			Author: &Author{
				FirstName: "Tushar", LastName: "Raja"},
		},
	)

	// Route Handlers

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/add-book", createBook).Methods("POST")
	r.HandleFunc("/api/del-book/{id}", delBook).Methods("DELETE")
	r.HandleFunc("/api/update-book/{id}", updateBook).Methods("PUT")

	// Run server
	log.Fatal(http.ListenAndServe(":8000", r))
}
