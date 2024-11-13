// domain/enrollment.go
package domain

type Enrollment struct {
	UserID   int    `bson:"user_id" json:"user_id"`
	CourseID string `bson:"course_id" json:"course_id"`
	Status   string `bson:"status" json:"status"`
}
