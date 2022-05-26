package repository

import (
	"github.com/mahendrafathan/go-api-boilerplate/models"

	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(database *gorm.DB) UserRepository {
	return &userRepoImpl{
		db: database,
	}
}

func (u *userRepoImpl) FindByUsernameOrEmail(cred string) *gorm.DB {
	return u.db.Where("username = ? OR email = ?", cred, cred)
}

func (u *userRepoImpl) Create(user models.User) *gorm.DB {
	return u.db.Create(&user)
}
