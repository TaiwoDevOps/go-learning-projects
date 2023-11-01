package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TaiwoDevOps/go-crud/pkg/models"
	"github.com/TaiwoDevOps/go-crud/pkg/utils"
	"github.com/gorilla/mux"
)

// create the books model variable
var NewBook models.Book

// create the get books controller function
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	// convert the data to JSON
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create the get book by id controller function
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// call the parse body function to parse the request body
	vars := mux.Vars(r)
	// call the get book by id function passed to the request body
	bookId := vars["bookId"]
	// id is coming as int, parse it as string
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create the create book controller function
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// create a new Book instance variable
	book := &models.Book{}
	// call the parse body function to parse the request body and make it readable for the DB
	utils.ParseBody(r, book)
	// call the create book function
	b := book.CreateBook()
	// convert the data to JSON and return the response
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create a delete book controller function
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create an update book controller function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// create a new Book instance variable
	var updateBook = &models.Book{}
	// call the parse body function to parse the request body and make it readable for the DB
	utils.ParseBody(r, updateBook)
	// create a variable to hold the body of the request
	vars := mux.Vars(r)
	// read the request body and pick the book id
	bookId := vars["bookId"]
	// convert the book id to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call the get book by id function and ensure the book exists
	bookDetails, db := models.GetBookById(ID)
	// update the book details from the request body
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// save  the book details in memory address
	db.Save(&bookDetails)
	// convert the data to JSON and send response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
