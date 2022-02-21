package utils

import (
	"time"

	"github.com/floydjones1/auth-server/model"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWTToken(user model.User) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24*365).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}