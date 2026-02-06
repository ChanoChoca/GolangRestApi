package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN no definida")
	}

	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
    	CREATE TABLE IF NOT EXISTS usuarios (
    	    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    	    nombre VARCHAR(100) NOT NULL,
    	    apellido VARCHAR(100) NOT NULL,
    	    dni VARCHAR(20) NOT NULL UNIQUE,
    	    telefono VARCHAR(50),
        
    	    rol VARCHAR(20) NOT NULL,
        
    	    supervisor_id BIGINT UNSIGNED,
    	    FOREIGN KEY (supervisor_id) REFERENCES usuarios(id),
        
    	    provincia VARCHAR(100) NOT NULL,
    	    localidad VARCHAR(100) NOT NULL,
    	    direccion VARCHAR(150) NOT NULL,
    	    codigo_postal VARCHAR(20) NOT NULL,
        
    	    password VARCHAR(255) NOT NULL,
    	    activo BOOLEAN NOT NULL DEFAULT true,
        
    	    creado_en TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            actualizado_en TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        
    	    version BIGINT UNSIGNED NOT NULL DEFAULT 1
        )
	;`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error al crear la tabla:", err)
	}

	log.Println("MySQL conectado")
}