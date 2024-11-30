package interfaces

import "lion-test/models"

type VoteRepositoryInterface interface {
	UpdateVoteRepository(Transaction *models.VotesDTO) (string, error)
	CheckVoteRepository(Transaction *models.Votes) (bool, error)
}
