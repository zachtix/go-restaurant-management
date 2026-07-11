package helper

import (
	"fmt"
	"restaurant-management/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user model.User, secret string, hour uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["first_name"] = user.First_name
	claims["last_name"] = user.Last_name
	claims["email"] = user.Email
	claims["avatar"] = user.Avatar
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(hour)).Unix()

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return string(t), nil
}

func verifyJWT(token string, secret string) error {
	_, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})

	return err
}
