// cursos-api/domain/course.go
package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Course representa el modelo de un curso
type Course struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"title"`
	Description  string             `bson:"description"`
	Instructor   string             `bson:"instructor"`
	Duration     int                `bson:"duration"`
	Level        string             `bson:"level"`
	Availability bool               `bson:"availability"`
}
