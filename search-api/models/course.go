// search-api/models/course.go
package models

type Course struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Instructor   string `json:"instructor"`
	Duration     int    `json:"duration"`
	Level        string `json:"level"`
	Availability bool   `json:"availability"`
}
