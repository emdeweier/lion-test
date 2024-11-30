package controller

import "lion-test/models"

type AuthControllerInterface interface {
	Login(Credential *models.Credential) (string, error)
}
