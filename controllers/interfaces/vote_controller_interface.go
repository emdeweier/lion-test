package controller

import "lion-test/models"

type VoteControllerInterface interface {
	UpdateVoteController(Transaction *models.VotesDTO) (string, error)
	CheckVoteController(Transaction *models.Votes) (bool, error)
}
