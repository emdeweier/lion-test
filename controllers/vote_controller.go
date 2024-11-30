package controllers

import (
	intrfc "lion-test/controllers/interfaces"
	"lion-test/models"
	intrfcRepo "lion-test/repositories/interfaces"
)

type VoteController struct {
	Repository intrfcRepo.VoteRepositoryInterface
}

func NewVoteController(repository intrfcRepo.VoteRepositoryInterface) intrfc.VoteControllerInterface {
	return &VoteController{
		Repository: repository,
	}
}

func (controller VoteController) UpdateVoteController(voteModel *models.VotesDTO) (string, error) {
	updateVote, err := controller.Repository.UpdateVoteRepository(voteModel)
	return updateVote, err
}

func (controller VoteController) CheckVoteController(voteModel *models.Votes) (bool, error) {
	checkVote, err := controller.Repository.CheckVoteRepository(voteModel)
	return checkVote, err
}
