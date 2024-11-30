package models

//#region Create

type InsertGenre struct {
	Id     int    `json:"id,omitempty"`
	Genre  string `json:"genre" validate:"required"`
	Viewed int64  `json:"viewed"`
}

//#endregion

//#region Read

type GetGenreByName struct {
	Genre string `json:"genre" validate:"required"`
}

type GetGenreData struct {
	Id     int    `json:"id,omitempty"`
	Genre  string `json:"genre"`
	Viewed int64  `json:"viewed"`
}

//#endregion

//#region Update

type UpdateGenre struct {
	InsertGenre
}

//#endregion

//#region Delete

//#endregion
