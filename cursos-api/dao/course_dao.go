// cursos-api/dao/course_dao.go
package dao

import (
	"context"
	"cursos-app/cursos-api/db"
	"cursos-app/cursos-api/domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Crear un curso
func CreateCourse(course domain.Course) (domain.Course, error) {
	course.ID = primitive.NewObjectID() // Genera un nuevo ObjectID
	_, err := db.MongoDB.Collection("courses").InsertOne(context.TODO(), course)
	if err != nil {
		log.Println("Error al crear el curso:", err)
		return domain.Course{}, err
	}
	return course, nil
}

// Obtener todos los cursos
func GetCourses() ([]domain.Course, error) {
	var courses []domain.Course
	cursor, err := db.MongoDB.Collection("courses").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error al obtener cursos:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var course domain.Course
		if err := cursor.Decode(&course); err != nil {
			log.Println("Error al decodificar curso:", err)
			continue
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// Obtener un curso por ID
func GetCourseByID(id string) (domain.Course, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ID inválido:", err)
		return domain.Course{}, err
	}
	var course domain.Course
	err = db.MongoDB.Collection("courses").FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&course)
	if err != nil {
		log.Println("Error al obtener curso:", err)
		return domain.Course{}, err
	}
	return course, nil
}

// Actualizar un curso por ID
func UpdateCourse(id string, updatedCourse domain.Course) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ID inválido:", err)
		return err
	}

	update := bson.M{
		"$set": updatedCourse,
	}
	_, err = db.MongoDB.Collection("courses").UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		log.Println("Error al actualizar el curso:", err)
		return err
	}
	return nil
}

// Crear inscripción de usuario a un curso
func CreateEnrollment(enrollment domain.Enrollment) error {
	_, err := db.MongoDB.Collection("enrollments").InsertOne(context.TODO(), enrollment)
	if err != nil {
		log.Println("Error al inscribir usuario en el curso:", err)
		return err
	}
	return nil
}

func GetCoursesByUserID(userID int) ([]domain.Course, error) {
	var courses []domain.Course
	ctx := context.TODO()

	// Encuentra todas las inscripciones del usuario en enrollments
	cursor, err := db.MongoDB.Collection("enrollments").Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Para cada inscripción, encuentra el curso correspondiente en courses
	for cursor.Next(ctx) {
		var enrollment domain.Enrollment
		if err := cursor.Decode(&enrollment); err == nil {
			courseID, _ := primitive.ObjectIDFromHex(enrollment.CourseID)
			course := domain.Course{}
			_ = db.MongoDB.Collection("courses").FindOne(ctx, bson.M{"_id": courseID}).Decode(&course)
			courses = append(courses, course)
		}
	}

	return courses, nil
}
