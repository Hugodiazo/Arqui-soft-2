// users-api/dao/user_dao.go
package dao

import (
	"database/sql"
	"users-app/users-api/db"
	"users-app/users-api/domain"
)

// Crear un nuevo usuario en la base de datos
func CreateUser(user domain.User) error {
	query := "INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, user.Name, user.Email, user.Password, user.Role)
	return err
}

// Obtener un usuario por correo electrónico
func GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, password, role FROM users WHERE email = ?"
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return user, nil
	}
	return user, err
}

// Obtener todos los usuarios
func GetAllUsers() ([]domain.User, error) {
	query := "SELECT id, name, email, role FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role); err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

// Obtener un usuario por ID
func GetUserByID(userID int) (domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, role FROM users WHERE id = ?"
	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	return user, err
}

// Actualizar un usuario
func UpdateUser(userID int, user domain.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, role = ? WHERE id = ?"

	// Usa db.DB.Exec para ejecutar la consulta de actualización
	_, err := db.DB.Exec(query, user.Name, user.Email, user.Password, user.Role, userID)
	if err != nil {
		return err
	}

	return nil
}
