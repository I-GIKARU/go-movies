package models

import (
	"gomysql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Image       string `json:"image"` // ✅ Image URL field
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	err := db.Create(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		return []Book{}, err
	}
	return books, nil
}

func GetBookById(id int) (Book, error) {
	var book Book
	err := db.First(&book, id).Error
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func DeleteBook(id int, book Book) (Book, error) {
	err := db.Delete(&book, id).Error
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func UpdateBook(id int, updatedData Book) (Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return Book{}, err
	}

	if updatedData.Title != "" {
		book.Title = updatedData.Title
	}
	if updatedData.Author != "" {
		book.Author = updatedData.Author
	}
	if updatedData.Publication != "" {
		book.Publication = updatedData.Publication
	}
	if updatedData.Image != "" {
		book.Image = updatedData.Image
	}

	err := db.Save(&book).Error
	return book, err
}
