package services

import (
	"errors"
	"flashpage/app/models"
	"flashpage/app/repositories"

	"golang.org/x/crypto/bcrypt"
)

func CreateUsuario(u *models.Usuario) error {

	if u.Rol == models.ASESOR && u.SupervisorID == nil {
		return errors.New("asesor requiere supervisor")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	u.Activo = true

	return repositories.CreateUsuario(u)
}

func GetUsuario(id uint) (*models.Usuario, error) {
	return repositories.FindUsuarioByID(id)
}

func ListUsuarios(page, pageSize int) (*models.Page[models.Usuario], error) {
	
	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	items, total, err := repositories.ListUsuariosPaged(pageSize, offset)
	if err != nil {
		return nil, err
	}

	totalPages := (total + pageSize - 1) / pageSize

	return &models.Page[models.Usuario]{
		Items: items,
		Total: total,
		Page: page,
		PageSize: pageSize,
		TotalPages: totalPages,
	}, nil
}

func UpdateUsuario(u *models.Usuario) error {

	if u.Rol == models.ASESOR && u.SupervisorID == nil {
		return errors.New("asesor requiere supervisor")
	}

	return repositories.UpdateUsuario(u)
}

func DisableUsuario(id uint) error {
	return repositories.DisableUsuario(id)
}