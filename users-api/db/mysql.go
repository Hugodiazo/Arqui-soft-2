package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB inicializa la conexión a la base de datos
func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:Pirata02@tcp(mysql:3306)/arqsoft2")
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error al verificar la conexión a la base de datos:", err)
	}

	log.Println("Conexión a la base de datos MySQL establecida con éxito")
}
