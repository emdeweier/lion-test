package interfaces

import "lion-test/models"

type UserRepositoryInterface interface {
	InsertUserRepository(User *models.InsertUser) (int, error)
	GetUserByUsernameRepository(User *models.GetUserByUsername) (models.GetUserData, error)
}
