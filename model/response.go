package model

type BaseResponse struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	AccessToken string `json:"accessToken"`
}
