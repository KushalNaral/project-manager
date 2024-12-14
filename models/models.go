package models

import "time"

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Project struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Task struct {
	ID          int       `json:"id,omitempty"`
	ProjectId   int       `json:"project_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status,omitempty"`
	Priority    Priority  `json:"priority,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Status string

const (
	StatusOpen       Status = "open"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusToDo       Status = "to_do"
)

type Priority int

const (
	PriorityLow    Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh   Priority = 3
)

type Comment struct {
	ID        int       `json:"id,omitempty"`
	TaskId    int       `json:"task_id,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Tag struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Assignment struct {
	TaskID int `json:"task_id,omitempty"`
	UserID int `json:"user_id,omitempty"`
}

type TaskTag struct {
	TaskID int `json:"task_id,omitempty"`
	TagID  int `json:"tag_id,omitempty"`
}
