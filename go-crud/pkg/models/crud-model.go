package models

import (
	"github.com/TaiwoDevOps/go-crud/pkg/config"
	"github.com/jinzhu/gorm"
)

// create db instance
var db *gorm.DB

// create a Book model
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// open the database connection
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// creating DB functions for handling controllers query methods
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// use this function to find a book by book object
func (b *Book) GetBookById() *Book {
	db.First(&b, b.ID)
	return b
}

// use this function to find a book by id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	// create a local variable to hold the book
	var getBook Book
	// find the book using the id
	db := db.Where("ID=?", Id).Find(&getBook)
	// return the book and the db instance
	return &getBook, db
}

// Note: GetBookById are created twice but with different signatures, find out why

// use this function to delete the book from record
func DeleteBookById(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
