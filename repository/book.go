package repository

import (
	"github.com/mahendrafathan/go-api-boilerplate/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(book *[]models.Book) *gorm.DB
	Create(models.Book) *gorm.DB
	FindOne(isbn string, book *models.Book) *gorm.DB
	Delete(isbn string, book *models.Book) *gorm.DB
}
