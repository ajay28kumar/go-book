package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-book/models"
	"go-book/repository/book"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {

}

var books [] models.Book
func logFatal (err error){
	if err != nil {
		log.Fatal(err)
	}
}
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)
		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		params:=mux.Vars(r) // we have id of books in map
		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		id, err := strconv.Atoi(params["id"])
		logFatal(err)
		book = bookRepo.GetBook(db, book,id)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB)  http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var book models.Book
		var bookID int
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookRepository.BookRepository{}
		bookID = bookRepo.AddBook(db, book)
		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db,book)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook (db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		params:=mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		id,err := strconv.Atoi(params["id"])
		logFatal(err)
		rowsUpdated := bookRepo.RemoveBook(db,id)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}