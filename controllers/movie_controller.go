package controllers

import (
	intrfc "lion-test/controllers/interfaces"
	"lion-test/models"
	intrfcRepo "lion-test/repositories/interfaces"
)

type MovieController struct {
	Repository intrfcRepo.MovieRepositoryInterface
}

func NewMovieController(repository intrfcRepo.MovieRepositoryInterface) intrfc.MovieControllerInterface {
	return &MovieController{
		Repository: repository,
	}
}

func (controller *MovieController) InsertMovieController(movieModel *models.InsertMovie) (int, error) {
	insertMovie, err := controller.Repository.InsertMovieRepository(movieModel)
	return insertMovie, err
}

func (controller *MovieController) GetMovieController(request *models.BaseRequestGetDataRepository) ([]models.GetMovieData, error) {
	listMovies, err := controller.Repository.GetMovieRepository(request)
	return listMovies, err
}

func (controller *MovieController) UpdateMovieController(movieModel *models.UpdateMovie) (int, error) {
	updateMovie, err := controller.Repository.UpdateMovieRepository(movieModel)
	return updateMovie, err
}

func (controller *MovieController) GetMovieDetailController(id, title string) (*models.GetMovieDetailData, error) {
	movieDetail, err := controller.Repository.GetMovieDetailRepository(id, title)
	return movieDetail, err
}

func (controller *MovieController) GetMovieDetailForUpdateController(id, title string) (*models.GetMovieDetailData, error) {
	movieDetail, err := controller.Repository.GetMovieDetailForUpdateRepository(id, title)
	return movieDetail, err
}

func (controller *MovieController) GetMostViewed() (*models.GetMostViewed, error) {
	movieDetail, err := controller.Repository.GetMostViewed()
	return movieDetail, err
}
