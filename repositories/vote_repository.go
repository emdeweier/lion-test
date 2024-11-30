package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"lion-test/models"
	intrfc "lion-test/repositories/interfaces"
	"strings"
)

type VoteRepository struct {
	DB *gorm.DB
}

func NewVoteRepository(DB *gorm.DB) intrfc.VoteRepositoryInterface {
	return &VoteRepository{
		DB: DB,
	}
}

func (repository VoteRepository) UpdateVoteRepository(voteModel *models.VotesDTO) (string, error) {
	alreadyVote, checkVoteErr := repository.CheckVoteRepository(&models.Votes{
		MovieId:      voteModel.MovieId,
		UserUsername: voteModel.UserUsername,
	})

	if checkVoteErr != nil {
		return "", checkVoteErr
	}

	if alreadyVote {
		deleteVote := repository.DB.Table("votes").Where("user_username = ? AND movie_id = ?", voteModel.UserUsername, voteModel.MovieId).Delete(voteModel)
		if deleteVote.Error != nil {
			return "", deleteVote.Error
		}

		return "delete", deleteVote.Error
	}

	insertVote := repository.DB.Table("votes").Create(voteModel)
	if insertVote.Error != nil {
		return "", insertVote.Error
	}

	return "insert", insertVote.Error
}

func (repository VoteRepository) CheckVoteRepository(voteModel *models.Votes) (bool, error) {
	var voteData *models.VotesDTO
	checkStatus := repository.DB.Table("votes").Select("*").Where("movie_id = ? and user_username = ?", voteModel.MovieId, strings.ToUpper(voteModel.UserUsername)).First(&voteData)
	if checkStatus.Error != nil && fmt.Sprintf("%s", checkStatus.Error) != "record not found" {
		return false, checkStatus.Error
	}

	if fmt.Sprintf("%s", checkStatus.Error) == "record not found" {
		return false, nil
	}

	return true, nil
}
