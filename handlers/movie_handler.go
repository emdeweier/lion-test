package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	controller "lion-test/controllers/interfaces"
	"lion-test/helper"
	"lion-test/models"
	"net/http"
)

type MovieHandler struct {
	Controller controller.MovieControllerInterface
	Validator  *validator.Validate
}

func NewMovieHandler(controller controller.MovieControllerInterface) *MovieHandler {
	validate := validator.New()
	return &MovieHandler{
		Controller: controller,
		Validator:  validate,
	}
}

func (handler *MovieHandler) InsertMovieHandler(ctx *gin.Context) {
	var movieParamInsert models.InsertMovie

	if err := ctx.BindJSON(&movieParamInsert); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&movieParamInsert); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	getMovie, err := handler.Controller.GetMovieDetailForUpdateController("", movieParamInsert.Title)
	if err != nil && err.Error() != "record not found" {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	} else if getMovie != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "This movie already exist."}})
		return
	}

	newMovieInsert := models.InsertMovie{
		Id:          helper.GenerateRandomString(8),
		Title:       movieParamInsert.Title,
		Description: movieParamInsert.Description,
		Duration:    movieParamInsert.Duration,
		Artists:     movieParamInsert.Artists,
		Genres:      movieParamInsert.Genres,
		WatchUrl:    movieParamInsert.WatchUrl,
		Viewed:      0,
		Voted:       0,
	}

	insertMovie, err := handler.Controller.InsertMovieController(&newMovieInsert)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": insertMovie}})
}

func (handler *MovieHandler) UpdateMovieHandler(ctx *gin.Context) {
	var movieParamUpdate models.UpdateMovie

	if err := ctx.BindJSON(&movieParamUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&movieParamUpdate); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	getMovieDetail, err := handler.Controller.GetMovieDetailForUpdateController(movieParamUpdate.Id, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	} else if getMovieDetail == nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "record not found"}})
		return
	}

	if movieParamUpdate.Title != getMovieDetail.Title {
		getMovie, err := handler.Controller.GetMovieDetailForUpdateController("", movieParamUpdate.Title)
		if err != nil && err.Error() != "record not found" {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		} else if getMovie != nil {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "This movie already exist."}})
			return
		}
	}

	movieUpdate := models.UpdateMovie{
		InsertMovie: models.InsertMovie{
			Id:          movieParamUpdate.Id,
			Title:       movieParamUpdate.Title,
			Description: movieParamUpdate.Description,
			Duration:    movieParamUpdate.Duration,
			Artists:     movieParamUpdate.Artists,
			Genres:      movieParamUpdate.Genres,
			WatchUrl:    movieParamUpdate.WatchUrl,
			Viewed:      getMovieDetail.Viewed,
			Voted:       getMovieDetail.Voted,
		},
	}

	insertMovie, err := handler.Controller.UpdateMovieController(&movieUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": insertMovie}})
}

func (handler *MovieHandler) GetMovieHandler(ctx *gin.Context) {
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
		search := map[string]interface{}{
			"title":       request.Search,
			"description": request.Search,
			"artists":     request.Search,
			"genres":      request.Search,
		}
		requestRepository.Search = search
	}

	requestRepository.Pagination = request.Pagination

	getMovie, err := handler.Controller.GetMovieController(&requestRepository)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getMovie}})
}

func (handler *MovieHandler) GetMovieDetailHandler(ctx *gin.Context) {
	var getMovieData models.GetMovieById

	if err := ctx.BindJSON(&getMovieData); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&getMovieData); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	role, _ := ctx.Get("role")

	if role == "admin" {
		getMovie, err := handler.Controller.GetMovieDetailForUpdateController(getMovieData.Id, "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getMovie}})
		return
	}

	getMovie, err := handler.Controller.GetMovieDetailController(getMovieData.Id, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getMovie}})
}

func (handler *MovieHandler) GetMostViewedHandler(ctx *gin.Context) {
	getMovie, err := handler.Controller.GetMostViewed()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": getMovie}})
}
