package interfaces

import "lion-test/models"

type GenreRepositoryInterface interface {
	InsertGenreRepository(movieModel *models.InsertGenre) (int, error)
	GetGenreRepository(request *models.BaseRequestGetDataRepository) ([]models.GetGenreData, error)
	GetGenreDetailRepository(genre string) (*models.GetGenreData, error)
	GetMostViewed() (*models.GetMostViewed, error)
}
