package models

//#region Create

type InsertMovie struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Duration    int64  `json:"duration" validate:"required"`
	Artists     string `json:"artists" validate:"required"`
	Genres      string `json:"genres" validate:"required"`
	WatchUrl    string `json:"watch_url" validate:"required"`
	Viewed      int64  `json:"viewed,omitempty"`
	Voted       int64  `json:"voted"`
}

//#endregion

//#region Read

type GetMovieById struct {
	Id string `json:"id" validate:"required"`
}

type GetMovieData struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
}

type GetMovieDetailData struct {
	GetMovieData
	WatchUrl string `json:"watch_url"`
	Viewed   int64  `json:"viewed"`
	Voted    int64  `json:"voted"`
}

//#endregion

//#region Update

type UpdateMovie struct {
	InsertMovie
}

//#endregion

//#region Delete

//#endregion
