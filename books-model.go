package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getBook(db *gorm.DB, id int) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func getBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		log.Fatalf("Error to get book: %v", result.Error)
	}

	return books
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Model(&Book{}).Where("id = ?", book.ID).Updates(book)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no book found with ID %d", book.ID)
	}
	return nil
}

func deleteBook(db *gorm.DB, id int) error {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no book found with ID %d", id)
	}
	return nil
}

func searchBook(db *gorm.DB, bookName string) []Book {
	var books []Book

	result := db.Where("name = ?", bookName).Order("price desc").Find(&books)
	if result.Error != nil {
		log.Fatalf("Search book failed: %v", result.Error)
	}
	return books
}
