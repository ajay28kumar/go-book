package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

var books [] Book

func main() {
	router:= mux.NewRouter()
	books = append(books, Book{0, "Golang pointers", "Mr Google", "2010"})

	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request)  {
	b := json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request)  {
	p:=mux.Vars(r)
	log.Println()
}
func addBook(w http.ResponseWriter, r *http.Request)  {
	log.Println("addBook")
}
func updateBook(w http.ResponseWriter, r *http.Request)  {
	log.Println("updateBook")
}
func removeBook(w http.ResponseWriter, r *http.Request)  {
	log.Println("removeBook")
}