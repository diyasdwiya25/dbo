package mapping

import (
	"dbo/entity"
)

type AuthFormater struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func Auth(user entity.Users, token string) AuthFormater {
	formater := AuthFormater{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}

	return formater
}
