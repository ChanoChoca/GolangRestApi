package models

import "time"

type Usuario struct {
	ID           uint
	Nombre       string
	Apellido     string
	DNI          string
	Telefono     string

	Rol          Rol

	SupervisorID *uint

	Provincia    string
	Localidad    string
	Direccion    string
	CodigoPostal string

	Password     string
	Activo       bool

	CreadoEn     time.Time
	ActualizadoEn time.Time

	Version      uint
}

type Rol string

const (
	ADMIN      Rol = "ADMIN"
	GERENTE    Rol = "GERENTE"
	SUPERVISOR Rol = "SUPERVISOR"
	ASESOR     Rol = "ASESOR"
)

var rolNivel = map[Rol]int{
	ADMIN:      4,
	GERENTE:    3,
	SUPERVISOR: 2,
	ASESOR:     1,
}

func (r Rol) Nivel() int {
	return rolNivel[r]
}

func (r Rol) EsMayorQue(otro Rol) bool {
	return r.Nivel() > otro.Nivel()
}

func (r Rol) EsIgualOMayorQue(otro Rol) bool {
	return r.Nivel() >= otro.Nivel()
}

func (r Rol) EsIgualOMenorQue(otro Rol) bool {
	return r.Nivel() <= otro.Nivel()
}