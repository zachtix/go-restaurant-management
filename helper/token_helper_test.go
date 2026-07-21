package helper

import (
	"os"
	model "restaurant-management/models"
	"testing"
)

func TestGenerateAndVerify(t *testing.T) {
	secret := "jwt-secret"
	secret = os.Getenv("SECRET_ACCESS")

	user := model.User{
		Email:      "test@mail.com",
		First_name: "John",
	}

	token, err := GenerateJWT(user, secret, 1)
	if err != nil {
		t.Fatalf("GenerateJWT error: %v", err)
	}

	claims, err := VerifyJWT(token, secret)
	if err != nil {
		t.Fatalf("VerifyJWT error: %v", err)
	}
	if claims["email"] != user.Email {
		t.Errorf("Email in Claims = %v; want %v", claims["email"], user.Email)
	}

	if _, err := VerifyJWT(token, "wrong-secret"); err == nil {
		t.Error("Secret invalid, but verify pass")
	}
}
