package models

//#region Create

type VotesDTO struct {
	MovieId      string `json:"movie_id" validate:"required"`
	UserUsername string `json:"user_username" validate:"required"`
}

//#endregion

//#region Read

type Votes struct {
	MovieId      string `json:"movie_id" validate:"required"`
	UserUsername string `json:"user_username" validate:"required"`
}

//#endregion

//#region Update

//#endregion

//#region Delete

//#endregion
