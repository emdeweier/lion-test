package models

import "time"

//#region Create

type InsertUser struct {
	Username   string    `json:"username" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Role       string    `json:"role"`
	InsertDate time.Time `json:"insert_date"`
	Address    string    `json:"address" validate:"required"`
}

type InsertUserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

//#endregion

//#region Read

type GetUserByUsername struct {
	Username string `json:"username" validate:"required"`
}

type GetUserData struct {
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	InsertDate time.Time `json:"insert_date"`
	Address    string    `json:"address"`
}

type GetUserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

//#endregion

//#region Update

//#endregion

//#region Delete

//#endregion
