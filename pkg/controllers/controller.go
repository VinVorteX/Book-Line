package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vinayak3010/Book-Line/pkg/models"
	"github.com/vinayak3010/Book-Line/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := uuid.Parse(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	book, _ := models.GetBookById(bookID)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)

	NewBook = *CreateBook.CreateBook()

	res, _ := json.Marshal(NewBook)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := uuid.Parse(params["Id"])

	if err != nil {
		log.Fatal(err)
	}

	book := models.DeleteBook(bookID)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}

	utils.ParseBody(r, updateBook)

	params := mux.Vars(r)

	bookID, err := uuid.Parse(params["Id"])

	if err != nil {
		log.Fatal(err)
	}

	book, db := models.GetBookById(bookID)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}
