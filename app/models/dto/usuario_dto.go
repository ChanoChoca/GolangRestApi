package dto

type CreateUsuarioRequest struct {
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	DNI          string `json:"dni"`
	Telefono     string `json:"telefono"`
	Rol          string `json:"rol"`
	SupervisorID *uint  `json:"supervisor_id"`

	Provincia    string `json:"provincia"`
	Localidad    string `json:"localidad"`
	Direccion    string `json:"direccion"`
	CodigoPostal string `json:"codigo_postal"`

	Password string `json:"password"`
}

type UpdateUsuarioRequest struct {
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Telefono     string `json:"telefono"`
	SupervisorID *uint  `json:"supervisor_id"`

	Provincia    string `json:"provincia"`
	Localidad    string `json:"localidad"`
	Direccion    string `json:"direccion"`
	CodigoPostal string `json:"codigo_postal"`

	Activo bool `json:"activo"`
}