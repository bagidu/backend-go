package user

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	u := User{
		Password: "secret",
	}

	if err := u.HashPassword(); err != nil {
		t.Errorf("Should not error: %s", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("secret")); err != nil {
		t.Errorf("Should valid compare hash: %s", err.Error())
	}
}

func TestCheckPassword(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)

	u := User{
		Password: string(hashed),
	}

	if err := u.CheckPassword("secret"); err != nil {
		t.Errorf("Password should valid: %s", err.Error())
	}
}
