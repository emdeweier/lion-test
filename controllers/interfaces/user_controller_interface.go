package controller

import "lion-test/models"

type UserControllerInterface interface {
	InsertUserController(User *models.InsertUser) (int, error)
	GetUserByUsernameController(User *models.GetUserByUsername) (models.GetUserData, error)
}
