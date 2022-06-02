package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashcott/go-project/go-bookstore/pkg/models"
	"github.com/hashcott/go-project/go-bookstore/pkg/utils"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, _ *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetail, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	deletedBookDetail := models.DeleteBook(ID)

	res, _ := json.Marshal(deletedBookDetail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetail, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetail.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetail.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetail.Publication = updateBook.Publication
	}

	db.Save(&bookDetail)

	res, _ := json.Marshal(bookDetail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
