package mapping

import (
	"dbo/entity"
	"time"
)

type usersFormater struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	RememberToken string    `json:"rememberToken"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func UsersRow(users entity.Users) usersFormater {
	formater := usersFormater{
		Id:            users.Id,
		Name:          users.Name,
		Email:         users.Email,
		Password:      users.Password,
		RememberToken: users.RememberToken,
		CreatedAt:     users.CreatedAt,
		UpdatedAt:     users.UpdatedAt,
	}

	return formater
}

func Users(users []entity.Users) []usersFormater {

	if len(users) == 0 {
		return []usersFormater{}
	}

	var usersArray []usersFormater
	for _, row := range users {
		usersRow := UsersRow(row)
		usersArray = append(usersArray, usersRow)
	}

	return usersArray
}
