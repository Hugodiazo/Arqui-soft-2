package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"users-app/users-api/domain"
	"users-app/users-api/services"
	"users-app/users-api/utils"

	"github.com/gorilla/mux"
)

// Handler para registrar un usuario
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Hashear la contraseña antes de registrar el usuario
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error al hashear la contraseña", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	err = services.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado con éxito"})
}

// Handler para iniciar sesión
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginUserHandler iniciado")

	var credentials domain.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Println("Error al decodificar las credenciales:", err)
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	log.Printf("Credenciales recibidas: Email=%s, Password=%s\n", credentials.Email, credentials.Password)

	// Llama al servicio para obtener el token
	token, err := services.LoginUser(credentials)
	if err != nil {
		log.Println("Error en el servicio de login:", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	log.Println("Login exitoso, enviando token")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Handler para obtener todos los usuarios
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler para obtener un usuario por ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByID(id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Handler para actualizar un usuario
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Llama al servicio para actualizar el usuario
	err = services.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "Error al actualizar usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario actualizado con éxito"})
}

// ProtectedHandler es un ejemplo de una ruta protegida por AuthMiddleware
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "¡Accediste a una ruta protegida!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AdminHandler es un ejemplo de una ruta protegida por AdminMiddleware
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "¡Accediste a una ruta de administrador!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
