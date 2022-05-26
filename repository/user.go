package repository

import (
	"github.com/mahendrafathan/go-api-boilerplate/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsernameOrEmail(username string) *gorm.DB
	Create(models.User) *gorm.DB
}
