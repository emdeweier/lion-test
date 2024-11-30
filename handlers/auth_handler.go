package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"lion-test/models"
	intrfc "lion-test/repositories/interfaces"
	"lion-test/utils"
	"net/http"
)

type AuthHandler struct {
	UserRepository intrfc.UserRepositoryInterface
	Validator      *validator.Validate
}

func NewAuthHandler(userRepository intrfc.UserRepositoryInterface) *AuthHandler {
	validate := validator.New()
	return &AuthHandler{
		UserRepository: userRepository,
		Validator:      validate,
	}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	var loginParam models.Credential

	if err := ctx.BindJSON(&loginParam); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&loginParam); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	getUser, err := handler.UserRepository.GetUserByUsernameRepository(&models.GetUserByUsername{Username: loginParam.Username})
	if err != nil || bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginParam.Password)) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(getUser.Username, getUser.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": token}})
}
