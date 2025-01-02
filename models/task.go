package models

// Task represents a To-Do item
type Task struct {
	ID    int    `json:"id"`    // Unique identifier
	Title string `json:"title"` // Task title
	Done  bool   `json:"done"`  // Completion status
}
