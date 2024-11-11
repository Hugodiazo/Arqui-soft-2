// cursos-api/services/course_service.go
package services

import (
	"cursos-app/cursos-api/dao"
	"cursos-app/cursos-api/domain"
)

func GetAllCourses() ([]domain.Course, error) {
	return dao.GetCourses()
}

func CreateCourse(course domain.Course) (domain.Course, error) {
	return dao.CreateCourse(course)
}

func GetCourseByID(id string) (domain.Course, error) {
	return dao.GetCourseByID(id)
}

func UpdateCourse(id string, course domain.Course) error {
	return dao.UpdateCourse(id, course)
}

func EnrollCourse(enrollment domain.Enrollment) error {
	return dao.CreateEnrollment(enrollment)
}

func GetCoursesByUserID(userID int) ([]domain.Course, error) {
	return dao.GetCoursesByUserID(userID)
}
