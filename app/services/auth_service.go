package services

import (
	"flashpage/app/models"
	"flashpage/app/repositories"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(dni, password string) (*models.Usuario, bool) {
	user, err := repositories.FindByDNI(dni)
	if err != nil {
		return nil, false
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, false
	}

	return user, true
}