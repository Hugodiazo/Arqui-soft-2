package domain

type Course struct {
	ID          string   `json:"id"`
	Title       []string `json:"title"`
	Description []string `json:"description"`
	Instructor  []string `json:"instructor"`
	Duration    []int    `json:"duration"`
	Level       []string `json:"level"`
}
