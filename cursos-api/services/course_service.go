// services/course_service.go

package services

import (
	"cursos-app/cursos-api/dao"
	"cursos-app/cursos-api/domain"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllCourses() ([]domain.Course, error) {
	return dao.GetCourses()
}

func CreateCourse(course domain.Course) (domain.Course, error) {
	return dao.CreateCourse(course)
}

func GetCourseByID(id string) (domain.Course, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error al convertir ID a ObjectID:", err)
		return domain.Course{}, err
	}
	return dao.GetCourseByID(objectID)
}

func UpdateCourse(id string, course domain.Course) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error al convertir ID a ObjectID:", err)
		return err
	}
	return dao.UpdateCourse(objectID, course)
}

func EnrollCourse(enrollment domain.Enrollment) error {
	return dao.CreateEnrollment(enrollment)
}

func GetEnrollmentsByUser(userID int) ([]domain.Enrollment, error) {
	return dao.GetEnrollmentsByUser(userID)
}

// DeleteCourseService llama al DAO para eliminar un curso
func DeleteCourseService(courseID string) error {
	deletedCount, err := dao.DeleteCourseDAO(courseID)
	if err != nil {
		return err
	}

	if deletedCount == 0 {
		return errors.New("Curso no encontrado")
	}

	return nil
}
