// users-api/dao/user_dao.go
package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"users-app/users-api/db"
	"users-app/users-api/domain"

	"github.com/bradfitz/gomemcache/memcache"
)

// Crear un nuevo usuario en la base de datos
func CreateUser(user domain.User) error {
	query := "INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, user.Username, user.Email, user.Password, user.Role)
	return err
}

// Obtener un usuario por correo electrónico (con cache)
func GetUserByEmail(email string) (domain.User, error) {
	var user domain.User

	query := "SELECT id, username, email, password, role FROM users WHERE email = ?"
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("Error al ejecutar la consulta SQL:", err)
		return user, err
	}

	return user, nil
}

// Obtener un usuario por ID (con cache)
func GetUserByID(userID int) (domain.User, error) {
	var user domain.User

	// Intentar obtener el usuario de Memcached
	cacheKey := fmt.Sprintf("user_%d", userID)
	item, err := db.Cache.Get(cacheKey)
	if err == nil {
		err := json.Unmarshal(item.Value, &user)
		if err == nil {
			log.Println("Usuario obtenido desde la caché")
			return user, nil
		}
	}

	// Buscar en la base de datos
	query := "SELECT id, username, email, role FROM users WHERE id = ?"
	err = db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return user, err
	}

	// Almacenar en la caché
	userJSON, _ := json.Marshal(user)
	db.Cache.Set(&memcache.Item{Key: cacheKey, Value: userJSON})
	log.Println("Usuario almacenado en la caché")

	return user, nil
}

// Obtener todos los usuarios
func GetAllUsers() ([]domain.User, error) {
	query := "SELECT id, username, email, role FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

// Actualizar un usuario e invalidar cache
func UpdateUser(userID int, user domain.User) error {
	query := "UPDATE users SET username = ?, email = ?, password = ?, role = ? WHERE id = ?"
	_, err := db.DB.Exec(query, user.Username, user.Email, user.Password, user.Role, userID)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("user_%d", userID)
	if err := db.Cache.Delete(cacheKey); err != nil {
		log.Println("No se pudo eliminar el caché:", err)
	}

	return nil
}
