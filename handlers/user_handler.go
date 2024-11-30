package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	controller "lion-test/controllers/interfaces"
	"lion-test/models"
	"net/http"
	"time"
)

type UserHandler struct {
	Controller controller.UserControllerInterface
	Validator  *validator.Validate
}

func NewUserHandler(controller controller.UserControllerInterface) *UserHandler {
	validate := validator.New()
	return &UserHandler{
		Controller: controller,
		Validator:  validate,
	}
}

func (handler *UserHandler) InsertUserHandler(ctx *gin.Context) {
	var userParamInsert models.InsertUser

	if err := ctx.BindJSON(&userParamInsert); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&userParamInsert); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userParamInsert.Password), bcrypt.DefaultCost)

	newUserInsert := models.InsertUser{
		Username:   userParamInsert.Username,
		Name:       userParamInsert.Name,
		Password:   string(hashedPassword),
		Role:       "user",
		InsertDate: time.Now(),
		Address:    userParamInsert.Address,
	}

	insertUser, err := handler.Controller.InsertUserController(&newUserInsert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": insertUser}})
}

func (handler *UserHandler) GenerateAdminHandler(ctx *gin.Context) {
	var userParamInsert models.InsertUser

	if err := ctx.BindJSON(&userParamInsert); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&userParamInsert); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userParamInsert.Password), bcrypt.DefaultCost)

	newUserInsert := models.InsertUser{
		Username:   userParamInsert.Username,
		Name:       userParamInsert.Name,
		Password:   string(hashedPassword),
		Role:       "admin",
		InsertDate: time.Now(),
		Address:    userParamInsert.Address,
	}

	insertUser, err := handler.Controller.InsertUserController(&newUserInsert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": insertUser}})
}

func (handler *UserHandler) GetUserByUsernameHandler(ctx *gin.Context) {
	var getUserByUsernameParam models.GetUserByUsername

	if err := ctx.BindJSON(&getUserByUsernameParam); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&getUserByUsernameParam); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	getUser, err := handler.Controller.GetUserByUsernameController(&getUserByUsernameParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getUser}})
}
