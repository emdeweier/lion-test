package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	controller "lion-test/controllers/interfaces"
	"lion-test/models"
	"net/http"
)

type GenreHandler struct {
	Controller controller.GenreControllerInterface
	Validator  *validator.Validate
}

func NewGenreHandler(controller controller.GenreControllerInterface) *GenreHandler {
	validate := validator.New()
	return &GenreHandler{
		Controller: controller,
		Validator:  validate,
	}
}

func (handler *GenreHandler) InsertGenreHandler(ctx *gin.Context) {
	var genreParamInsert models.InsertGenre

	if err := ctx.BindJSON(&genreParamInsert); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&genreParamInsert); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	getGenre, err := handler.Controller.GetGenreDetailController(genreParamInsert.Genre)
	if err != nil && err.Error() != "record not found" {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	} else if getGenre != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "This genre already exist."}})
		return
	}

	newGenreInsert := models.InsertGenre{
		Genre:  genreParamInsert.Genre,
		Viewed: 0,
	}

	insertMovie, err := handler.Controller.InsertGenreController(&newGenreInsert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": insertMovie}})
}

func (handler *GenreHandler) GetGenreHandler(ctx *gin.Context) {
	var request models.BaseRequestGetDataHandler
	var requestRepository models.BaseRequestGetDataRepository

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&request); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	if request.Search != "" {
		search := map[string]interface{}{"genre": request.Search}
		requestRepository.Search = search
	}

	requestRepository.Pagination = request.Pagination

	getGenre, err := handler.Controller.GetGenreController(&requestRepository)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getGenre}})
}

func (handler *GenreHandler) GetGenreDetailHandler(ctx *gin.Context) {
	var getGenreData models.GetGenreByName

	if err := ctx.BindJSON(&getGenreData); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&getGenreData); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	getMovie, err := handler.Controller.GetGenreDetailController(getGenreData.Genre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getMovie}})
}

func (handler *GenreHandler) GetMostViewedHandler(ctx *gin.Context) {
	getGenre, err := handler.Controller.GetMostViewed()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getGenre}})
}
