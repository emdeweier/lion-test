package repositories

import (
	"gorm.io/gorm"
	"lion-test/models"
	intrfc "lion-test/repositories/interfaces"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) intrfc.UserRepositoryInterface {
	return &UserRepository{
		DB: DB,
	}
}

func (repository UserRepository) InsertUserRepository(User *models.InsertUser) (int, error) {
	insertUser := repository.DB.Table("users").Create(User)
	if insertUser.Error != nil {
		return 0, insertUser.Error
	}

	return int(insertUser.RowsAffected), insertUser.Error
}

func (repository UserRepository) GetUserByUsernameRepository(User *models.GetUserByUsername) (models.GetUserData, error) {
	var userDetail models.GetUserData
	getUserDetail := repository.DB.Table("users").Select("*").Where("username = ?", User.Username).First(&userDetail)
	if getUserDetail.Error != nil {
		return models.GetUserData{}, getUserDetail.Error
	}

	return userDetail, getUserDetail.Error
}
