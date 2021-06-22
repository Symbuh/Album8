package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"ID"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init books variable as a slice Book struct
// You need to define the length of arrays in go but slice has variable length

// Declaring the collection of books as a slice
var books []Book

func uploadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Content-Type", "application/json")
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	// Loop through books and find with id

	for _, item := range books {
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
	book.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID - not safe (could generate same ID)
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(100000))
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			// Weird syntax for remaking our books array without the specific index??
			// Really don't understand how it works but it will get rid of the book
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

func main() {
	// Initialize router
	// := specifies type inference
	r := mux.NewRouter()

	// Route handlers / endpoints

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})

	books = append(books, Book{ID: "2", Isbn: "847654", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Deer"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
