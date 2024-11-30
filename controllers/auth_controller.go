package controllers

import (
	intrfc "lion-test/controllers/interfaces"
	"lion-test/models"
	intrfcRepo "lion-test/repositories/interfaces"
)

type AuthController struct {
	Repository intrfcRepo.UserRepositoryInterface
}

func NewAuthController(repository intrfcRepo.UserRepositoryInterface) intrfc.AuthControllerInterface {
	return &AuthController{
		Repository: repository,
	}
}

func (controller AuthController) Login(Credential *models.Credential) (string, error) {
	//TODO implement me
	panic("implement me")
}
