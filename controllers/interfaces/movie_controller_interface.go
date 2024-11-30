package controller

import "lion-test/models"

type MovieControllerInterface interface {
	InsertMovieController(movieModel *models.InsertMovie) (int, error)
	GetMovieController(request *models.BaseRequestGetDataRepository) ([]models.GetMovieData, error)
	UpdateMovieController(movieModel *models.UpdateMovie) (int, error)
	GetMovieDetailController(id, title string) (*models.GetMovieDetailData, error)
	GetMovieDetailForUpdateController(id, title string) (*models.GetMovieDetailData, error)
	GetMostViewed() (*models.GetMostViewed, error)
}
