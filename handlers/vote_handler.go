package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	controller "lion-test/controllers/interfaces"
	"lion-test/helper"
	"lion-test/models"
	"net/http"
)

type VoteHandler struct {
	Controller      controller.VoteControllerInterface
	MovieController controller.MovieControllerInterface
	Validator       *validator.Validate
}

func NewVoteHandler(controller controller.VoteControllerInterface, movieController controller.MovieControllerInterface) *VoteHandler {
	validate := validator.New()
	return &VoteHandler{
		Controller:      controller,
		MovieController: movieController,
		Validator:       validate,
	}
}

func (handler VoteHandler) UpdateVoteHandler(ctx *gin.Context) {
	var voteDTO models.VotesDTO

	if err := ctx.BindJSON(&voteDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&voteDTO); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	username, _ := ctx.Get("username")

	if voteDTO.UserUsername != username {
		ctx.JSON(http.StatusUnauthorized, models.BaseResponse{Status: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	updateVoteDTO := models.VotesDTO{
		MovieId:      voteDTO.MovieId,
		UserUsername: voteDTO.UserUsername,
	}

	updateVote, err := handler.Controller.UpdateVoteController(&updateVoteDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if updateVote == "insert" {
		getMovie, errGetMovie := handler.MovieController.GetMovieDetailForUpdateController(updateVoteDTO.MovieId, "")
		if errGetMovie != nil {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errGetMovie.Error()}})
			return
		}

		duration, err := helper.ConvertTimeToSeconds(getMovie.Duration)
		if err != nil {
			return
		}

		movieUpdate := models.UpdateMovie{
			InsertMovie: models.InsertMovie{
				Id:          getMovie.Id,
				Title:       getMovie.Title,
				Description: getMovie.Description,
				Duration:    int64(duration),
				Artists:     getMovie.Artists,
				Genres:      getMovie.Genres,
				WatchUrl:    getMovie.WatchUrl,
				Viewed:      getMovie.Viewed,
				Voted:       getMovie.Voted + 1,
			},
		}

		updateVoteMovie, errUpdateVoteMovie := handler.MovieController.UpdateMovieController(&movieUpdate)
		if errUpdateVoteMovie != nil {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errUpdateVoteMovie.Error()}})
			return
		}

		ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateVoteMovie}})
	} else if updateVote == "delete" {
		getMovie, errGetMovie := handler.MovieController.GetMovieDetailForUpdateController(updateVoteDTO.MovieId, "")
		if errGetMovie != nil {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errGetMovie.Error()}})
			return
		}

		duration, err := helper.ConvertTimeToSeconds(getMovie.Duration)
		if err != nil {
			return
		}

		movieUpdate := models.UpdateMovie{
			InsertMovie: models.InsertMovie{
				Id:          getMovie.Id,
				Title:       getMovie.Title,
				Description: getMovie.Description,
				Duration:    int64(duration),
				Artists:     getMovie.Artists,
				Genres:      getMovie.Genres,
				WatchUrl:    getMovie.WatchUrl,
				Viewed:      getMovie.Viewed,
				Voted:       getMovie.Voted - 1,
			},
		}

		updateVoteMovie, errUpdateVoteMovie := handler.MovieController.UpdateMovieController(&movieUpdate)
		if errUpdateVoteMovie != nil {
			ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errUpdateVoteMovie.Error()}})
			return
		}

		ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateVoteMovie}})
	}
}

func (handler VoteHandler) CheckVoteHandler(ctx *gin.Context) {
	var checkVoteParam models.Votes

	if err := ctx.BindJSON(&checkVoteParam); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := handler.Validator.Struct(&checkVoteParam); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	checkVote, err := handler.Controller.CheckVoteController(&checkVoteParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": checkVote}})
}
