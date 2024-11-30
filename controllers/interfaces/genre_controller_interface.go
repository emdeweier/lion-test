package controller

import "lion-test/models"

type GenreControllerInterface interface {
	InsertGenreController(movieModel *models.InsertGenre) (int, error)
	GetGenreController(request *models.BaseRequestGetDataRepository) ([]models.GetGenreData, error)
	GetGenreDetailController(genre string) (*models.GetGenreData, error)
	GetMostViewed() (*models.GetMostViewed, error)
}
