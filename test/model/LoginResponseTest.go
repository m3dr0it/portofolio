package model

import "portofolio/model"

type LoginResponseTest struct {
	Message string              `json:message`
	Data    model.LoginResponse `json:data`
}
