package service

import "portofolio/model"

func UserLogin(userLogin model.UserLogin) bool {
	var isSigned bool = false
	password := "root"

	if password == userLogin.Password {
		isSigned = true
	}

	return isSigned
}
