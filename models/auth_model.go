package models

//#region Read

type Credential struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CredentialResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

//#endregion
