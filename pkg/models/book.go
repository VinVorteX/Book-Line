package models

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/vinayak3010/Book-Line/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Id 		uuid.UUID `gorm:"type:uuid;primary_key;"json:"id"`
	Name        string `gorm:"type:varchar(100)"json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
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

func GetBookById(Id uuid.UUID) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id uuid.UUID) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}
