package models

import (
	"github.com/hashcott/go-project/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Initialize() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return getBook, db
}

func DeleteBook(Id int64) Book {
	var deletedBook Book
	_ = db.Where("ID=?", Id).Delete(&deletedBook)
	return deletedBook
}
