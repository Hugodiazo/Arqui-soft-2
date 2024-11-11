// cursos-api/domain/enrollment.go
package domain

type Enrollment struct {
	UserID   int    `json:"user_id"`
	CourseID string `json:"course_id"`
	Status   string `json:"status"`
}
