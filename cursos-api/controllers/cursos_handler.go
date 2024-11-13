// cursos-api/controllers/course_controller.go
package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"cursos-app/cursos-api/db"
	"cursos-app/cursos-api/domain"
	"cursos-app/cursos-api/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Obtener todos los cursos
func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := services.GetAllCourses()
	if err != nil {
		http.Error(w, "Error al obtener cursos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Crear un curso
func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course domain.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	createdCourse, err := services.CreateCourse(course)
	if err != nil {
		http.Error(w, "Error al crear curso", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCourse)
}

// Obtener un curso por ID
func GetCourseByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	course, err := services.GetCourseByID(id)
	if err != nil {
		http.Error(w, "Curso no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

// Actualizar un curso por ID
func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var course domain.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	if err := services.UpdateCourse(id, course); err != nil {
		http.Error(w, "Error al actualizar el curso", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Curso actualizado con éxito"})
}

// Inscribir un usuario en un curso
func EnrollCourseHandler(w http.ResponseWriter, r *http.Request) {
	var enrollment domain.Enrollment
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	if err := services.EnrollCourse(enrollment); err != nil {
		http.Error(w, "Error al inscribirse en el curso", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inscripción exitosa"})
}

// Handler para obtener los cursos en los que está inscrito un usuario
func GetEnrollmentsByUser(userID string) ([]domain.Enrollment, error) {
	var enrollments []domain.Enrollment
	collection := db.MongoDB.Collection("enrollments") // Asegúrate de que esta colección exista en MongoDB
	filter := bson.M{"user_id": userID}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var enrollment domain.Enrollment
		if err := cursor.Decode(&enrollment); err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

// Obtener inscripciones por usuario
func GetEnrollmentsByUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	// Convertir el userID de string a int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	enrollments, err := services.GetEnrollmentsByUser(userID)
	if err != nil {
		http.Error(w, "Error al obtener inscripciones", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrollments)
}
