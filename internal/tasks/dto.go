package tasks

import (
	"time"

	"github.com/google/uuid"
)

type AssignTaskRequest struct {
	AssignedTo  string `json:"assigned_to" binding:"required,uuid"`
	Title       string `json:"title" binding:"required,min=3,max=255"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date" binding:"required"`
}

type TaskResponse struct {
	ID             uuid.UUID `json:"id"`
	ProjectID      uuid.UUID `json:"project_id"`
	AssignedBy     uuid.UUID `json:"assigned_by"`
	AssignedByName string    `json:"assigned_by_name"`
	AssignedTo     uuid.UUID `json:"assigned_to"`
	AssignedToName string    `json:"assigned_to_name"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	DueDate        time.Time `json:"due_date"`
	CreatedAt      time.Time `json:"created_at"`
}

type DashboardTaskResponse struct {
	TaskID       uuid.UUID `json:"task_id"`
	ProjectID    uuid.UUID `json:"project_id"`
	ProjectTitle string    `json:"project_title"`

	AssignedBy     uuid.UUID `json:"assigned_by"`
	AssignedByName string    `json:"assigned_by_name"`

	AssignedTo     uuid.UUID `json:"assigned_to"`
	AssignedToName string    `json:"assigned_to_name"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`

	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
