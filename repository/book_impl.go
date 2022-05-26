package repository

import (
	"github.com/mahendrafathan/go-api-boilerplate/models"

	"gorm.io/gorm"
)

type bookRepoImpl struct {
	db *gorm.DB
}

func NewBookRepoImpl(database *gorm.DB) BookRepository {
	return &bookRepoImpl{
		db: database,
	}
}

func (b *bookRepoImpl) FindAll(book *[]models.Book) *gorm.DB {
	return b.db.Find(&book)
}
func (b *bookRepoImpl) Create(book models.Book) *gorm.DB {
	return b.db.Create(&book)
}

func (b *bookRepoImpl) FindOne(isbn string, book *models.Book) *gorm.DB {
	return b.db.Where("isbn = ?", isbn).Assign(&book)
}
func (b *bookRepoImpl) Delete(isbn string, book *models.Book) *gorm.DB {
	return b.db.Where("isbn = ?", isbn).Delete(&book)
}
