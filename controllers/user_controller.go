package controllers

import (
	intrfc "lion-test/controllers/interfaces"
	"lion-test/models"
	intrfcRepo "lion-test/repositories/interfaces"
)

type UserController struct {
	Repository intrfcRepo.UserRepositoryInterface
}

func NewUserController(repository intrfcRepo.UserRepositoryInterface) intrfc.UserControllerInterface {
	return &UserController{
		Repository: repository,
	}
}

func (controller *UserController) InsertUserController(User *models.InsertUser) (int, error) {
	insertUser, err := controller.Repository.InsertUserRepository(User)
	return insertUser, err
}

func (controller *UserController) GetUserByUsernameController(User *models.GetUserByUsername) (models.GetUserData, error) {
	userDetail, err := controller.Repository.GetUserByUsernameRepository(User)
	return userDetail, err
}
