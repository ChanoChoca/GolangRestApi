package dto

type LoginRequest struct {
	DNI      string `json:"dni"`
	Password string `json:"password"`
}