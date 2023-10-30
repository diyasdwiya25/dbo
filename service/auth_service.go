package service

import (
	"dbo/entity"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Auth interface {
	GenerateToken(users entity.Users) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type auth struct {
}

func NewAuthService() *auth {
	return &auth{}
}

func (s *auth) GenerateToken(users entity.Users) (string, error) {
	SECRET_KEY := []byte(os.Getenv("SECRET_KEY"))

	claim := jwt.MapClaims{}
	claim["user_id"] = users.Id
	claim["name"] = users.Name
	claim["email"] = users.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *auth) ValidateToken(encodedToken string) (*jwt.Token, error) {

	SECRET_KEY := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token tidak benar")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
