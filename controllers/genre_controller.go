package controllers

import (
	intrfc "lion-test/controllers/interfaces"
	"lion-test/models"
	intrfcRepo "lion-test/repositories/interfaces"
)

type GenreController struct {
	Repository intrfcRepo.GenreRepositoryInterface
}

func NewGenreController(repository intrfcRepo.GenreRepositoryInterface) intrfc.GenreControllerInterface {
	return &GenreController{
		Repository: repository,
	}
}

func (controller *GenreController) InsertGenreController(genreModel *models.InsertGenre) (int, error) {
	insertGenre, err := controller.Repository.InsertGenreRepository(genreModel)
	return insertGenre, err
}

func (controller *GenreController) GetGenreController(request *models.BaseRequestGetDataRepository) ([]models.GetGenreData, error) {
	listGenres, err := controller.Repository.GetGenreRepository(request)
	return listGenres, err
}

func (controller *GenreController) GetGenreDetailController(genre string) (*models.GetGenreData, error) {
	genreDetail, err := controller.Repository.GetGenreDetailRepository(genre)
	return genreDetail, err
}

func (controller *GenreController) GetMostViewed() (*models.GetMostViewed, error) {
	genreDetail, err := controller.Repository.GetMostViewed()
	return genreDetail, err
}
