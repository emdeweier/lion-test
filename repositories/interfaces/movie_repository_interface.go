package interfaces

import "lion-test/models"

type MovieRepositoryInterface interface {
	InsertMovieRepository(movieModel *models.InsertMovie) (int, error)
	GetMovieRepository(request *models.BaseRequestGetDataRepository) ([]models.GetMovieData, error)
	UpdateMovieRepository(movieModel *models.UpdateMovie) (int, error)
	GetMovieDetailRepository(id, title string) (*models.GetMovieDetailData, error)
	GetMovieDetailForUpdateRepository(id, title string) (*models.GetMovieDetailData, error)
	GetMostViewed() (*models.GetMostViewed, error)
	UpdateViewedGenreRepository(movieModel *models.UpdateGenre) (int, error)
}
