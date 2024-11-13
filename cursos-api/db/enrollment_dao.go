// cursos-api/db/enrollment_dao.go

package db

import (
	"context"
	"cursos-app/cursos-api/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCoursesByUserID(userID string) ([]domain.Course, error) {
	var courses []domain.Course
	collection := MongoDB.Collection("enrollments")
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var course domain.Course
		if err := cursor.Decode(&course); err == nil {
			courses = append(courses, course)
		}
	}

	return courses, nil
}
