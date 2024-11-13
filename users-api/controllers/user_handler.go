package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"users-app/users-api/domain"
	"users-app/users-api/services"

	"github.com/gorilla/mux"
)

// Handler para registrar un usuario
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	err := services.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado con éxito"})
}

// Handler para iniciar sesión
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var credentials domain.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Llama al servicio para obtener el token y el rol
	token, err := services.LoginUser(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

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
