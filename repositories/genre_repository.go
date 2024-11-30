package repositories

import (
	"gorm.io/gorm"
	"lion-test/helper"
	"lion-test/models"
	intrfc "lion-test/repositories/interfaces"
)

type GenreRepository struct {
	DB *gorm.DB
}

func NewGenreRepository(DB *gorm.DB) intrfc.GenreRepositoryInterface {
	return &GenreRepository{
		DB: DB,
	}
}

func (repository *GenreRepository) InsertGenreRepository(genreModel *models.InsertGenre) (int, error) {
	insertMovie := repository.DB.Table("genres").Create(genreModel)
	if insertMovie.Error != nil {
		return 0, insertMovie.Error
	}

	return int(insertMovie.RowsAffected), insertMovie.Error
}

func (repository *GenreRepository) GetGenreRepository(request *models.BaseRequestGetDataRepository) ([]models.GetGenreData, error) {
	var listAllGenre []models.GetGenreData
	var totalCount int64
	query := repository.DB.Table("genres")
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	query = helper.SearchData(query, request.Search)
	offset := (request.Pagination.Page - 1) * request.Pagination.Size
	if err := query.Limit(request.Pagination.Size).Offset(offset).Find(&listAllGenre).Error; err != nil {
		return nil, err
	}

	return listAllGenre, nil
}

func (repository *GenreRepository) GetGenreDetailRepository(genre string) (*models.GetGenreData, error) {
	var genreData *models.GetGenreData

	query := repository.DB.Table("genres").Where("genre = ?", genre).First(&genreData)
	if query.Error != nil {
		return nil, query.Error
	}

	return genreData, nil
}

func (repository *GenreRepository) GetMostViewed() (*models.GetMostViewed, error) {
	var result models.GetMostViewed
	var genreData []models.GetGenreData
	getMovieDetail := repository.DB.Table("genres").Where("viewed <> 0").Order("viewed desc").Limit(5).Find(&genreData)
	if getMovieDetail.Error != nil {
		return &result, getMovieDetail.Error
	}

	result.Data = genreData
	return &result, nil
}
