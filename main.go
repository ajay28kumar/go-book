package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go-book/controllers"
	"go-book/driver"
	"go-book/models"
	"log"
	"net/http"
)
var books [] models.Book
var db *sql.DB
func init()  {
	gotenv.Load()
}



func main() {
	db=driver.ConnectDB()
	router:= mux.NewRouter()
	controller := controllers.Controller{}
	router.HandleFunc("/books",controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}",controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books",controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books",controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}",controller.RemoveBook(db)).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8000", router))
}
