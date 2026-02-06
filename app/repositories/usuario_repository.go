package repositories

import (
	"flashpage/app/config"
	"flashpage/app/models"
)

func FindByDNI(dni string) (*models.Usuario, error) {

	query := `
	SELECT id, nombre, apellido, dni, telefono, rol,
	       supervisor_id, provincia, localidad, direccion,
	       codigo_postal, password, activo,
	       creado_en, actualizado_en, version
	FROM usuarios
	WHERE dni=?`

	row := config.DB.QueryRow(query, dni)

	var u models.Usuario

	err := row.Scan(
		&u.ID,
		&u.Nombre,
		&u.Apellido,
		&u.DNI,
		&u.Telefono,
		&u.Rol,
		&u.SupervisorID,
		&u.Provincia,
		&u.Localidad,
		&u.Direccion,
		&u.CodigoPostal,
		&u.Password,
		&u.Activo,
		&u.CreadoEn,
		&u.ActualizadoEn,
		&u.Version,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func CreateUsuario(u *models.Usuario) error {

	query := `
	INSERT INTO usuarios
	(nombre, apellido, dni, telefono, rol, supervisor_id,
	 provincia, localidad, direccion, codigo_postal,
	 password, activo, creado_en, actualizado_en, version)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,NOW(),NOW(),1)
	`

	res, err := config.DB.Exec(query,
		u.Nombre, u.Apellido, u.DNI, u.Telefono, u.Rol,
		u.SupervisorID,
		u.Provincia, u.Localidad, u.Direccion, u.CodigoPostal,
		u.Password, u.Activo,
	)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return nil
}

func FindUsuarioByID(id uint) (*models.Usuario, error) {

	query := `SELECT id,nombre,apellido,dni,telefono,rol,
	supervisor_id,provincia,localidad,direccion,codigo_postal,
	password,activo,creado_en,actualizado_en,version
	FROM usuarios WHERE id=?`

    row := config.DB.QueryRow(query, id)

	var u models.Usuario

	err := row.Scan(
		&u.ID, &u.Nombre, &u.Apellido, &u.DNI, &u.Telefono,
		&u.Rol, &u.SupervisorID,
		&u.Provincia, &u.Localidad, &u.Direccion, &u.CodigoPostal,
		&u.Password, &u.Activo,
		&u.CreadoEn, &u.ActualizadoEn, &u.Version,
	)

	if err != nil {
		return nil, err
	}

	return &u, err
}

func ListUsuariosPaged(limit, offset int) ([]models.Usuario, int, error) {

	var total int
	err := config.DB.QueryRow(`SELECT COUNT(*) FROM usuarios`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := config.DB.Query(`
		SELECT id, nombre, apellido, dni, telefono, rol, supervisor_id,
		        provincia, localidad, direccion, codigo_postal, password,
				activo, creado_en, actualizado_en, version
		FROM usuarios ORDER BY id LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []models.Usuario

	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(
		    &u.ID, &u.Nombre, &u.Apellido, &u.DNI, &u.Telefono,
		    &u.Rol, &u.SupervisorID,
		    &u.Provincia, &u.Localidad, &u.Direccion, &u.CodigoPostal, &u.Password,
		    &u.Activo, &u.CreadoEn, &u.ActualizadoEn, &u.Version,
	    )
		if err != nil {
			return nil, 0, err
		}
		list = append(list, u)
	}

	return list, total, nil
}

func UpdateUsuario(u *models.Usuario) error {

	query := `
	UPDATE usuarios SET
	nombre=?, apellido=?, telefono=?, supervisor_id=?,
	provincia=?, localidad=?, direccion=?, codigo_postal=?,
	activo=?, actualizado_en=NOW(), version=version+1
	WHERE id=?`

	_, err := config.DB.Exec(query,
		u.Nombre, u.Apellido, u.Telefono, u.SupervisorID,
		u.Provincia, u.Localidad, u.Direccion, u.CodigoPostal,
		u.Activo, u.ID,
	)

	return err
}

func DisableUsuario(id uint) error {
	_, err := config.DB.Exec(`UPDATE usuarios SET activo=false WHERE id=?`, id)
	return err
}