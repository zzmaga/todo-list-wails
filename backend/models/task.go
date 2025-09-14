package models

import "time"

// Priority представляет приоритет задачи
type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

// Task представляет задачу в системе
type Task struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"createdAt"`
	DueAt     *time.Time `json:"dueAt,omitempty"` // Дедлайн (опционалка)
	Priority  Priority   `json:"priority"`        // Приоритет задачи
}
