package repositories

import (
	"gorm.io/gorm"
	"lion-test/helper"
	"lion-test/models"
	intrfc "lion-test/repositories/interfaces"
	"strings"
)

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(DB *gorm.DB) intrfc.MovieRepositoryInterface {
	return &MovieRepository{
		DB: DB,
	}
}

func (repository *MovieRepository) InsertMovieRepository(movieModel *models.InsertMovie) (int, error) {
	insertMovie := repository.DB.Table("movies").Create(movieModel)
	if insertMovie.Error != nil {
		return 0, insertMovie.Error
	}

	return int(insertMovie.RowsAffected), insertMovie.Error
}

func (repository *MovieRepository) GetMovieRepository(request *models.BaseRequestGetDataRepository) ([]models.GetMovieData, error) {
	var listAllMovies []models.GetMovieData
	var totalCount int64
	query := repository.DB.Table("movies").Select("*, SEC_TO_TIME(duration) AS duration")
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	query = helper.SearchData(query, request.Search)
	offset := (request.Pagination.Page - 1) * request.Pagination.Size
	if err := query.Limit(request.Pagination.Size).Offset(offset).Find(&listAllMovies).Error; err != nil {
		return nil, err
	}

	return listAllMovies, nil
}

func (repository *MovieRepository) UpdateMovieRepository(movieModel *models.UpdateMovie) (int, error) {
	updateViewed := repository.DB.Table("movies").Where("id = ?", movieModel.Id).Save(&movieModel)
	if updateViewed.Error != nil {
		return 0, updateViewed.Error
	}

	return int(updateViewed.RowsAffected), updateViewed.Error
}

func (repository *MovieRepository) GetMovieDetailForUpdateRepository(id, title string) (*models.GetMovieDetailData, error) {
	var movieData *models.GetMovieDetailData

	query := repository.DB.Table("movies")
	if id != "" && title == "" {
		query = query.Select("*, SEC_TO_TIME(duration) AS duration").Where("id = ?", id).First(&movieData)
		if query.Error != nil {
			return nil, query.Error
		}
	} else if title != "" && id == "" {
		query = query.Select("*, SEC_TO_TIME(duration) AS duration").Where("title = ?", title).First(&movieData)
		if query.Error != nil {
			return nil, query.Error
		}
	}
	return movieData, nil
}

func (repository *MovieRepository) GetMovieDetailRepository(id, title string) (*models.GetMovieDetailData, error) {
	var movieData *models.GetMovieDetailData

	query := repository.DB.Table("movies")
	if id != "" && title == "" {
		query = query.Select("*, SEC_TO_TIME(duration) AS duration").Where("id = ?", id).First(&movieData)
		if query.Error != nil {
			return nil, query.Error
		}

		movieData.Viewed += 1

		duration, err := helper.ConvertTimeToSeconds(movieData.Duration)
		if err != nil {
			return nil, err
		}

		_, err = repository.UpdateMovieRepository(&models.UpdateMovie{
			InsertMovie: models.InsertMovie{
				Id:          movieData.Id,
				Title:       movieData.Title,
				Description: movieData.Description,
				Duration:    int64(duration),
				Artists:     movieData.Artists,
				Genres:      movieData.Genres,
				WatchUrl:    movieData.WatchUrl,
				Viewed:      movieData.Viewed,
				Voted:       movieData.Voted,
			},
		})

		genres := strings.Split(movieData.Genres, ",")
		for _, genre := range genres {
			_, err = repository.UpdateViewedGenreRepository(&models.UpdateGenre{
				InsertGenre: models.InsertGenre{
					Genre: genre,
				},
			})
		}

		if err != nil {
			return nil, err
		}
	} else if title != "" && id == "" {
		query = query.Select("*, SEC_TO_TIME(duration) AS duration").Where("title = ?", title).First(&movieData)
		if query.Error != nil {
			return nil, query.Error
		}
	}
	return movieData, nil
}

func (repository *MovieRepository) GetMostViewed() (*models.GetMostViewed, error) {
	var result models.GetMostViewed
	var movieData []models.GetMovieDetailData
	getMovieDetail := repository.DB.Table("movies").Where("viewed <> 0").Order("viewed desc").Limit(5).Find(&movieData)
	if getMovieDetail.Error != nil {
		return &result, getMovieDetail.Error
	}

	result.Data = movieData
	return &result, nil
}

func (repository *MovieRepository) UpdateViewedGenreRepository(movieModel *models.UpdateGenre) (int, error) {
	var genreData models.GetGenreData
	movieModel.Genre = strings.Replace(movieModel.Genre, "'", "", -1)
	getDataGenre := repository.DB.Table("genres").Where("genre = ?", movieModel.Genre)
	if getDataGenre.Error != nil {
		return 0, getDataGenre.Error
	}

	getGenreDetail := getDataGenre.First(&genreData)
	if getGenreDetail.Error != nil {
		return 0, getGenreDetail.Error
	}
	movieModel.Viewed = genreData.Viewed + 1
	getDataGenre.UpdateColumn("viewed", movieModel.Viewed)

	return int(getDataGenre.RowsAffected), getDataGenre.Error
}
